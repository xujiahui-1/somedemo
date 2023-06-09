package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func demo() {
	tmpl := "Hello {{.Name}}!"
	user := User{"James"}
	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
