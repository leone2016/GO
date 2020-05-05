package main

import (
	"fmt"
	"strings"
	"time"
)

/*
Go rutines: son el equivalentes a los promesas, asyn - await

fmt.Scanln(&x): la << & >> signigica puntero

channels: 	son canales como su nombre lo indica, que permiten a un Go rutines envie información a otro main u otro
			go rutine,
			¿para qué nos permite enviar información?
			Para que cada Función Paralela que se este desarrollando en el procesador, pueda dialogar con otra
enviando información

 */
func main() {
	// esto trabaja de manera asyncrona
	go miNombreLento("leonardo Medina")
	fmt.Println(" Estoy aqui")
	var x string
	fmt.Scanln(&x) // al terminar de escribir la funcion se termina, sin importar si se termino de imprimir mi nombre lento

}

func miNombreLento( nombre string){
	letras := strings.Split(nombre,"")
	for _, letra := range letras {
		time.Sleep(1000*time.Millisecond)
		fmt.Println(letra)
	}
}