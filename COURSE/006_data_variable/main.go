package main

import (
	"log"
	"os"
	"text/template"
)
/*
consola: go run main.go
 */
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Hola soy un texto desde Golang, un Gusto`)
	if err != nil {
		log.Fatalln(err)
	}
}
