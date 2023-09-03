package models

import (
	"time"
)

type User struct {
	Id           int       `db:"id"`
	Username     string    `db:"username"`
	EmailAddress string    `db:"email_address"`
	Password     string    `db:"password"`
	IsValidated  bool      `db:"is_validated"`
	CreatedAt    time.Time `db:"created_at"`
}

func NewUser(id int, username, email, password string) *User {
	return &User{
		Id:           id,
		Username:     username,
		EmailAddress: email,
		Password:     password,
		IsValidated:  false,
		CreatedAt:    time.Now(),
	}
}
