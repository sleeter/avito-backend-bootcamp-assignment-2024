package entity

import "time"

type House struct {
	Id        int32     `db:"id"`
	Address   string    `db:"address"`
	Year      int32     `db:"year"`
	Developer *string   `db:"developer"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"update_at"`
}
