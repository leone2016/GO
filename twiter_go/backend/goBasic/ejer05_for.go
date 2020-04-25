package main

import "fmt"

func main() {
	j:=1
	for j<=10{
		fmt.Println(j)
		j++
	}
	numero:=0
	for{
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln( &numero)
		if numero == 100 {
			break
		}
	}
	fmt.Println(":========================:")
	var i = 0
	for i<=10{
		fmt.Printf("\n valor de $i: %d", i)
		if i==5 {
			fmt.Printf("\n\t multiplicamos por 2 \n")
			i = i * 2
			continue // no siguen las lineas de abajo
		}
		fmt.Printf("\n\tPasó por aqui")
		i++
	}
	fmt.Println(":========================:")
	var k int = 0
	RUTINA:
		fmt.Println("Ingresa")
		for k<10 {
			if k==5{
				k= k+2
				fmt.Println("voy a RUTINA")
				goto RUTINA
			}
			fmt.Printf("valor de i: %d\n", k)
			k++
		}
}
