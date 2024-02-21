package database

import (
	"database/sql"
	"time"
)

type DbPoolOptions struct {
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	MaxIdleConns int
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	MaxOpenConns int
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	ConnMaxLifetime time.Duration
}

var DbPoolDefault = &DbPoolOptions{
	MaxIdleConns:    20,
	MaxOpenConns:    100,
	ConnMaxLifetime: 30 * time.Minute,
}

func SetConnectionsPool(db *sql.DB, pool *DbPoolOptions) {
	db.SetMaxIdleConns(pool.MaxIdleConns)
	db.SetMaxOpenConns(pool.MaxOpenConns)
	db.SetConnMaxLifetime(pool.ConnMaxLifetime)
}
