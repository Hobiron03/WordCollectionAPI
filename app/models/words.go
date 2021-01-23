package models

import "log"

type Word struct {
	ID        int
	UserID    int
	Word      string
	Mean      string
	Pronounce string
	Genre     string
	color     string
}

func (u *User) CreateWord(word string, mean string, pronounce string, genre string, color string) (err error) {
	cmd := `insert into words (user_id, word, mean, pronounce, genre, color, created_at) values (?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, word, mean, pronounce, genre, color)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
