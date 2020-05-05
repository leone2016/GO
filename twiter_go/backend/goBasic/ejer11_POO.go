package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type aninal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

//Genero Humano
type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre	bool
	vivo bool
}

type mujer struct{
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer() { h.comiendo = true }
func (h *hombre) pensar() { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre { return "hombre" }else{ return "Mujer"}
}
func (h *hombre) estaVivo() bool{  return h.vivo }
/*
ingresa una interface tipo humano, atravez de este se puede acceder a  hombre o en donde se
haya implementado

 */
func HumanosRespirando( hu humano){
	hu.respirar()
	fmt.Printf("\f Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}
/*
ANIMALES
 */
type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}
func (p *perro) respirar() { p.respirando = true }
func (p *perro) comer() { p.comiendo = true }
func (p *perro) esCarnivoro()  bool {return p.carnivoro }
func (p *perro) estaVivo() bool{  return p.vivo }
func AnimalesRespirar( an aninal ){
	an.respirar();
	fmt.Printf( "\f Soy un animal y estoy respirando\n ")
}
/*
Ser vivo
en esta interface se puede enviar, una persona o un perro.

esto es una diferencia de otros lenguajes de programacion
 */

func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}
func AnimalEsCarnivoro(an aninal) int {
	if an.esCarnivoro() {
		return 1
	}
	return 0
}

func main() {
	fmt.Println("HUMANOS")

	Pedro := new(hombre)
	Pedro.esHombre = true
	Pedro.vivo= true

	HumanosRespirando(Pedro)
	fmt.Printf( "\t Estoy vivo = %t \n ", estoyVivo(Pedro))

	Maria := new(hombre)
	Maria.vivo= false
	HumanosRespirando(Maria)
	fmt.Printf( "\t Estoy vivo = %t \n ", estoyVivo(Maria))

	// ANIMAL TIPO PERRO
// ANIMAL TIPO PERRO
// ANIMAL TIPO PERRO
fmt.Println("PERRO")
	totalCanivoros := 0
	dogo := new(perro)
	dogo.carnivoro = true
	dogo.vivo = true
	AnimalesRespirar(dogo)
	totalCanivoros =+ AnimalEsCarnivoro(dogo)

	fmt.Printf( "\ttotal carnivoros %d \n", totalCanivoros)
	fmt.Printf( "\t Estoy vivo = %t \n ", estoyVivo(dogo))


}
