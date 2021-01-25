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

	fmt.Println("access signup")
	json.NewDecoder(r.Body).Decode(&authFrom)
	fmt.Println(authFrom.Username)
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
