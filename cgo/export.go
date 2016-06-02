package main

import (
	"github.com/mkideal/learngo/cgo/lib"
	"unsafe"
)

import "C"

//export Exec
func Exec(ptr unsafe.Pointer) {
	xx := (*lib.T)(ptr)
	lib.Exec(xx)
}

func main() {
}
