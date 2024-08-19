package entity

type User struct {
	Id       string `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	UserType string `db:"user_type"`
	Dummy    bool   `db:"dummy"`
}
