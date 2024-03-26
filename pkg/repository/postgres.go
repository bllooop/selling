package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	userListTable    = "userlist"
	sellingListTable = "sellinglist"
	userSellingTable = "usersellingtable"
)

const (
	host     = "localhost"
	port     = "5433"
	user     = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
	password = "54321"
)

func NewPostgresDB() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode))
	if err != nil {
		return nil, err
	}
	return db, nil
}
