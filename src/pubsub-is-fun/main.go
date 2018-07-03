package main

import (
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func main() {
	log.Print("aaa")
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "my-project-id"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new topic.
	topicName := "my-new-topic"

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	log.Printf("Topic %v created.\n", topic)
}
