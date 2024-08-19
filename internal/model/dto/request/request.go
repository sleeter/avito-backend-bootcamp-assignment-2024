package request

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
type DummyLogin struct {
	UserType string `json:"user_type"`
}
type Login struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type CreateFlat struct {
	HouseId int32  `json:"house_id" binding:"required"`
	Price   int32  `json:"price" binding:"required"`
	Rooms   *int32 `json:"rooms"`
	Status  string `json:"status"`
}
type UpdateFlat struct {
	Id     int32  `json:"id" binding:"required"`
	Status string `json:"status"`
}

type House struct {
	Address   string  `json:"address" binding:"required"`
	Year      int32   `json:"year" binding:"required"`
	Developer *string `json:"developer"`
}

type Subscriber struct {
	HouseId int32  `json:"house_id"`
	Email   string `json:"email" binding:"required"`
}
