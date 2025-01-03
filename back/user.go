package main

import (
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Voted bool   `json:"voted"`
	Vote  string `json:"vote"`
}

func NewUser(id string, name string) *User {
	user := &User{
		ID:    id,
		Name:  name,
		Voted: false,
	}

	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	return user
}
