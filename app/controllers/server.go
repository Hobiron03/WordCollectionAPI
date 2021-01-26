package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"wordcollection/app/models"
	"wordcollection/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Error struct {
	Message string
}

type Username struct {
	Username string
}

type WordID struct {
	ID int `json:"id"`
}

type WordEdit struct {
	ID        int
	Word      string
	Pronounce string
	Mean      string
	Genre     string
	Color     string
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
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	json.NewEncoder(w).Encode(data)

}

func StartAPIServer() error {

	router := mux.NewRouter()

	router.HandleFunc("/", top)
	router.HandleFunc("/fetchmyword", TokenVerifyMiddleWare(fetchMyWordHandler)).Methods("POST")
	router.HandleFunc("/addmyword", TokenVerifyMiddleWare(addMyWordHandler))
	router.HandleFunc("/deletemyword", TokenVerifyMiddleWare(deleteMyWordHandler))
	router.HandleFunc("/alldeletemyword", TokenVerifyMiddleWare(allDeleteMyWordHandler))
	router.HandleFunc("/updatemyword", TokenVerifyMiddleWare(updateMyWordHandler))
	router.HandleFunc("/deleteuser", TokenVerifyMiddleWare(deleteUserHandler))
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signin", signinHandler).Methods("POST")
	router.HandleFunc("/validation", TokenVerifyMiddleWare(validation)).Methods("GET", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug:              true,
		OptionsPassthrough: false,
	})

	handler := c.Handler(router)

	return http.ListenAndServe(":"+config.Config.Port, handler)
}

func top(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// ここを追加
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
}

func fetchMyWordHandler(w http.ResponseWriter, r *http.Request) {
	var username Username
	username.Username = r.FormValue("username")

	user, err := models.GetUserByName(username.Username)
	if err != nil {
		log.Fatalln(err)
	}

	words, err := user.GetWordAll()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("fetchMyWordHandler")
	fmt.Println(words)
	responseJSON(w, words)
}

func addMyWordHandler(w http.ResponseWriter, r *http.Request) {
	var addEditWordPost AddEditWordPost

	addEditWordPost.Username = r.FormValue("username")
	addEditWordPost.Word = r.FormValue("word")
	addEditWordPost.Pronounce = r.FormValue("pronounce")
	addEditWordPost.Mean = r.FormValue("mean")
	addEditWordPost.Genre = r.FormValue("genre")
	addEditWordPost.Color = r.FormValue("color")

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

	//todo: id返してあげないとダメ
	// responseJSON(w)

	w.WriteHeader(http.StatusOK)
}

func deleteMyWordHandler(w http.ResponseWriter, r *http.Request) {
	var wordID WordID
	json.NewDecoder(r.Body).Decode(&wordID)

	word, err := models.GetWord(wordID.ID)
	if err != nil {
		log.Fatalln(err)
	}

	err = word.DeleteWord()
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
}

func updateMyWordHandler(w http.ResponseWriter, r *http.Request) {
	var postedWord models.Word

	postedWord.ID, _ = strconv.Atoi(r.FormValue("id"))
	postedWord.Word = r.FormValue("word")
	postedWord.Pronounce = r.FormValue("pronounce")
	postedWord.Mean = r.FormValue("mean")
	postedWord.Genre = r.FormValue("genre")
	postedWord.Color = r.FormValue("color")

	word, err := models.GetWord(postedWord.ID)
	if err != nil {
		log.Fatalln(err)
	}

	word.Word = postedWord.Word
	word.Mean = postedWord.Mean
	word.Pronounce = postedWord.Pronounce
	word.Genre = postedWord.Genre
	word.Color = postedWord.Color
	err = word.UpdateWord()
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
}

func allDeleteMyWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("TopHandler")
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
