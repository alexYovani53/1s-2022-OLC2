package entorno

import arrayList "github.com/colegno/arraylist"

type Clase struct {
	Id            string          `json:'Id'`
	Instrucciones *arrayList.List `json:'Instruccion'`
}

func NewClase(id string, instrucciones *arrayList.List) *Clase {
	return &Clase{Id: id, Instrucciones: instrucciones}
}
