package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mkideal/learngo/tools/share"
	"github.com/mkideal/log"
)

func main() {
	flag.Parse()

	log.SetLevel(share.LogLevel())

	book := share.Parse(share.Chapters()...)
	log.WithJSON(book).Debug("book")
	err := book.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
