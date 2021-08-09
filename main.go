package main

import (
	"GoRepository/model"
	"GoRepository/repository"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	cache := repository.NewRepository()

	var randomId uuid.UUID

	if newUUID, err := uuid.NewUUID(); err != nil {
		fmt.Println("Error while creating a new UUID")
		return
	} else {
		randomId = newUUID
	}

	test := model.User{Identifier: randomId, Name: "hola", Password: "si"}
	cache.Save(test)

	user := cache.Find(randomId).(model.User)
	fmt.Println(user)
}

