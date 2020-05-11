package main

import (
	"fmt"
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

un canal es un espacio de memoria de dialogo, en el que van a dialogar distintas rutinas,
y cuando se aloje un valor en ese espacio de memoria la rutina que esta pidiendo un valor
a cambio va actuar en consecuencia.

Ese espacio de memoria es importante entenderlo, ya que debe ser de un tipo de dato, string, int, en
difinitiva lo que va trasportar nuestro canal es una variable(tipo de dato)

*/
func main() {
	canal1 := make(chan time.Duration)
	go bucle( canal1 )
	fmt.Println("Llegué hasta aquí")

	// sintaxis para decir que estoy esperando a canal 1
	msg := <- canal1 // esto es similar al await
	fmt.Println(msg)

}

func bucle(canal chan time.Duration){
	inicio := time.Now()
	for i:=0; i<100000000000;i++{

	}
	final := time.Now()
	//almacena en espacio de memoria la diferencia
	canal <- final.Sub(inicio)
}
