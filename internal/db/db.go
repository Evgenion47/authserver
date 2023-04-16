package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const dsn = "host=localhost port=5432 user=postgres password=7123 dbname=postgres sslmode=disable"

func New(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, dsn)
}
