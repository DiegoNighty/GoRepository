package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Model interface {
	Id() uuid.UUID
}

type User struct {
	Identifier uuid.UUID
	Name string
	Password string
}

func (user User) Id() uuid.UUID {
	return user.Identifier
}

func (user User) String() string {
	return fmt.Sprintf("Usuario con la id: %v y el nombre %v", user.Id(), user.Name)
}