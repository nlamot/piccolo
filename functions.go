package planner

import (
	"net/http"
	"sync"

	"go.uber.org/dig"
	"piccolo.com/planner/pkg/genetic/population"
)

var container *dig.Container
var containerOnce sync.Once

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
	containerOnce.Do(func() {
		container = population.NewContainer()
	})

	err := container.Invoke(func(handler population.PopulationController) {
		handler.Generate(w, r)
	})

	if err != nil {
		panic(err)
	}
}
