package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"piccolo.com/planner/pkg/common/config"
)


type FirestoreClient interface {
	Close() error
}

func ProvideFirestoreClient(config *config.GCPConfiguration) FirestoreClient {
	client, err := firestore.NewClient(context.Background(), config.ProjectID)
	if err != nil {
		log.Printf("firestore.NewClient: %v", err)
		return nil
	}
	return client
}