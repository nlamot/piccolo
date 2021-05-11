package gcp

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)


type FirestoreClient interface {
	Close() error
}

func ProvideFirestoreClient(config *GCPConfiguration) FirestoreClient {
	client, err := firestore.NewClient(context.Background(), config.ProjectID)
	if err != nil {
		log.Printf("firestore.NewClient: %v", err)
		return nil
	}
	return client
}