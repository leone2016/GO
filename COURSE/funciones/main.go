package main

import "fmt"

// func (r receptor) identificador(parametros) retorno(s){ CODE }
func main() {
	foo()
	bar("LEONARDO")  // se envia argumentos
	fmt.Println( sum( 6))
	fmt.Println( saludar( "leonardo", "MEdina"))
}

func foo()  {
	fmt.Println("Hola desde foo")
}
func bar( s string){ // se definen parametros
	fmt.Println(" HOla", s)
}
func sum(val int) int{ // se definen parametros
	return val + 1
}

func saludar(n,a string)(string, bool){
	p:= fmt.Sprint(n, " ", a, " dice HOLA")
	q:= true

	return p,q
}