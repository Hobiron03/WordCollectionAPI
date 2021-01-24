package controllers

import (
	"fmt"
	"net/http"
	"wordcollection/config"
)

func StartAPIServer() error {

	http.HandleFunc("/", topHandler)
	http.HandleFunc("/fetchmyword", fetchMyWordHandler)
	http.HandleFunc("/addmyword", addMyWordHandler)
	http.HandleFunc("/deletemyword", deleteMyWordHandler)
	http.HandleFunc("/alldeletemyword", allDeleteMyWordHandler)
	http.HandleFunc("/updatemyword", updateMyWordHandler)
	http.HandleFunc("/deleteuser", deleteUserHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/validation", logoutHandler)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func fetchMyWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func addMyWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func deleteMyWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func allDeleteMyWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func updateMyWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
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
