package main

import (
	"github.com/eSlider/itunes-ranking-service/api"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
