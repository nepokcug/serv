package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Server starting on http://localhost:8080")

	// Используем логгер для Fatal ошибок
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("Server failed:", err)
	}
}
