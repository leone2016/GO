package main

import "fmt"

/*
Ejercicio práctico #2
Usa el código del ejemplo anterior y almacena los valores de tipo persona en un mapa con la llave apellido.
Accede a cada valor en el mapa. Imprime los valores. Imprime también los valores haciendo range sobre el slice.
solution: https://play.golang.org/p/h5Uf4CwZt2S
video: 087


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
		apellido: "Guartazaca",
		sabores:[]string{
			"frutos del bosque",
			"Mora",
			"Torta Suiza",
		},
	}
	m := map[string]persona{
		x.apellido: x,
		x1.apellido: x1,
	}
	fmt.Println(m)
	for _, v := range m{
		fmt.Println(v.nombre, v.apellido)
		for k, val:=range v.sabores{
			fmt.Println( "\t",k, val)
		}
	}




}
