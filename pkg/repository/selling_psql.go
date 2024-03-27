package repository

import (
	"context"
	"fmt"
	"selling"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}
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
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, date, description, url, price) VALUES ($1,$2,$3,$4,$5) RETURNING *", sellingListTable)
	row := tr.QueryRow(context.Background(), createListQuery, list.Title, time.Now(), list.Description, list.PicURL, list.Price)
	if err := row.Scan(&listres.Id, &listres.Title, &listres.Price, &listres.Date, &listres.Description, &listres.PicURL); err != nil {
		tr.Rollback(context.Background())
		return listres, err
	}
	var username string
	loginQuery := fmt.Sprintf("SELECT username FROM %s WHERE id=$1", userListTable)
	row = tr.QueryRow(context.Background(), loginQuery, userId)
	if err := row.Scan(&username); err != nil {
		tr.Rollback(context.Background())
		return listres, err
	}
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_login, user_id, list_id) VALUES ($1,$2,$3)", userSellingTable)
	_, err = tr.Exec(context.Background(), createUserListQuery, username, userId, listres.Id)
	if err != nil {
		tr.Rollback(context.Background())
		return listres, err
	}
	return listres, tr.Commit(context.Background())
}
func (r *SellingPostgres) ListSellings(userId int, order, sortby string, page int) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	limit := 10
	offset := limit * (page - 1)
	data["Page"] = r.pagination("users", limit, page)
	var lists []selling.SellingList
	query := ""
	if userId == 0 {
		query = fmt.Sprintf(`SELECT sl.id, sl.title, sl.description, sl.date, sl.url, sl.price, ul.user_login FROM %s sl LEFT JOIN %s ul on sl.id = ul.list_id 
		ORDER BY %s %s limit %d offset %d`, sellingListTable, userSellingTable, order, sortby, limit, offset)
	} else {
		query = fmt.Sprintf(`SELECT sl.id, sl.title, sl.description, sl.date, sl.url, sl.price, ul.user_login FROM %s sl LEFT JOIN %s ul on sl.id = ul.list_id 
		WHERE ul.user_id=%d ORDER BY %s %s limit %d offset %d`, sellingListTable, userSellingTable, userId, order, sortby, limit, offset)
	}
	row, err := r.pg.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var err error
		k := selling.SellingList{}
		if userId == 0 {
			err = row.Scan(&k.Id, &k.Title, &k.Description, &k.Date, &k.PicURL, &k.Price, &k.UserLogin)
		} else {
			err = row.Scan(&k.Id, &k.Title, &k.Description, &k.Date, &k.PicURL, &k.Price, &k.UserLogin)
		}
		if err != nil {
			return nil, err
		}
		lists = append(lists, k)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	data["Sellings"] = lists
	return data, nil
}

func (r *SellingPostgres) pagination(table string, limit, page int) *Pagination {
	var (
		tmpl        = Pagination{}
		recordcount int
	)

	sqltable := fmt.Sprintf("SELECT count(id) FROM %s", table)

	r.pg.QueryRow(context.Background(), sqltable).Scan(&recordcount)

	total := (recordcount / limit)

	remainder := (recordcount % limit)
	if remainder == 0 {
		tmpl.TotalPage = total
	} else {
		tmpl.TotalPage = total + 1
	}

	tmpl.CurrentPage = page
	tmpl.RecordPerPage = limit

	if page <= 0 {
		tmpl.Next = page + 1
	} else if page < tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = page + 1
	} else if page == tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = 0
	}

	return &tmpl
}
