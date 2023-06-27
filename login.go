package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strings"
)

func Login() http.HandlerFunc {

	type loginRequest struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}

	type loginResponse struct {
		Message string `json:"message"`
	}

	validCredentials := map[string]string{
		"rob": "password",
		"nic": "password",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userName := mux.Vars(r)["username"]

		//remoteAddr := r.RemoteAddr

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("error parsing remote address:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ip = strings.TrimSuffix(ip, "%")
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req loginRequest
		req.Username = userName
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println("error decoding JSON")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if password, ok := validCredentials[req.Username]; ok && req.Password == password {
			resp := loginResponse{
				Message: "Login Successful",
			}

			log.Printf("Login Successful: User: %s Source IP Address: %s \n", req.Username, ip)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			resp := loginResponse{
				Message: "Permission Denied",
			}

			log.Printf("Permission Denied: User: %s Source IP Address: %s \n", req.Username, ip)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
