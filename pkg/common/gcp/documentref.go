package gcp

import "cloud.google.com/go/firestore"

type DocumentRef interface {
	ID() string
}

type documentRef struct {
	ref *firestore.DocumentRef
}

func (d documentRef) ID() string {
	return d.ref.ID
}