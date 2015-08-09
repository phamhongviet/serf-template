package main

import (
	"os"
	"text/template"
)

func RenderTemplate(src string, dest string, env interface{}) {
	// parse template
	tpl, err := template.ParseFiles(src)
	if err != nil {
		panic(err)
	}

	// render template
	result_file, err := os.Create(dest)
	defer result_file.Close()
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(result_file, env)
	if err != nil {
		panic(err)
	}
}
