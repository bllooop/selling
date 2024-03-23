package repository

import (
	"selling"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user selling.User) (int, error)
	SignUser(username, password string) (selling.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(pg *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(pg),
	}
}
