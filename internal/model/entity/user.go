package entity

type User struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Type     string `db:"type"`
	Dummy    bool   `db:"dummy"`
}

const (
	USERTYPE_CLIENT    = "client"
	USERTYPE_MODERATOR = "moderator"
)
