package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wordcollection/app/models"

	"github.com/dgrijalva/jwt-go"
)

type AuthFrom struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

func GenerateToken(username string) string {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iss":      "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err)
	}

	return tokenString
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var authFrom AuthFrom
	var jwt JWT
	var error Error

	json.NewDecoder(r.Body).Decode(&authFrom)
	if authFrom.Username == "" {
		error.Message = "ユーザネームが入力されていません。"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if authFrom.PassWord == "" {
		error.Message = "パスワードが入力されていません。"
		respondWithError(w, http.StatusBadRequest, error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//create new user & register to database
	newUser := &models.User{}
	newUser.Name = authFrom.Username
	newUser.PassWord = authFrom.PassWord
	err := newUser.CreateUser()
	if err != nil {
		error.Message = "すでにその名前は使用されています"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	// gererate jwt token & return token
	w.Header().Set("Content-Type", "application/json")
	jwt.Token = GenerateToken(newUser.Name)
	responseJSON(w, jwt)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var authFrom AuthFrom
	user := models.User{}
	var jwt JWT
	var error Error

	json.NewDecoder(r.Body).Decode(&authFrom)
	if authFrom.Username == "" {
		error.Message = "ユーザネームが入力されていません。"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if authFrom.PassWord == "" {
		error.Message = "パスワードが入力されていません。"
		respondWithError(w, http.StatusBadRequest, error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := authFrom.PassWord
	row := models.Db.QueryRow("select * from users where name=$1", authFrom.Username)
	fmt.Println("row")
	fmt.Println(row)
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.PassWord, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "ユーザが存在しません"
			respondWithError(w, http.StatusBadRequest, error)
			return
		} else {
			fmt.Println("else error")
			log.Fatalln(err)
		}
	}

	hashedPassword := user.PassWord
	if hashedPassword != models.Encrypt(password) {
		error.Message = "パスワードが適切ではありません"
		respondWithError(w, http.StatusUnauthorized, error)
		return
	}

	jwt.Token = GenerateToken(user.Name)

	responseJSON(w, jwt)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func varidation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}
