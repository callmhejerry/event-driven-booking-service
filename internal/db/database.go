package db

import (
	"context"
	"time"

	"github.com/callmhejerry/event-driven-booking-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	*Queries
	pool *pgxpool.Pool
}

func NewDatabase(dbConfig config.DatabaseConfig) (*Database, error) {
	poolConfig, err := pgxpool.ParseConfig(dbConfig.DbUrl)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = 4
	poolConfig.MinConns = 2
	poolConfig.MinIdleConns = 2
	poolConfig.MaxConnIdleTime = time.Second * 20

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	queries := New(pool)

	database := &Database{
		Queries: queries,
		pool:    pool,
	}

	return database, nil
}
