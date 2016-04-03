package learngo

import (
	"fmt"
)

type ExampleFunc func()

type example struct {
	name string
	fn   ExampleFunc
}

var examples = []example{}

func Register(name string, fn ExampleFunc) bool {
	examples = append(examples, example{name: name, fn: fn})
	return true
}

func Run() {
	for _, e := range examples {
		fmt.Printf(">> example: %s\n", e.name)
		e.fn()
		fmt.Printf("\n")
	}
}
