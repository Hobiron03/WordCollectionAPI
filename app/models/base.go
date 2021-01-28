package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
	tableNameWord = "words"
)

func init() {
	// Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// fmt.Println(Db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	databaseURL := os.Getenv("DATABASE_URL")
	pgURL, err := pq.ParseURL(databaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pgURL)
	Db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(Db)

	//sqlite
	// cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING UNIQUE,
	// 	password STRING,
	// 	created_at DATETIME)`, tableNameUser)
	// _, err = Db.Exec(cmdU)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//postgres
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid TEXT NOT NULL UNIQUE,
		name TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP)`, tableNameUser)
	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	//sqlite
	// cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	user_id INTEGER,
	// 	word TEXT,
	// 	mean TEXT,
	// 	pronounce TEXT,
	// 	genre TEXT,
	// 	color TEXT,
	// 	created_at DATETIME)`, tableNameWord)

	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		word TEXT,
		mean TEXT,
		pronounce TEXT,
		genre TEXT,
		color TEXT,
		created_at TIMESTAMP)`, tableNameWord)
	_, err = Db.Exec(cmdW)
	if err != nil {
		log.Fatalln(err)
	}

}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
