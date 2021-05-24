package roster_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"piccolo.com/planner/pkg/planning/roster"
)

var repository = roster.ProvideRosterRepository()

func TestCreateRoster(t *testing.T) {
	repository.Create(aRoster)
}

type RosterRepositoryMock struct {
	mock.Mock
}

func (r *RosterRepositoryMock) Create(roster roster.Roster) string {
	args := r.Called(roster)
	return args.String(0)
}