package entity

import "time"

type House struct {
	Id           int32     `db:"id"`
	Address      string    `db:"address"`
	Year         int32     `db:"year"`
	Developer    string    `db:"developer"`
	CreationDate time.Time `db:"creation_date"`
	UpdateDate   time.Time `db:"update_date"`
}
