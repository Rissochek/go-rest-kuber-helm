package main

import (
	"log"
	"net/http"

	"github.com/Rissochek/go-rest-kuber-helm/internal/server"
)

func main() {
	router := server.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
