package roster_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/planning/roster"
)

var repositoryMock = new(RosterRepositoryMock)
var service = roster.ProvideRosterService(repositoryMock)

func TestCreateRosterPassesInformationToTheRepository(t *testing.T) {
	var UUID = uuid.New().String()
	repositoryMock.On("Create", aNewRoster).Return(UUID)
	assert.Equal(t, UUID, service.Create(aNewRoster))
}