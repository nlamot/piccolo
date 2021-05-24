package population_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"piccolo.com/planner/pkg/genetic/population"
)

var repositoryMock = new(PopulationRepositoryMock)
var service = population.ProvidePopulationService(repositoryMock)

func TestCreate(t *testing.T) {
	service.Create(population.CreatePopulationRequest{})
}

type PopulationServiceMock struct {
	mock.Mock
}

func (p *PopulationServiceMock) Create(r population.CreatePopulationRequest) {
	p.Called(r)
}