package main

import (
	"flag"
	_ "github.com/mkideal/learngo/tools/share"
)

var (
	flChapter = flag.Int("ch", 0, "Chapter number, >= 1")
	flSection = flag.Int("sec", 0, "Section number. >= 1")
)

func main() {
	flag.Parse()
}
