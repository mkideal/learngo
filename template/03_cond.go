package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mkideal/learngo"
)

var _ = learngo.Register("03_cond", func() {
	type T struct {
		Map map[string]string
	}
	arg := T{
		Map: map[string]string{
			"a": "A",
			"b": "B",
			"c": "C",
		},
	}
	t, err := template.New("03_cond").
		Parse(`
{{range $k,$v := .Map}}
	{{if eq $k "a"}}		{{/*if*/}}
		{{printf "key=a"}}
	{{else if eq $k "b"}}	{{/*else if*/}}
		{{printf "val=%v where key=b" $v}}
	{{else}}				{{/*else*/}}
		what's this: {{$k}}
	{{end}}
{{end}}`)
	if err != nil {
		fmt.Printf("parse template 03_cond error: %v\n", err)
		return
	}
	t.Execute(os.Stdout, arg)
})
