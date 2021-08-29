package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitPostgresDatabase() *pgxpool.Pool {
	username := os.Getenv("POSTGRES_USER")
	passwd := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, passwd, "localhost", db)

	connection, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		panic(err)
	}
	return connection
}
