package population_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"piccolo.com/planner/pkg/genetic/population"
)

var repositoryMock = new(PopulationRepositoryMock)
var service = population.ProvidePopulationService(repositoryMock)

func TestGenerate(t *testing.T) {
	service.Generate(population.GeneratePopulationRequest{})
}

type PopulationServiceMock struct {
	mock.Mock
}

func (p *PopulationServiceMock) Generate(r population.GeneratePopulationRequest) {
	p.Called(r)
}