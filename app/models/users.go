package models

import (
	"log"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		password,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, CreateUUID(), u.Name, Encrypt(u.PassWord), time.Now())

	return err
}

func GetUserByID(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, password, created_at
	from users where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func GetUserByName(username string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, password, created_at
	from users where name = ?`

	err = Db.QueryRow(cmd, username).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
