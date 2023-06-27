package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	if gcpProjectID == "" {
		fmt.Println("ERROR: GCP_PROJECT_ID environment variable is not set.")
		os.Exit(1)
	}

	gcpPubSubTopicId := os.Getenv("GCP_PUBSUB_TOPIC_ID")
	if gcpPubSubTopicId == "" {
		fmt.Println("ERROR: GCP_PUBSUB_TOPIC_ID environment variable is not set.")
		os.Exit(1)
	}

	log.Println("starting Vault API server")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8200", router))

}
