package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Login() http.HandlerFunc {

	type loginRequest struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}

	type Metadata struct {
		Username string `json:"username"`
	}

	type Auth struct {
		ClientToken    string      `json:"client_token"`
		Accessor       string      `json:"accessor"`
		Policies       []string    `json:"policies"`
		TokenPolicies  []string    `json:"token_policies"`
		Metadata       Metadata    `json:"metadata"`
		LeaseDuration  int         `json:"lease_duration"`
		Renewable      bool        `json:"renewable"`
		EntityId       string      `json:"entity_id"`
		TokenType      string      `json:"token_type"`
		Orphan         bool        `json:"orphan"`
		MfaRequirement interface{} `json:"mfa_requirement"`
		NumUses        int         `json:"num_uses"`
	}

	type loginResponse struct {
		RequestId     string      `json:"request_id"`
		LeaseId       string      `json:"lease_id"`
		Renewable     bool        `json:"renewable"`
		LeaseDuration int         `json:"lease_duration"`
		Data          interface{} `json:"data"`
		WrapInfo      interface{} `json:"wrap_info"`
		Warnings      interface{} `json:"warnings"`
		Auth          Auth        `json:"auth"`
	}

	type ErrorMessage struct {
		Errors []string `json:"errors"`
	}

	validCredentials := map[string]string{
		"rob": "password",
		"nic": "password",
	}

	failedAttempts := make(map[string]int)

	return func(w http.ResponseWriter, r *http.Request) {
		userName := mux.Vars(r)["username"]

		var ip string
		if len(r.Header["X-Forwarded-For"]) == 0 {
			//fmt.Printf("LENGTH IS %v\n", len(r.Header["X-Forwarded-For"]))
			ip = r.RemoteAddr
		} else {
			ip = r.Header["X-Forwarded-For"][0]
		}

		if r.Method != "POST" && r.Method != "PUT" {
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

		if password, ok := validCredentials[req.Username]; ok {
			if req.Password == password {
				resp := loginResponse{
					RequestId:     requestId,
					LeaseId:       "",
					Renewable:     false,
					LeaseDuration: 0,
					Data:          nil,
					WrapInfo:      nil,
					Warnings:      nil,
					Auth: Auth{
						ClientToken:   id,
						Accessor:      accessor,
						Policies:      []string{"default", "policy1", "policy2"},
						TokenPolicies: []string{"default", "policy1", "policy2"},
						Metadata: Metadata{
							Username: req.Username,
						},
						LeaseDuration:  2764800,
						Renewable:      renewable,
						EntityId:       entityId,
						TokenType:      tokenType,
						Orphan:         orphan,
						MfaRequirement: nil,
						NumUses:        numUses,
					},
				}
				delete(failedAttempts, req.Username)
				log.Printf("Login Successful: User: %s Source IP Address: %s \n", req.Username, ip)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			} else {
				failedAttempts[req.Username]++
				resp := ErrorMessage{
					Errors: []string{"permission denied"},
				}

				log.Printf("Permission Denied: User: %s Source IP Address: %s \n", req.Username, ip)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				if failedAttempts[req.Username] == 3 {
					log.Printf("Multiple failed login attempts: User: %s Source IP Address: %s\n", req.Username, ip)

					message := Message{
						Ip: ip,
					}

					messageJSON, err := json.Marshal(message)
					if err != nil {
						log.Println("error encoding message JSON:", err)
						w.WriteHeader(http.StatusInternalServerError)
						return
					}

					pubSub(string(messageJSON))
				}
			}
		}
	}
}
