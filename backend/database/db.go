package database

import (
	"context"
	"fmt"
	"os"
	"server/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Connect() error {
	config := config.Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	dbUrl := config.Url()

	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return err
	}

	pool = dbPool
	return nil
}

func GetPool() (*pgxpool.Pool, error) {
	if pool == nil {
		return nil, fmt.Errorf("database connection pool is not initialized")
	}

	return pool, nil
}

func Close() {
	if pool != nil {
		pool.Close()
	}
}
