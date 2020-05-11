package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(3)
	fmt.Println(numero,estado)

	fmt.Println(calculo(5,1,2))
}
func uno(numero int)(z int) {
	z = numero*2
	return z
}
func dos(numero int)(int,bool) {
	return numero*2,false
}
/*
PARAMETROS VARIABLES: se utiliza cuando desconosco el numero de
valores de entrada
 */

func calculo( numero ...int) int {
	total :=0
	for item, num:= range numero{
		// fmt.Printf(string(num)) no sirve imprimir aqui
		fmt.Printf("\n\t Item %d", item)
		total=total+num
	}
	return total
}