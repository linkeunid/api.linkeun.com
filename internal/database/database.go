package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	maxIdleConnections    = 10
	maxOpenConnections    = 100
	connectionMaxLifetime = 5 * time.Minute
	connectionTimeout     = 3 * time.Second // Timeout for the database connection attempt
)

type DB struct {
	*gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	// Set a timeout for the database connection
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database: " + err.Error())
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("failed to get sql.DB from gorm DB: " + err.Error())
	}

	// Configure the connection pool
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetConnMaxLifetime(connectionMaxLifetime)

	// Test the database connection with the provided context
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, errors.New("failed to ping database within the timeout: " + err.Error())
	}

	return &DB{gormDB}, nil
}

// WithContext wraps GORM queries with a context.
func (db *DB) WithContext(ctx context.Context) *gorm.DB {
	return db.DB.WithContext(ctx)
}
