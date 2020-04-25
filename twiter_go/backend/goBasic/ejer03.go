package main

import "fmt"

var estado bool
func main() {
	estado=true
	if estado=false;estado {
		fmt.Println("Todo bien")
	}else{
		fmt.Println("caso contrario")
	}

	switch numero:= 5; numero {
	case 1:
		fmt.Println(" Número uno")
	case 2:
		fmt.Println(" Número dos")
	case 3:
		fmt.Println(" Número tres")
	case 4:
		fmt.Println(" Número cuatro")
	case 5:
		fmt.Println(" Número cinco")
	default:
		fmt.Println("> a 5 ")
	}
}
