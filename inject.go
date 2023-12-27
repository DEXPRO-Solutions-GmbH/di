package di

import "go.uber.org/dig"

func Inject[T any](container *dig.Container) (T, error) {
	var value T
	err := container.Invoke(func(t T) {
		value = t
	})
	return value, err
}

func MustInject[T any](container *dig.Container) T {
	value, err := Inject[T](container)
	must(err)
	return value
}
