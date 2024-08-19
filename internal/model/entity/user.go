package entity

type User struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Type     string `db:"type"`
}

const (
	USERTYPE_CLIENT    = "client"
	USERTYPE_MODERATOR = "moderator"
)

var UserTypes = []string{USERTYPE_CLIENT, USERTYPE_MODERATOR}
