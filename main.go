package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting Vault API server")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":820", router))
}
