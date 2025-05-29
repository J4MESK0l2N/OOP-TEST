package main

import (
  "log"
  "net/http"

  "test-oop-golang/handlers"
)

func main() {
	http.Handle("/api/v1/users", handlers.AuthMiddleWare(http.HandlerFunc(handlers.GetUsers)))
	http.HandleFunc("/api/v1/user/create", handlers.CreateUser)
	http.HandleFunc("/api/v1/login", handlers.Login)

	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}