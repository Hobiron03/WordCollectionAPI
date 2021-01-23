package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"wordcollection/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
	tableNameWord = "words"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	fmt.Println(Db)
	if err != nil {
		log.Fatal(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		word TEXT,
		mean TEXT,
		pronounce TEXT,
		genre TEXT,
		color TEXT,
		created_at DATETIME)`, tableNameWord)
	Db.Exec(cmdW)

}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))

	return cryptext
}
