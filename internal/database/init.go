package database

import (
	"context"
	"fmt"

	"auth-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConn interface {
	Connect(ctx context.Context) (*pgxpool.Pool, error)
}

type dbConn struct {
	dbConfig *config.DBConfig
}

func NewConnector(dbc *config.DBConfig) DBConn {
	return &dbConn{
		dbConfig: dbc,
	}
}

func (c *dbConn) Connect(ctx context.Context) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(c.dbConfig.ConnString)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error initializing connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return pool, nil
}
