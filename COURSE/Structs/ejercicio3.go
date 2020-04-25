package main

import "fmt"

/*
Ejercicio práctico #3
Crea un nuevo tipo: vehículo.
El tipo subyacente es un struct.
Los campos:
puertas
color
Crea dos nuevos tipos: camión & sedán.
El tipo subyacente de cada uno de esos tipo es un struct.
Embebe el tipo “vehículo” en ambos camión y sedán.
Dale al camión el campo “cuatroRuedas” el cual será un booleano.
Dale al sedán el campo “lujoso” el cual será un booleano.
Con los structs vehículo, camión y sedán:
Usando un composite literal, crea un valor de tipo camión y asígnale valor a los campos.
Usando un composite literal, crea un valor de tipo sedan y asígnale valor a los campos.
Imprime cada uno de los valores.
Imprime un solo valor de cada uno de eso valores.
solución: https://play.golang.org/p/GchZPj5QDdk
video: 088

*/
type vehiculo struct {
	puertas int
	color string
}
type camion struct {
	vehiculo
	cuatroRuedas bool
}
type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	c1 := camion{
		vehiculo: vehiculo{
			puertas: 5,
			color: "rojo",
		},
		cuatroRuedas: true,
	}
	c2 := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color: "negro",
		},
		lujoso: false,
	}
	fmt.Println(c1)
	fmt.Println(c1.puertas)
	fmt.Println(c2)
	fmt.Println(c2.puertas)
}
