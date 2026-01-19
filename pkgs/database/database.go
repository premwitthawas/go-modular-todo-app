package databasePkg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPoolDb() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/db?sslmode=disable")
	if err != nil {
		log.Fatalf("error: can't connected db %v", err)
	}
	return pool
}
