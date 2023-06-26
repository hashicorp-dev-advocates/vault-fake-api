package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/login", Login()).Methods("POST")

	return r
}
