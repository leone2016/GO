package user

import "time"

type Usuario struct {
	id int
	nombre string
	fechaAlta time.Time
	status bool
}
/*
apuntador
(this *Usuario)
this = Usuario
 */
func (this *Usuario) AltaUsuario( id int, nombre string, fechaalta time.Time, status bool)  {
	this.id = id
	this.nombre = nombre
	this.fechaAlta = fechaalta
	this.status = status
}