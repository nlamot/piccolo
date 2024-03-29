package firestore

import "cloud.google.com/go/firestore"

//go:generate mockery --name DocumentRef --output ./mock/ --outpkg mock
type DocumentRef interface {
	ID() string
}

type documentRef struct {
	ref *firestore.DocumentRef
}

func (d documentRef) ID() string {
	return d.ref.ID
}
