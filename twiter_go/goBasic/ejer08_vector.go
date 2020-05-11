package main

import "fmt"

var tabla[10] int
var matriz[5][7]int
func main() {
tabla[0]=1
tabla[9]=4
tablax :=[10]int {1,2,4,3,5,8,7,9,12}
fmt.Println(tabla)
fmt.Println(tablax)
for i:=0; i<len(tablax); i++{
	fmt.Println(tablax[i])
}

/*
matrices
 */
matriz[0][0]=1
/*
slice vectorec que puedo ampliar en tiempo de ejecucion
 */
slice := []int{2,3,4}
fmt.Println(slice)
variante2()
variante3()
variante4()
}

func variante2(){
	elementos:= [5]int{1,2,3,4,5}
	porsion := elementos[3:] // porsion se crea como slice desde el vector elementos
	fmt.Println( porsion)

}
func variante3()  {
	elementos := make([]int,5,20) //typo, largo, capacidad
	fmt.Printf("largo %d, Capacidad %d", len(elementos), cap(elementos))
}
func variante4()  {
	elementos := make([]int,0,0) //typo, largo, capacidad
	for i:=0; i<100; i++{
		elementos=append(elementos, i)
	}
	fmt.Printf("\n largo %d, Capacidad %d", len(elementos), cap(elementos))
}