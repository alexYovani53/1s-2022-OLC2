package Accesos

import (
	"OLC2/analizador/entorno"
	arrayList "github.com/colegno/arraylist"
)

type AccessoArr struct {
	Identificador string
	Dimensiones   *arrayList.List
}

func NewAccessoArr(identificador string, dimensiones *arrayList.List) AccessoArr {
	return AccessoArr{Identificador: identificador, Dimensiones: dimensiones}
}

func (a AccessoArr) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	return entorno.Result3D{}
}

func (a AccessoArr) ObtenerIntDimensiones(ent *entorno.Entorno) *arrayList.List {

	return nil

}

func (a AccessoArr) Obtener3DRef(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}
