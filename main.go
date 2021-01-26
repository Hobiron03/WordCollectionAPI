package main

import "wordcollection/app/controllers"

func main() {
	controllers.StartAPIServer()

	// fmt.Println(models.Db)

	// userの作成例
	// user1 := &models.User{}
	// user1.Name = "first user"
	// user1.PassWord = "test2"
	// user1.CreateUser()

	//userの取得例 & wordの追加例
	// user, _ := models.GetUser(1)
	// user.CreateWord("word23232", "mean2", "pronounce2", "genre2", "colo2")
	// fmt.Println(user.GetUserWordAll())

	//userの取得例2
	// user, _ := models.GetUserByName("first user")
	// user.CreateWord("word23232daf", "meandaf2", "pronoudfance2", "genrdfafe2", "colofad2")
	// fmt.Println(user)
	// fmt.Println(user.GetUserWordAll())

	// wordの取得例
	// word, _ := models.GetWord(2)
	// fmt.Println(word)

	// // wordの更新例
	// word.Word = "dfadfaa"
	// word.Mean = "更新してみた"
	// word.UpdateWord()
}
