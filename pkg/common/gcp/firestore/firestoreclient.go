package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"piccolo.com/planner/pkg/common/gcp"
)

// ProvideFirestoreClient will setup the firestore client.
// This is not allowed to fail and the function will cause the service to exit if it does.
func ProvideFirestoreClient(config *gcp.Configuration) (FirestoreClient, error) {
	client, err := firestore.NewClient(context.Background(), config.ProjectID)
	return &firestoreClient{
		client: client,
	}, err
}

//go:generate mockery --name FirestoreClient --output ./mock/ --outpkg mock
type FirestoreClient interface {
	Collection(string) CollectionRef
}

type firestoreClient struct {
	client *firestore.Client
}

func (f *firestoreClient) Collection(path string) CollectionRef {
	return &collectionRef{ref: f.client.Collection(path)}
}
