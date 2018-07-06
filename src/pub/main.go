package main

import (
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

const topicName string = "message2"
const projectID string = "my-project-id"
const subID string = "message-sub"

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	topic := CreateTopicIfNotExist(client)
	log.Printf("Topic %v created.\n", topic)

	CreateSubscriptionIfNotExist(client, topic)
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

// CreateTopicIfNotExist ...
func CreateTopicIfNotExist(c *pubsub.Client) *pubsub.Topic {
	ctx := context.Background()
	topic := c.Topic(topicName)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return topic
	}

	topic, err = c.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatal("Failed to create topic")
	}
	return topic
}

// CreateSubscriptionIfNotExist ...
func CreateSubscriptionIfNotExist(c *pubsub.Client, topic *pubsub.Topic) *pubsub.Subscription {
	ctx := context.Background()
	sub := c.Subscription(subID)
	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return sub
	}

	sub, err = c.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		log.Fatal(err)
	}
	return sub
}
