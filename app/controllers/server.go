package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wordcollection/config"

	"github.com/gorilla/mux"
)

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func StartAPIServer() error {

	router := mux.NewRouter()

	router.HandleFunc("/", topHandler)
	router.HandleFunc("/fetchmyword", fetchMyWordHandler)
	router.HandleFunc("/addmyword", addMyWordHandler)
	router.HandleFunc("/deletemyword", deleteMyWordHandler)
	router.HandleFunc("/alldeletemyword", allDeleteMyWordHandler)
	router.HandleFunc("/updatemyword", updateMyWordHandler)
	router.HandleFunc("/deleteuser", deleteUserHandler)
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signin", signinHandler)
	router.HandleFunc("/logout", logoutHandler)
	router.HandleFunc("/validation", logoutHandler)

	return http.ListenAndServe(":"+config.Config.Port, router)
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

}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return nil
}
