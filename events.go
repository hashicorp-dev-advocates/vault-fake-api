package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
	"os"
)

type Message struct {
	Ip string `json:"ip"`
}

func pubSub(payload string) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		log.Fatalf("GCP_PROJECT_ID needs to be set")
	}

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the id for the new topic.
	topicID := os.Getenv("GCP_PUBSUB_TOPIC_ID")
	if topicID == "" {
		log.Fatalf("GCP_PUBSUB_TOPIC_ID needs to be set")
	}

	// Creates the new topic.
	topic := client.Topic(topicID)

	_ = topic.Publish(ctx, &pubsub.Message{Data: []byte(payload)})
	topic.Stop()
}
