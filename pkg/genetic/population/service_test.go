package population_test

import (
	"testing"

	"piccolo.com/planner/pkg/genetic/population"
)

var service = population.ProvidePopulationService(repositoryMock)

func TestGenerate(t *testing.T) {
	service.Generate(population.GeneratePopulationRequest{})
}
