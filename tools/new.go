package main

import (
	"flag"
	"github.com/mkideal/gostart/tools/share"
)

var (
	flChapter = flag.Int("ch", 0, "Chapter number, >= 1")
	flSection = flag.Int("sec", 0, "Section number. >= 1")
)

func main() {
	flag.Parse()
}
