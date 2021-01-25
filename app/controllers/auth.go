package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthFrom struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var authFrom AuthFrom
	var error Error

	json.NewDecoder(r.Body).Decode(&authFrom)
	if authFrom.Username == "" {
		error.Message = "ユーザネームが入力されていません。"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if authFrom.Password == "" {
		error.Message = "パスワードが入力されていません。"
		respondWithError(w, http.StatusBadRequest, error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//register new user to database

	// gererate jwt token & return token
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func varidation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}
