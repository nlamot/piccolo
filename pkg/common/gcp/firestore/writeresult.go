package firestore

import "cloud.google.com/go/firestore"

//go:generate mockery --name WriteResult --output ./mock/ --outpkg mock
type WriteResult interface {
}

type writeResult struct {
	res *firestore.WriteResult
}
