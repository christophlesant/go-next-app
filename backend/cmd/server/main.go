package main

import (
	"log"
	"net/http"

	"github.com/christophlesant/go-next-app/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Router
	r := mux.NewRouter()
	r.HandleFunc("/app/page", handlers.HomePage).Methods("GET") // Home page

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
	})

	handler := c.Handler(r)
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
