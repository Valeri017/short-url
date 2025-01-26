package main

import (
	"fmt"
	"net/http"
	"short-url/configs"
	"short-url/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: router,
	}

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	fmt.Println("Сервер запущен на 8081")
	server.ListenAndServe()
}
