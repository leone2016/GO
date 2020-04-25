package main

import "fmt"

/*
Ejercicios - Ninja Nivel 5
Ejercicio práctico #1
Crea tu propio tipo “persona” el cual tendrá un tipo subyacente tipo “struct” de manera que pueda almacenar la siguiente data:
Nombre
Apellido
Sabores de helado favoritos
Crea dos VALORES de TIPO persona. Imprime los valores, usa range sobre los elementos en el slice el cual almacena los valores
de helados favoritos.
solución:
https://play.golang.org/p/3_LXuPsT3Q3
video: 086

 */


type persona struct {
	nombre string
	apellido string
	sabores []string
}
func main() {
	x := persona{
		nombre:"leonardo",
		apellido: "medina",
		sabores:[]string{
			"mora",
			"frutos del bosque",
			"maracuya",
		},
	}

	x1 := persona{
		nombre:"Lorena",
		apellido: "G",
		sabores:[]string{
			"frutos del bosque",
			"Mora",
			"Torta Suiza",
		},
	}
	fmt.Println(x.nombre)

	for i, val := range x.sabores {
		fmt.Println("\t",i, val)
	}
	fmt.Println(x1.nombre)
	for i, val := range x1.sabores {
		fmt.Println("\t",i, val)
	}



}
