package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mkideal/learngo"
)

type Addr struct {
	City   string
	Street string
}

type People struct {
	Name string
	Addr Addr
}

func (p People) Say(text string) string { return p.Name + " say: " + text }

var _ = learngo.Register("01_struct", func() {
	people := People{
		Name: "Gopher",
		Addr: Addr{City: "Bei Jing", Street: "Hong Jun"},
	}

	t, err := template.New("01_struct").Parse(`{{.Say "Hello!"}}everyone!
I'm {{.Name}}. now, I'm in {{.Addr.Street}}, {{.Addr.City}}.`)
	if err != nil {
		fmt.Printf("parse template 01_struct error: %v\n", err)
		return
	}
	t.Execute(os.Stdout, people)
})
