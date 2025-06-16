package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDatabase(databaseUrl string) (*DB, error) {
	config, errConfig := pgxpool.ParseConfig(databaseUrl)
	if errConfig != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", errConfig)
	}
	pool, errPool := pgxpool.NewWithConfig(context.Background(), config)
	if errPool != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", errPool)
	}
	if errPing := pool.Ping(context.Background()); errPing != nil {
		return nil, fmt.Errorf("failed to ping database: %w", errPing)
	}
	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) GetConn(ctx context.Context) (*pgxpool.Conn, error) {
	conn, errConn := db.Pool.Acquire(ctx)
	if errConn != nil {
		return nil, fmt.Errorf("failed to acquire connection: %w", errConn)
	}
	return conn, nil
}
