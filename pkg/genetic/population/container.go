package population

import (
	"go.uber.org/dig"
	"piccolo.com/planner/pkg/common/gcp"
)

func NewContainer() *dig.Container {
	container := dig.New()
	container.Provide(gcp.GetConfiguration)
	container.Provide(gcp.ProvideFirestoreClient)
	container.Provide(ProvidePopulationController)

	return container
}