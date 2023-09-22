package user

import "fmt"

type user struct{}

func New() *user {
	return &user{}
}

func (u *user) Update() {
	fmt.Println("Updated!")
}
