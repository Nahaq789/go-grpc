package model

type User struct {
	id    int
	name  string
	email string
}

func NewUser(id int, name, email string) *User {
	return &User{id: id, name: name, email: email}
}
