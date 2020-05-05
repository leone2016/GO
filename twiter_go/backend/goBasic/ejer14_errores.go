package main

import (
	"fmt"
	"log"
	"os"
)

/*
* DEFER: es una instrucción que se va ejecutar si o si, cuando se va por un return, error ó fin de funcion
* os.Exit(1): finaliza la ejecución del programa
* Cada vez que se hace un open es necesario cerrar, para no ocupar buffer de memoria
* Panic: es una funcion de go para forzar un error, de esa manera el sistema va abortar y muestra el texto que
se le ingrese, esto funciona cuando se necesita un dato que es importante.
* Recover: Para poder controlar el el Panic y va tomar el control del mensaje, en el caso que no exista ningun panic, retorna nil

 */
func main() {
	archivo := "prueba.txt"
	ejemploPanic()
	f,err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}



}
/*
para utilizar el defer con mas de una acción es necesario crear una función anonima

 */
func ejemploPanic(){
	defer func() {
		reco := recover()
		if reco != nil{
			log.Fatalf("ocurrio un error que generó Panic \n %v", reco)
		}
		return
	}()
	a:= 1
	if a == 1{
		panic("Se encontro el el valor 1")
	}

}