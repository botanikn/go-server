package entities

type User struct {
	Id       int
	Username string
	PassHash string
	Email    string
	Role     string
}
