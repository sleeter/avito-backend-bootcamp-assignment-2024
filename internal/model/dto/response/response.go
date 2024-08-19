package response

import "time"

type Register struct {
	UserId string `json:"user_id"`
}
type Login struct {
	Token string `json:"token"`
}

type Error struct {
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Code      int32  `json:"code"`
}

type Flat struct {
	Id      int32  `json:"id"`
	HouseId int32  `json:"house_id"`
	Price   int32  `json:"price"`
	Rooms   int32  `json:"rooms"`
	Status  string `json:"status"`
}

type House struct {
	Id        int32     `json:"id"`
	Address   string    `json:"address"`
	Year      int32     `json:"year"`
	Developer *string   `json:"developer"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
