package gcp

import "cloud.google.com/go/firestore"

type WriteResult interface {

}

type writeResult struct {
	res *firestore.WriteResult
}