package main

import (
	"fmt"
	"wordcollection/app/models"
)

func main() {
	// controllers.StartAPIServer()

	fmt.Println(models.Db)

	user1 := &models.User{}
	user1.Name = "first user"
	user1.PassWord = "test"
	user1.CreateUser()
}
