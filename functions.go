package planner

import (
	"net/http"

	"go.uber.org/dig"
	"piccolo.com/planner/pkg/common/config"
	"piccolo.com/planner/pkg/common/db"
	"piccolo.com/planner/pkg/genetic/population"
)

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
	container := dig.New()
	container.Provide(config.GetGCPConfiguration)
	container.Provide(db.ProvideFirestoreClient)
	container.Provide(population.ProvidePopulationHandler)

	err := container.Invoke(func(handler population.PopulationHandler) {
		handler.Generate(w, r)
	})

	if err != nil {
		panic(err)
	}
}
