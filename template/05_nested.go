package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mkideal/learngo"
)

var _ = learngo.Register("05_nested", func() {
	t, err := template.New("05_nested").Parse(`
	{{define "T1"}} Hello {{.}}	{{end}}
	{{with $x := "Gopher"}}{{template "T1" $x}}{{end}}`)
	if err != nil {
		fmt.Printf("parse template 05_nested error: %v\n", err)
		return
	}
	t.Execute(os.Stdout, nil)
})
