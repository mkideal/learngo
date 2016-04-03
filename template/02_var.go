package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mkideal/learngo"
)

var _ = learngo.Register("02_var", func() {
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
	t, err := template.New("02_var").
		Parse(`{{with $x:="hello"}}say {{$x}}!{{end}}
{{range $k,$v := .Map}}{{printf "key=%v,val=%v\n" $k $v}}{{end}}`)
	if err != nil {
		fmt.Printf("parse template 02_var error: %v\n", err)
		return
	}
	t.Execute(os.Stdout, arg)
})
