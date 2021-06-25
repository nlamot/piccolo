package roster

import (
	"go.uber.org/dig"
	"piccolo.com/planner/pkg/common/gcp"
	"piccolo.com/planner/pkg/common/gcp/firestore"
)

func NewContainer() *dig.Container {
	container := dig.New()
	container.Provide(gcp.GetConfiguration)
	container.Provide(firestore.ProvideFirestoreClient)
	container.Provide(ProvideRosterService)
	container.Provide(ProvideRosterRepository)
	return container
}