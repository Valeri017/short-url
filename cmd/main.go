package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"short-url/configs"
	"short-url/internal/auth"
	"short-url/internal/link"
	"short-url/pkg/db"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		fmt.Printf("Пока пока %v\n", sig)
		os.Exit(0)
	}()

	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: router,
	}
	//Repository
	linkRepository := link.NewLinkRepository(db)
	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	fmt.Println("Сервер запущен на 8081")
	server.ListenAndServe()
}
