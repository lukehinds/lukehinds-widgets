package main

import (
	"log"
	"net/http"

	"github.com/lukehinds/widgets/internal/server"
)

func main() {
	s := server.NewServer()
	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", s.Router()); err != nil {
		log.Fatal(err)
	}
}