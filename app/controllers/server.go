package controllers

import (
	"fmt"
	"net/http"
	"todo_app/config"
)

func StartAPIServer() error {
	// http.HandleFunc("/", nil)
	fmt.Println("start server")
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
