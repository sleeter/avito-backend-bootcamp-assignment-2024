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
	HouseId int32  `json:"house_id" binding:"required,gte=1"`
	Price   int32  `json:"price" binding:"required,gte=0"`
	Rooms   *int32 `json:"rooms"`
	Status  string `json:"status"`
}
type UpdateFlat struct {
	Id     int32  `json:"id" binding:"required,gte=1"`
	Status string `json:"status"`
}

type House struct {
	Address   string  `json:"address" binding:"required"`
	Year      int32   `json:"year" binding:"required,gte=0"`
	Developer *string `json:"developer" binding:""`
}

type Subscriber struct {
	HouseId int32  `json:"house_id"`
	Email   string `json:"email" binding:"required,email"`
}
