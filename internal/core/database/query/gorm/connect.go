package gorm

import (
	"backend/internal/core/database"
	"backend/internal/core/database/query/gorm/otel"
	applogger "backend/internal/core/logger"

	errs "github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"

	"fmt"

	"gorm.io/gorm/logger"
)

type Options struct {
	Host           string
	Port           string
	DBName         string
	User           string
	Password       string
	ConnectTimeout int64
	Logger         applogger.Logger
	PoolOptions    *database.DbPoolOptions
	TraceEnable    bool
}

// Connect ...
func Connect(options *Options) (*gorm.DB, error) {
	appLogger := options.Logger
	connection := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable connect_timeout=%v TimeZone=Asia/Bangkok",
		options.Host,
		options.Port,
		options.User,
		options.DBName,
		options.Password,
		options.ConnectTimeout,
	)
	dbLogger := NewDbLogger(appLogger)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		return nil, errs.Wrap(err, "Failed to connect to database.")
	}
	db = db.Debug()
	sqlDB, _ := db.DB()
	if err = sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, errs.Wrap(err, "Failed to connect to database.")
	}

	dbPool := database.DbPoolDefault
	if options.PoolOptions != nil {
		dbPool = options.PoolOptions
	}
	database.SetConnectionsPool(sqlDB, dbPool)

	appLogger.Info("DB connect successfully..")

	if options.TraceEnable {
		err := db.Use(otel.New())
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func Close(db *gorm.DB) {
	if db != nil {
		sqlDb, err := db.DB()
		if err != nil {
			return
		}
		sqlDb.Close()
	}
}

func NewDbLogger(l applogger.Logger) logger.Interface {
	gormLogger := zapgorm2.New(l.Logger())
	gormLogger.IgnoreRecordNotFoundError = true
	gormLogger.LogLevel = gormlogger.Warn
	gormLogger.SetAsDefault()
	return gormLogger
}

func AutoMigrate(db *gorm.DB, dst ...interface{}) {
	_ = db.AutoMigrate(dst...)
}
