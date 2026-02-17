package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)
 
func InitDB(ctx context.Context, dbUrl string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
 
	_, err = pool.Exec(ctx, 
	`CREATE TABLE IF NOT EXISTS links(
		short_link VARCHAR(100) NOT NULL UNIQUE,
		main_link VARCHAR(500) NOT NULL UNIQUE
	)`)

	if err != nil {
		log.Fatal(err)
	}

	return pool
}