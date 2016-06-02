package main

import (
	"fmt"
	"github.com/mkideal/learngo/cgo/lib"
	"unsafe"
)

// #include "app.h"
import "C"

type Impl struct{}

func (*Impl) Hello() string {
	fmt.Println("hello")
	return "hello"
}

func (*Impl) World() {
	fmt.Println("world")
}

func main() {
	C.load()
	v := lib.New(new(Impl))
	C.exec(unsafe.Pointer(v))
}
