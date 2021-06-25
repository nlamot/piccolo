package gcs

import (
	"context"

	"cloud.google.com/go/storage"
)

// ProvideStorageClient will setup the firestore client.
// This is not allowed to fail and the function will cause the service to exit if it does.
func ProvideStorageClient() (StorageClient, error) {
	client, err := storage.NewClient(context.Background())
	return &storageClient {
		client: client, 
	}, err
}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name StorageClient -output ./mock/ -outpkg mock
type StorageClient interface {
	Bucket(string) BucketHandle
}

type storageClient struct {
	client *storage.Client
}

func (c *storageClient) Bucket(name string) BucketHandle {
	return &bucketHandle{bucket: c.client.Bucket(name)}
}
