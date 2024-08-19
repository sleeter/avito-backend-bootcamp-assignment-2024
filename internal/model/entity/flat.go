package entity

type Flat struct {
	Id      int32  `db:"id"`
	HouseId int32  `db:"house_id"`
	Number  int32  `db:"number"`
	Price   int32  `db:"price"`
	Rooms   int32  `db:"rooms"`
	Status  string `db:"status"`
}

const (
	FLATSTATUS_CREATED       = "created"
	FLATSTATUS_APPROVED      = "approved"
	FLATSTATUS_DECLINED      = "declined"
	FLATSTATUS_ON_MODERATION = "on_moderation"
)
