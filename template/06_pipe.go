package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/mkideal/learngo"
)

var _ = learngo.Register("06_pipe", func() {
	t := template.New("06_pipe")
	t.Funcs(template.FuncMap{
		"prepend": func(pre, s string) string { return pre + s },
		"upper":   strings.ToUpper,
	})
	t, err := t.Parse(`{{ "output" | prepend "pre-" | upper }}`)
	if err != nil {
		fmt.Printf("parse template 06_pipe error: %v\n", err)
		return
	}
	t.Execute(os.Stdout, nil)
})
