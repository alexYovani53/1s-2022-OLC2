package expresion2

import (
	"OLC2/analizador/entorno"
	arrayList "github.com/colegno/arraylist"
)

type Llamada struct {
	IdFuncion        string
	ListaExpresiones *arrayList.List
}

func NewLlamada(idFuncion string, ListaExpresiones *arrayList.List) Llamada {
	return Llamada{IdFuncion: idFuncion, ListaExpresiones: ListaExpresiones}
}

func (l Llamada) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	return entorno.Result3D{}
}

func (l Llamada) Get3D(ent *entorno.Entorno) string {

	l.Obtener3D(ent)

	return ""
}
