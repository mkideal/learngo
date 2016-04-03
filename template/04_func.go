package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mkideal/learngo"
)

func one() string {
	return "function one"
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("y == 0")
	}
	return x / y, nil
}

var _ = learngo.Register("04_func", func() {
	for i, text := range []string{
		`{{one}} {{divide 10 0}} `,
		`{{one}} {{divide 500 2}} `,
	} {
		name := fmt.Sprintf("04_func_%d", i+1)
		t := template.New(name)
		t.Funcs(template.FuncMap{
			"one":    one,
			"divide": divide,
		})
		t, err := t.Parse(text)
		if err != nil {
			fmt.Printf("parse template %s error: %v\n", name, err)
			continue
		}
		if err := t.Execute(os.Stdout, nil); err != nil {
			fmt.Printf("Execute %q error: %v\n", text, err)
		}
	}
})
