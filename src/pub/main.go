package main

import (
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "my-project-id"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new topic.
	topicName := "message2"

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	log.Printf("Topic %v created.\n", topic)

	_, err = client.CreateSubscription(ctx, "message-sub",
		pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		log.Fatalf("Failed to Create Subscription: %v", err)
	}
	log.Print("Subscription was created.\n")

	// Publish Message
	msg1 := &pubsub.Message{Data: []byte("hello!")}
	res := topic.Publish(ctx, msg1)
	if _, err := res.Get(ctx); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	log.Printf("Message %v sent.\n", msg1)

	msg2 := &pubsub.Message{Data: []byte("good bye!")}
	res = topic.Publish(ctx, msg2)
	if _, err := res.Get(ctx); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	log.Printf("Message %v sent.\n", msg2)
}
