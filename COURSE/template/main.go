package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
/*
at the terminal:
go run main.go Leonardo
*/
func main() {
	fmt.Println("ñ",os.Args[0])
	fmt.Println(os.Args[0])
	fmt.Println("----",os.Args[1])
	fmt.Println("sssss",os.Args[1])
	fmt.Println(os.Args[1])
	name := os.Args[1]

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}
