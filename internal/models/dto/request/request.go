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

type Flat struct {
	HouseId int32 `json:"house_id"`
	Price   int32 `json:"price"`
	Rooms   int32 `json:"rooms"`
}

type House struct {
	Address   string `json:"address"`
	Year      int32  `json:"year"`
	Developer string `json:"developer"`
	Status    string `json:"status"`
}
