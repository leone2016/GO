package main

import "fmt"

func main() {
	paises := make( map[string]string)
	paises["Ecuador"] = "QUITO"
	paises["COLOMBIA"] = "BOGOTA"

	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Junior": 30,
	}
	campeonato["Chivas"]=100
	fmt.Println(campeonato)
	delete(campeonato, "Real Madrid")
	fmt.Println(paises)

	for equipo, puntaje := range campeonato{
		fmt.Printf(" El equipo de %s,\t tiene un puntaje de: %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Liga de Quito"]
	fmt.Println(puntaje, existe)


}
