package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// manejo de archivos en GO
func main() {
	//leoArchivo()
	leoArchivo2()
	grabarArchivo()
	grabarArchivo2()
}

/*
Lee un archivo de un solo intento, pero no deja manipular mucho
 */
func leoArchivo(){
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("hubo un error", err )
	}else{
		fmt.Println( string(archivo))
	}
}
func leoArchivo2(){
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("hubo un error", err )
	}else{
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan(){
			registro := scanner.Text()
			fmt.Printf("Linea > "+registro+"\n")
		}
	}
}

/*
grabar archivos
os.Create, primero borra el contenido del archivo y luego guarda las nuevas lineas
 */

func grabarArchivo(){
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("hubo un error", err )
		return
	}
	// empieza a guardar en el archivo
	fmt.Fprintln(archivo, "Esta es una linea nueva BB")
	archivo.Close()
}
/*
a diferencia de grabarArchivo, este mantiene las linas
o el texto anterior en que aya tenido el documentos
 */
func grabarArchivo2(){
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es una segunda Linea") == false{
		fmt.Println("Error en la 2da linea " )
	}

}
func Append(archivo string, texto string ) bool {
	arch, err := os.OpenFile(archivo,	os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("Hubo un error APPEND")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}

