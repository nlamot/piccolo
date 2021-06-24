package planner

import (
	"net/http"
	"sync"

	"go.uber.org/dig"
	"piccolo.com/planner/pkg/genetic/population"
)

var populationContainer *dig.Container
var populationContainerOnce sync.Once

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
	populationContainerOnce.Do(func() {
		populationContainer = population.NewContainer()
	})

	err := populationContainer.Invoke(func(handler population.PopulationController) {
		handler.Generate(w, r)
	})

	if err != nil {
		panic(err)
	}
}
