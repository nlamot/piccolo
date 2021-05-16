package gcp

import (
	"context"

	"cloud.google.com/go/firestore"
)


type FirestoreClient interface {
}

// ProvideFirestoreClient will setup the firestore client. 
// This is not allowed to fail and the function will cause the service to exit if it does.
func ProvideFirestoreClient(config *Configuration) (FirestoreClient, error) {
	client, err := firestore.NewClient(context.Background(), config.ProjectID)
	return client, err
}