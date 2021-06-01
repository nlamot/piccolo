package roster_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"piccolo.com/planner/pkg/planning/roster"
)

var firestoreClientMock = new(FirestoreClientMock)
var repository = roster.ProvideRosterRepository(firestoreClientMock)

var request roster.Roster

func TestCreateRosterNewRoster(t *testing.T) {
	givenAValidNewRoster()
	repository.Create(request)
}

func givenAValidNewRoster(){
	request = aNewRoster
}


type RosterRepositoryMock struct {
	mock.Mock
}

func (r *RosterRepositoryMock) Create(roster roster.Roster) string {
	args := r.Called(roster)
	return args.String(0)
}