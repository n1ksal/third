package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создание логгера
	logger := log.New(os.Stdout, "myapp: ", log.LstdFlags)

	// Создание и настройка сервера
	srv := server.NewServer(logger)

	// Запуск сервера
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
