package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	fullname  string
	age       int
	createdAt time.Time
}

func NewUser(fullname string, age int) (*User, error) {

	if fullname == "" {
		return nil, errors.New("el nombre no puede ir vacio")
	}
	return &User{
		fullname:  fullname,
		age:       age,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ClearUserName() {
	u.fullname = ""
}

func (u *User) PrintUserInfo() {
	fmt.Println(u)
}
