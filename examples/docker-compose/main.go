package main

import (
	"context"
	"fmt"
	"log/slog"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
)

const projectID = "example-project"

type Issue struct {
	Title string
}

// $ FIRESTORE_EMULATOR_HOST=localhost:8080 DATASTORE_EMULATOR_HOST=localhost:8081 PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go
func main() {
	ctx := context.Background()

	if err := firestoreAccess(ctx); err != nil {
		slog.ErrorContext(ctx, "Firestore access", "error", err)
	}

	if err := datastoreAccess(ctx); err != nil {
		slog.ErrorContext(ctx, "Datastore access", "error", err)
	}

	if err := pubsubAccess(ctx); err != nil {
		slog.ErrorContext(ctx, "Pub/Sub access", "error", err)
	}

	slog.InfoContext(ctx, "Completed!")
}

func firestoreAccess(ctx context.Context) error {
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("new client: %w", err)
	}
	defer c.Close()

	docRef := c.Collection("issues").Doc("1")

	if _, err := docRef.Set(ctx, Issue{Title: "Firestore Issue"}); err != nil {
		return fmt.Errorf("set: %w", err)
	}

	doc, err := docRef.Get(ctx)
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}

	var issue Issue
	if err := doc.DataTo(&issue); err != nil {
		return fmt.Errorf("data to: %w", err)
	}

	slog.InfoContext(ctx, "Firestore", "issue", issue)
	return nil
}

func datastoreAccess(ctx context.Context) error {
	c, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("new client: %w", err)
	}
	defer c.Close()

	key := datastore.IDKey("Issue", 1, nil)

	if _, err := c.Put(ctx, key, &Issue{Title: "Datastore Issue"}); err != nil {
		return fmt.Errorf("put: %w", err)
	}

	var issue Issue
	if err := c.Get(ctx, key, &issue); err != nil {
		return fmt.Errorf("get: %w", err)
	}

	slog.InfoContext(ctx, "Datastore", "issue", issue)
	return nil
}

func pubsubAccess(ctx context.Context) error {
	c, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("new client: %w", err)
	}
	defer c.Close()

	topic, err := c.CreateTopic(ctx, "example-topic")
	if err != nil {
		return fmt.Errorf("create topic: %w", err)
	}

	sub, err := c.CreateSubscription(ctx, "example-subscription", pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		return fmt.Errorf("create subscription: %w", err)
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello, Pub/Sub!"),
	})
	if _, err := result.Get(ctx); err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	var received string
	cctx, cancel := context.WithCancel(ctx)
	if err := sub.Receive(cctx, func(_ context.Context, msg *pubsub.Message) {
		received = string(msg.Data)
		msg.Ack()
		cancel()
	}); err != nil {
		return fmt.Errorf("receive: %w", err)
	}

	slog.InfoContext(ctx, "Pub/Sub", "received", received)
	return nil
}
