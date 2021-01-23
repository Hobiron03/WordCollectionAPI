package main

import (
	"fmt"
	"wordcollection/app/models"
)

func main() {
	// controllers.StartAPIServer()

	user, _ := models.GetUser(1)
	fmt.Println(user)
}
