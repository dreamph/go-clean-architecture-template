package app

import (
	"backend/cmd/api/handler"
	_ "backend/docs"
	"backend/internal/config"
	"backend/internal/core/api"
	apicommons "backend/internal/core/api/commons"
	"backend/internal/core/api/middleware"
	"backend/internal/core/auth/jwt"
	"backend/internal/core/database/query"
	"backend/internal/core/database/query/bun"
	"backend/internal/core/errors"
	"backend/internal/core/logger"
	"backend/internal/core/logger/zap"
	"backend/internal/core/utils"
	authusecase "backend/internal/modules/auth/usecase"
	companyusecase "backend/internal/modules/company/usecase"
	"backend/internal/repository"
	"os"
	"time"

	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	fibercasbin "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	c := config.New("config/local.env", "config/config_api.yml", false)

	cfg, err := config.LoadConfig(c)
	if err != nil {
		logger.LogErrorAndExit(err)
	}

	log := zap.NewLogger(&logger.Options{
		FilePath: cfg.Log.FilePath,
		Level:    cfg.Log.Level,
		Format:   cfg.Log.Format,
		ProdMode: cfg.App.ProdMode,
	})
	defer log.Sync()

	bunDB, err := bun.Connect(&bun.Options{
		Host:           cfg.Database.Host,
		Port:           cfg.Database.Port,
		DBName:         cfg.Database.DBName,
		User:           cfg.Database.User,
		Password:       cfg.Database.Password,
		ConnectTimeout: cfg.Database.ConnectTimeout,
		Logger:         log.WithOptionsAddCallerSkip(5),
		TraceEnable:    false,
	})
	if err != nil {
		logger.LogErrorAndExit(err, log)
	}
	defer bun.Close(bunDB)

	appDB := &query.AppIDB{BunDB: bunDB}
	dbTx := bun.NewDBTx(bunDB)

	jwtToken := jwt.NewJwtToken(&jwt.ConfigInfo{
		Issuer:            cfg.App.Jwt.Issuer,
		SecretKey:         []byte(cfg.App.Jwt.SecretKey),
		Expiration:        time.Duration(cfg.App.Jwt.ExpirationMinutes) * time.Minute,
		RefreshExpiration: time.Duration(cfg.App.Jwt.RefreshExpirationMinutes) * time.Minute,
	})

	companyRepository := repository.NewCompanyRepository(appDB)
	companyUseCase := companyusecase.NewCompanyUseCase(dbTx, companyRepository)
	authUseCase := authusecase.NewAuthUseCase(jwtToken)

	app := fiber.New(api.NewServerConfig())
	app.Use(middleware.NewRequestTime())
	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: log.Logger(),
		Fields: []string{"requestId", "latency", "status", "method", "url"},
	}))
	app.Use(favicon.New())

	app.Use(fiberrecover.New())
	router := app.Group(cfg.App.ContextPath)

	apiHandler := api.NewApiHandler(&api.HandlerOption{
		JwtToken: jwtToken,
	})

	authz := fibercasbin.New(fibercasbin.Config{
		ModelFilePath: cfg.Permissions.ModelFile,
		PolicyAdapter: fileadapter.NewAdapter(cfg.Permissions.PolicyFile),
		Lookup: func(c *fiber.Ctx) string {
			info := apiHandler.GetUserRequestInfo(c)
			return info.Scope
		},
		Forbidden: func(c *fiber.Ctx) error {
			return apicommons.ResponseError(c, errors.ErrPermissionDenied)
		},
	})

	jwtAuth := middleware.NewJwtAuth(&middleware.JwtAuthOption{
		JwtToken: jwtToken,
		Enable:   true,
	})

	authPermissions := middleware.NewAuthPermissions(authz, true)

	handler.NewSwaggerAPIHandler(apiHandler, router).Init(cfg.App.ContextPath, cfg.App.ProdMode)
	handler.NewAppAPIHandler(apiHandler, router).Init()
	handler.NewAuthAPIHandler(apiHandler, router, authUseCase).Init()
	handler.NewCompanyAPIHandler(apiHandler, router, companyUseCase).Init(jwtAuth, authPermissions)

	port := utils.NumberToString(cfg.App.Port)

	log.Info("Starting...")
	log.Info("App : " + cfg.App.Name)
	log.Info("env : " + cfg.App.Env)
	log.Info("Swagger URL : http://localhost:" + port + cfg.App.ContextPath + "/swagger/index.html")

	app.Hooks().OnShutdown(func() error {
		log.Info("OnShutdown")
		return nil
	})

	go func() {
		err = app.Listen(":" + port)
		if err != nil {
			logger.LogErrorAndExit(err, log)
		}
	}()

	log.Info("Running...")

	utils.WaitSignal(func(sig os.Signal) {
		log.Info("Gracefully shutting down...")
		log.Info("Waiting for all request to finish")
		err = app.Shutdown()
		if err != nil {
			logger.LogErrorAndExit(err, log)
		}
		log.Info("Running cleanup tasks...")
		log.Info("Server was successful shutdown.")
	})
}
