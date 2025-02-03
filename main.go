package main

import (
	"back/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/generate", handlers.PasswordHandler)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
