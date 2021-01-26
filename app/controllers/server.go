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

type Username struct {
	Username string
}

type AddEditWordPost struct {
	Username  string
	Word      string
	Pronounce string
	Mean      string
	Genre     string
	Color     string
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

	router.HandleFunc("/fetchmyword", TokenVerifyMiddleWare(fetchMyWordHandler))
	router.HandleFunc("/addmyword", TokenVerifyMiddleWare(addMyWordHandler))
	router.HandleFunc("/deletemyword", TokenVerifyMiddleWare(deleteMyWordHandler))
	router.HandleFunc("/alldeletemyword", TokenVerifyMiddleWare(allDeleteMyWordHandler))
	router.HandleFunc("/updatemyword", TokenVerifyMiddleWare(updateMyWordHandler))
	router.HandleFunc("/deleteuser", TokenVerifyMiddleWare(deleteUserHandler))
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signin", signinHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler)
	router.HandleFunc("/validation", TokenVerifyMiddleWare(validation))

	return http.ListenAndServe(":"+config.Config.Port, router)
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
	var addEditWordPost AddEditWordPost
	json.NewDecoder(r.Body).Decode(&addEditWordPost)

	user, err := models.GetUserByName(addEditWordPost.Username)
	if err != nil {
		log.Fatalln(err)
	}

	err = user.CreateWord(
		addEditWordPost.Word,
		addEditWordPost.Mean,
		addEditWordPost.Pronounce,
		addEditWordPost.Genre,
		addEditWordPost.Color,
	)

	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
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
