package main

import (
	"context"
	"fmt"
	"log/slog"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/firestore"
)

const projectID = "example-project"

type Issue struct {
	Title string
}

// $ DATASTORE_EMULATOR_HOST=localhost:8081 FIRESTORE_EMULATOR_HOST=localhost:8080 go run main.go
func main() {
	ctx := context.Background()

	if err := firestoreAccess(ctx); err != nil {
		slog.ErrorContext(ctx, "Firestore access", "error", err)
	}

	if err := datastoreAccess(ctx); err != nil {
		slog.ErrorContext(ctx, "Datastore access", "error", err)
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
