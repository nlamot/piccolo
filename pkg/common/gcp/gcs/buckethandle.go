package gcs

import (
	"cloud.google.com/go/storage"
)

//go:generate mockery --name BucketHandle --output ./mock/ --outpkg mock
type BucketHandle interface {
	Object(string) ObjectHandle
}

type bucketHandle struct {
	bucket *storage.BucketHandle
}

func (b *bucketHandle) Object(name string) ObjectHandle {
	return &objectHandle{b.bucket.Object(name)}
}
