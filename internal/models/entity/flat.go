package entity

type Flat struct {
	Id       int32 `db:"id"`
	House_id int32 `db:"house_id"`
	Number   int32 `db:"number"`
	Price    int32 `db:"price"`
	Rooms    int32 `db:"rooms"`
	Status   int32 `db:"status"`
}
