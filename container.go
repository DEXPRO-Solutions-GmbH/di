package di

import "go.uber.org/dig"

func NewContainer() *dig.Container {
	container := dig.New()

	// let the container provide itself
	MustProvideVal(container, container)
	return container
}
