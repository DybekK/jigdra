package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type database struct {
	Connection *pgx.Conn
}

func Initialize() (*pgx.Conn, error) {
	username := os.Getenv("POSTGRES_USER")
	passwd := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, passwd, "localhost", db)
	return pgx.Connect(context.Background(), url)
}
