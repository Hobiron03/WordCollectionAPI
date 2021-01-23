package models

import (
	"log"
	"time"
)

type Word struct {
	ID        int
	UserID    int
	Word      string
	Mean      string
	Pronounce string
	Genre     string
	Color     string
}

func (u *User) CreateWord(word string, mean string, pronounce string, genre string, color string) (err error) {
	cmd := `insert into words (user_id, word, mean, pronounce, genre, color, created_at) values (?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, word, mean, pronounce, genre, color, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) GetUserWordAll() (words []Word, err error) {
	cmd := `select id, user_id, word, mean, pronounce, genre, color from words where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)

	for rows.Next() {
		var word Word
		err = rows.Scan(&word.ID, &word.UserID, &word.Word, &word.Mean, &word.Pronounce, &word.Genre, &word.Color)
		if err != nil {
			log.Fatalln(err)
		}
		words = append(words, word)
	}
	rows.Close()

	return words, err
}
