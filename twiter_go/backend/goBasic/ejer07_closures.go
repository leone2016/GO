package main

import "fmt"

/*
funciones anomimas
 */

var Calculo func(int, int)int = func(num1 int,num2 int)int{
	return num1 * num2
}
func Operaciones(){
	resultado := func() int{
		var a int = 23
		var b int = 13
		return a+b
	}
	fmt.Println(resultado());
}

/*
Closures: son un concepto de funciones anonimas que tienen que ver con la protección he isolación de codigo, de mantener
nuestras variales alejadas de los curiosos
*/


func main() {
	fmt.Printf("Sumo 5 * 7 = %d \n", Calculo(5,7))

	// restamos
	Calculo = func(num1 int,num2 int)int{
		return num1 - num2
	}
	fmt.Printf("Rest 5 / 7 = %d \n", Calculo(5,7))
	// Dividir
	Calculo = func(num1 int,num2 int)int{
		return num1 /num2
	}
	fmt.Printf("Div 5 - 7 = %d \n", Calculo(5,7))
	Operaciones()

	tablaDel := 2
	fmt.Println("TABLA DE MULTIPLICAR DEL :", tablaDel)
	miTabla := Tabla(tablaDel) // miTabla se convierte en la funcion que retorna
	for i:=0; i<11; i++{
		fmt.Println(miTabla())
	}
}
/*
Closures: en mis palabras
te peremite utilizar solo la funcion embebida dentro de otra
 */
func Tabla(valor int)func()int  {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}