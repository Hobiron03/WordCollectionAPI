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
	cmd := `insert into words (user_id, word, mean, pronounce, genre, color, created_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err = Db.Exec(cmd, u.ID, word, mean, pronounce, genre, color, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) GetWordAll() (words []Word, err error) {
	cmd := `select id, user_id, word, mean, pronounce, genre, color from words where user_id = $1;`
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
	cmd := `select id, user_id, word, mean, pronounce, genre, color from words where id = $1;`

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

func (u *User) GetNewestWordID() (id int, err error) {
	var newestID int
	cmd := `select MAX(id) from words where user_id = $1;`
	err = Db.QueryRow(cmd, u.ID).Scan(&newestID)

	return newestID, err
}

func (w *Word) UpdateWord() (err error) {
	cmd := `update words set word = $1, mean = $2, pronounce = $3, genre = $4, color = $5 where id = $6;`
	_, err = Db.Exec(cmd, w.Word, w.Mean, w.Pronounce, w.Genre, w.Color, w.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (w *Word) DeleteWord() (err error) {
	cmd := `delete from words where id = $1;`
	_, err = Db.Exec(cmd, w.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) DeleteWordAll() (err error) {
	cmd := `delete from words where user_id = $1;`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
