package entity

type Subscriber struct {
	Id      int32  `db:"id"`
	HouseId int32  `db:"house_id"`
	Email   string `db:"email"`
}
