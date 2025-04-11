package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"short-url/configs"
	"short-url/internal/auth"
	"short-url/pkg/db"
	"syscall"
)

func main() {

	// Создание канала для сигналов
	sigChan := make(chan os.Signal, 1)

	// Регистрация сигналов
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Горутина для обработки сигнала
	go func() {
		sig := <-sigChan
		fmt.Printf("Пока пока %v\n", sig)
		// Выполнение необходимых действий перед завершением
		os.Exit(0)
	}()

	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
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
	// Горутина для обработки сигнала
	go func() {
		sig := <-sigChan
		fmt.Printf("Получен сигнал: %v\n", sig)
		// Выполнение необходимых действий перед завершением
		os.Exit(0)
	}()

	fmt.Println("Ожидание сигнала...")
	select {} // Бесконечное ожидание
}
