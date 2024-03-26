package repository

import (
	"context"
	"fmt"
	"selling"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SellingPostgres struct {
	pg *pgxpool.Pool
}

func NewSellingPostgres(pg *pgxpool.Pool) *SellingPostgres {
	return &SellingPostgres{pg: pg}
}
func (r *SellingPostgres) CreateSelling(userId int, list selling.SellingList) (selling.SellingList, error) {
	var listres selling.SellingList
	tr, err := r.pg.Begin(context.Background())
	if err != nil {
		return listres, err
	}
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, date, description, url, price) VALUES ($1,$2.$3,$4,$5) RETURNING *", sellingListTable)
	row := tr.QueryRow(context.Background(), createListQuery, list.Title, list.Date, list.Description, list.PicURL, list.Price)
	if err := row.Scan(&listres.Id, &listres.Title, &listres.Description, &listres.Date, &listres.PicURL, &listres.Price); err != nil {
		tr.Rollback(context.Background())
		return listres, err
	}
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1,$2)", userListTable)
	_, err = tr.Exec(context.Background(), createUserListQuery, userId, listres.Id)
	if err != nil {
		tr.Rollback(context.Background())
		return listres, err
	}
	return listres, tr.Commit(context.Background())
}
func (r *SellingPostgres) ListSellings(userId int, order string) ([]selling.SellingList, error) {
	var lists []selling.SellingList
	query := ""
	if userId != 0 {
		query = fmt.Sprintf("SELECT sl.id, sl.title, sl.description, sl.date, sl.url, sl.price, ul.id (where ul.id is not null)  FROM %s sl LEFT JOIN %s ul on sl.id = ul.list_id",
			sellingListTable, userListTable)
	} else {
		query = fmt.Sprintf("SELECT id, title, description, date, url, price FROM %s",
			sellingListTable)
	}
	row, err := r.pg.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		k := selling.SellingList{}
		err := row.Scan(&k.Id, &k.Title, &k.Description, &k.Date, &k.PicURL, &k.Price)
		if err != nil {
			return nil, err
		}
		lists = append(lists, k)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}
