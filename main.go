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

	testUser := model.NewUser(randomId, "hola", "si")
	cache.Save(*testUser)

	user := cache.Find(randomId).(model.User)
	fmt.Println(user)
}

