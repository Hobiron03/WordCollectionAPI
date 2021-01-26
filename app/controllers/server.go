package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"wordcollection/app/models"
	"wordcollection/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Error struct {
	Message string
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

type Username struct {
	Username string
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func StartAPIServer() error {

	router := mux.NewRouter()

	router.HandleFunc("/", topHandler)
	router.HandleFunc("/fetchmyword", TokenVerifyMiddleWare(fetchMyWordHandler))
	router.HandleFunc("/addmyword", addMyWordHandler)
	router.HandleFunc("/deletemyword", deleteMyWordHandler)
	router.HandleFunc("/alldeletemyword", allDeleteMyWordHandler)
	router.HandleFunc("/updatemyword", updateMyWordHandler)
	router.HandleFunc("/deleteuser", deleteUserHandler)
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signin", signinHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler)
	router.HandleFunc("/validation", TokenVerifyMiddleWare(validation))

	return http.ListenAndServe(":"+config.Config.Port, router)
}

func topHandler(w http.ResponseWriter, r *http.Request) {

}

func fetchMyWordHandler(w http.ResponseWriter, r *http.Request) {
	var username Username
	json.NewDecoder(r.Body).Decode(&username)

	user, err := models.GetUserByName(username.Username)
	if err != nil {
		log.Fatalln(err)
	}

	words, err := user.GetWordAll()
	if err != nil {
		log.Fatalln(err)
	}

	responseJSON(w, words)
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("エラー発生")
				}
				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "tokenが正しくありません"
			respondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}

	})
}
