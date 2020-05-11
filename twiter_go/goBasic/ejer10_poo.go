package main

import (
	us "./user"
	"fmt"
	"time"
)

type pepe struct {
	us.Usuario
}

func main() {
	persona := new(pepe)
	persona.AltaUsuario(1, "Leonardo Medina", time.Now(), false)
	fmt.Println(persona.Usuario)
}
