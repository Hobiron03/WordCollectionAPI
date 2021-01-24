package main

import "wordcollection/app/controllers"

func main() {
	controllers.StartAPIServer()
	// fmt.Println(models.Db)
	// user1 := &models.User{}
	// user1.Name = "first user2"
	// user1.PassWord = "test2"
	// user1.CreateUser()

	// user, _ := models.GetUser(1)
	// // user.CreateWord("word23232", "mean2", "pronounce2", "genre2", "colo2")
	// fmt.Println(user.GetUserWordAll())
}
