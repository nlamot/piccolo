package smartschool

import (
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()
	container.Provide(ProvideRosterManager)
	return container
}