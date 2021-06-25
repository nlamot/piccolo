package firestore

import "cloud.google.com/go/firestore"

//go:generate go run github.com/vektra/mockery/cmd/mockery -name WriteResult -output ./mock/ -outpkg mock
type WriteResult interface {

}

type writeResult struct {
	res *firestore.WriteResult
}