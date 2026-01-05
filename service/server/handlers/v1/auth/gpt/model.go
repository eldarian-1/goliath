package gpt

type User struct {
	ID       string
	Email    string
	Password string // hash
	Role     string
}
