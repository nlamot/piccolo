package planner

import (
	"net/http"

	"piccolo.com/planner/pkg/genetic/population"
)


// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
	container := population.NewContainer()

	err := container.Invoke(func(handler population.PopulationController) {
		handler.Generate(w, r)
	})

	if err != nil {
		panic(err)
	}
}
