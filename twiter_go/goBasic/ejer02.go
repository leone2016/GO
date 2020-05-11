package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool
func main() {
	numero2, numero3, numero5,texto, status := 2,4,67,  "Este es mi texto", false

	var moneda int64 = 1991
	numero2 = int(moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero5)
	fmt.Println(texto)
	fmt.Println(status)
}
