package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	r.PathPrefix("/v1").Handler(http.StripPrefix("/v1", r))
	r.HandleFunc("/auth/userpass/login/{username}", Login()).Methods("POST", "PUT")
	r.HandleFunc("/sys/seal-status", Status()).Methods("GET")
	return r
}
