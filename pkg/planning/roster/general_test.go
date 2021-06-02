package roster_test

import "piccolo.com/planner/mocks"

var firestoreClientMock = new(mocks.FirestoreClient)
var repositoryMock = new(mocks.RosterRepository)
var collectionRefMock = new(mocks.CollectionRef)
var documentRefMock = new(mocks.DocumentRef)