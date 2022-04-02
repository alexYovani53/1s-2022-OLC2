package expresion2

import (
	"OLC2/analizador/entorno"
)

type InstanciaObjeto struct {
	Id string
}

func NewInstanciaObjeto(id string) InstanciaObjeto {
	return InstanciaObjeto{Id: id}
}

func (i InstanciaObjeto) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	return entorno.Result3D{}
}
