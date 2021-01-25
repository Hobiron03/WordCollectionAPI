package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wordcollection/app/models"
)

type AuthFrom struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
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
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}

	// gererate jwt token & return token
	w.Header().Set("Content-Type", "application/json")
	// responseJSON(w, newUser)
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
