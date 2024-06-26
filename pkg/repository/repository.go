package repository

import (
	"selling"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user selling.User) (selling.User, error)
	SignUser(username, password string) (selling.User, error)
}
type Selling interface {
	CreateSelling(userId int, list selling.SellingList) (selling.SellingList, error)
	ListSellings(userId int, order, sortby string, page int) (map[string]interface{}, error)
}

type Repository struct {
	Authorization
	Selling
}

func NewRepository(pg *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(pg),
		Selling:       NewSellingPostgres(pg),
	}
}
