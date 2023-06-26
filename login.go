package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Login() http.HandlerFunc {

	type loginRequest struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}

	type loginResponse struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println("error decoding JSON")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp := loginResponse{
			Message: "Permission Denied",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
