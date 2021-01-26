package models

import (
	"log"
	"time"
)

type Word struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Word      string `json:"word"`
	Mean      string `json:"mean"`
	Pronounce string `json:"pronounce"`
	Genre     string `json:"genre"`
	Color     string `json:"color"`
}

func (u *User) CreateWord(word string, mean string, pronounce string, genre string, color string) (err error) {
	cmd := `insert into words (user_id, word, mean, pronounce, genre, color, created_at) values (?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, u.ID, word, mean, pronounce, genre, color, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) GetWordAll() (words []Word, err error) {
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

func GetWord(id int) (word Word, err error) {
	cmd := `select id, user_id, word, mean, pronounce, genre, color from words where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&word.ID,
		&word.UserID,
		&word.Word,
		&word.Mean,
		&word.Pronounce,
		&word.Genre,
		&word.Color,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return word, err
}

func (w *Word) UpdateWord() (err error) {
	cmd := `update words set word = ?, mean = ?, pronounce = ?, genre = ?, color = ? where id = ?`
	_, err = Db.Exec(cmd, w.Word, w.Mean, w.Pronounce, w.Genre, w.Color, w.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (w *Word) DeleteWord() (err error) {
	cmd := `delete from words where id = ?`
	_, err = Db.Exec(cmd, w.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
