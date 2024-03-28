package repository

import (
	"context"
	"fmt"
	"selling"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthPostgres struct {
	pg *pgxpool.Pool
}

func NewAuthPostgres(pg *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{pg: pg}
}

func (r *AuthPostgres) CreateUser(user selling.User) (selling.User, error) {
	var res selling.User
	query := fmt.Sprintf(`INSERT INTO %s (username,password) VALUES ($1,$2) RETURNING id, username`, userListTable)
	row := r.pg.QueryRow(context.Background(), query, user.Username, user.Password)
	if err := row.Scan(&res.Id, &res.Username); err != nil {
		return res, err
	}
	return res, nil

}
func (r *AuthPostgres) SignUser(username, password string) (selling.User, error) {
	var user selling.User
	query := fmt.Sprintf(`SELECT id,username,password FROM %s WHERE username=$1 AND password=$2`, userListTable)
	res := r.pg.QueryRow(context.Background(), query, username, password)
	err := res.Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}
