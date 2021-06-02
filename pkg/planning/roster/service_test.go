package roster_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/planning/roster"
)

var service = roster.ProvideRosterService(repositoryMock)

func TestCreateRosterPassesInformationToTheRepository(t *testing.T) {
	var UUID = uuid.New().String()
	var organisationUUID = uuid.New().String()
	repositoryMock.On("Create", organisationUUID, aNewRoster).Return(UUID, nil)
	var result, err = service.Create(organisationUUID, aNewRoster)
	assert.Equal(t, UUID, result)
	assert.Nil(t, err)
}