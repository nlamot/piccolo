package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

//go:generate mockery --name CollectionRef --output ./mock/ --outpkg mock
type CollectionRef interface {
	Add(context.Context, interface{}) (DocumentRef, WriteResult, error)
}

type collectionRef struct {
	ref *firestore.CollectionRef
}

func (c *collectionRef) Add(ctx context.Context, data interface{}) (DocumentRef, WriteResult, error) {
	dr, wr, err := c.ref.Add(ctx, data)

	return documentRef{ref: dr}, writeResult{res: wr}, err
}
