package main

import "fmt"

/*
Ejercicio práctico #4
Crea y usa un struct anónimo.
solución: https://play.golang.org/p/mYvXb0nuc2S
video: 089

 */
func main() {
	x:= struct {
		nombre string
		apellido string
		amigos map[string]int
		bebidas []string
	}{
		nombre: "LEo",
		apellido: "Medina",
		amigos: map[string]int{
			"Andre morillo": 99722331,
			"Andre morillo1": 99722331,
			"Andre morillo2": 99722331,
			"Lorena G1": 99865346,
			"Lorena G2": 99865346 ,
			"Lorena G3": 99865346,
			"Lorena G4": 9865346,
		},
		bebidas: []string{
			"te verde",
			"te rojo",
		},
	}

	fmt.Println(x)
}
