package gcs

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
)

//go:generate mockery --name ObjectHandle --output ./mock/ --outpkg mock
type ObjectHandle interface {
	NewReader(context.Context) (io.Reader, error)
}

type objectHandle struct {
	objectHandle *storage.ObjectHandle
}

func (o *objectHandle) NewReader(ctx context.Context) (io.Reader, error) {
	return o.objectHandle.NewReader(ctx)
}
