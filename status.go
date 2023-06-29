package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Status() http.HandlerFunc {

	type Status struct {
		Type         string `json:"type"`
		Initialized  bool   `json:"initialized"`
		Sealed       bool   `json:"sealed"`
		T            int    `json:"t"`
		N            int    `json:"n"`
		Progress     int    `json:"progress"`
		Nonce        string `json:"nonce"`
		Version      string `json:"version"`
		Migration    bool   `json:"migration"`
		ClusterName  string `json:"cluster_name"`
		ClusterId    string `json:"cluster_id"`
		RecoverySeal bool   `json:"recovery_seal"`
		BuildDate    string `json:"build_date"`
		StorageType  string `json:"storage_type"`
		HaEnabled    bool   `json:"ha_enabled"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var StatusResponse Status

		StatusResponse.Type = "shamir"
		StatusResponse.Initialized = true
		StatusResponse.Sealed = false
		StatusResponse.T = 1
		StatusResponse.N = 1
		StatusResponse.Progress = 0
		StatusResponse.Nonce = ""
		StatusResponse.Version = "1.11.0"
		StatusResponse.Migration = false
		StatusResponse.ClusterName = "vault-cluster-30a35421"
		StatusResponse.ClusterId = "62b51e90-effc-19b8-b53b-027c98b2848b"
		StatusResponse.RecoverySeal = false
		StatusResponse.StorageType = "inmem"
		StatusResponse.HaEnabled = false
		StatusResponse.BuildDate = "2022-05-03T08:34:11Z"

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		respBody, err := json.Marshal(StatusResponse)
		if err != nil {
			log.Println("error encoding message JSON:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)

	}
}
