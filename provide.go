package di

import (
	"reflect"

	"go.uber.org/dig"
)

func MustProvide(container *dig.Container, fn any) {
	if kind := reflect.TypeOf(fn).Kind(); kind != reflect.Func {
		panic("mustProvide: fn must be a function")
	}

	must(container.Provide(fn))
}

func MustProvideVal[Val any](container *dig.Container, val Val) {
	if kind := reflect.TypeOf(val).Kind(); kind == reflect.Func {
		panic("mustProvideVal: val must not be a function but a plain value")
	}

	must(container.Provide(func() Val {
		return val
	}))
}
