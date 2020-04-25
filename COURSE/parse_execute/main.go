package main
import (
	"log"
	"os"
	"text/template"
)
func main() {
	tpl, err := template.ParseFiles("index.gohtml") // ("index.gohtml", "index2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
/*	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}*/

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
