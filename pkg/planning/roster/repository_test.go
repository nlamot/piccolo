package roster_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/firestore/v1"
	"piccolo.com/planner/pkg/planning/roster"
)

var repository = roster.ProvideRosterRepository(firestoreClientMock)

var request roster.Roster
var organisationUUID string
var rosterUUID string

var firestoreWriteResult *firestore.WriteResult
var firestoreWriteError error

func TestCreateRosterNewRoster(t *testing.T) {
	givenAValidNewRosterRequest()
	givenTheFirestoreClientCreatesTheDocumentCorrectly()
	createResult, err := repository.Create(organisationUUID, request)
	assert.Nil(t, err)
	assert.Equal(t, rosterUUID, createResult)
}

func givenAValidNewRosterRequest() {
	request = aNewRoster
	organisationUUID = uuid.New().String()
}

func givenTheFirestoreClientCreatesTheDocumentCorrectly() {
	firestoreWriteResult = &firestore.WriteResult{}
	firestoreWriteError = nil
	rosterUUID = uuid.New().String()
	firestoreClientMock.On("Collection", "/organisation-data/" + organisationUUID + "/rosters").Return(collectionRefMock)
	collectionRefMock.On("Add", context.Background(), request).Return(documentRefMock, firestoreWriteResult, firestoreWriteError)
	documentRefMock.On("ID").Return(rosterUUID)
}
