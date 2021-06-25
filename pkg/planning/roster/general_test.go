package roster_test

import "piccolo.com/planner/pkg/common/gcp/firestore/mock"
import rmock "piccolo.com/planner/pkg/planning/roster/mock"

var firestoreClientMock = new(mock.FirestoreClient)
var repositoryMock = new(rmock.RosterRepository)

var collectionRefMock = new(mock.CollectionRef)
var documentRefMock = new(mock.DocumentRef)