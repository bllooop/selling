package selling

import "github.com/jackc/pgtype"

type SellingList struct {
	Id          int         `json:"id"`
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" db:"description"`
	PicURL      string      `json:"url"  db:"url"`
	Price       int         `json:"price" db:"price"`
	Date        pgtype.Date `json:"date"`
}

/*type SellListsItem struct {
	Id            int
	SellintListId int
	SellinItemId  int
}*/

type UsersSell struct {
	Id     int
	UserId int
	SellId int
}
