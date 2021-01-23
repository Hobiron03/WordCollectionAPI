package controllers

import (
	"net/http"
	"wordcollection/config"
)

func StartAPIServer() error {

	http.HandleFunc("/", topHandler)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}

func topHandler(w http.ResponseWriter, r *http.Request) {

}
