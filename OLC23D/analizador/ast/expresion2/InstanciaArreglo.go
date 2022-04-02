package expresion2

import (
	"OLC2/analizador/entorno"
	"github.com/colegno/arraylist"
)

type InstanciaArreglo struct {
	Tipo        entorno.TipoDato
	Dimensiones *arraylist.List
}

func NewInstanciaArreglo(tipo entorno.TipoDato, dimensiones *arraylist.List) InstanciaArreglo {
	return InstanciaArreglo{Tipo: tipo, Dimensiones: dimensiones}
}

func (i InstanciaArreglo) Obtener3D(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}

func (i InstanciaArreglo) ObtenerIntDimensiones(ent *entorno.Entorno) *arraylist.List {

	listaDimensiones := arraylist.New()
	return listaDimensiones

}

func (i InstanciaArreglo) agregarValores(list *arraylist.List) []interface{} {
	nuevaLista := list.Clone()

	anchoPiv := nuevaLista.GetValue(0).(int)
	nuevaLista.RemoveAtIndex(0)

	s := make([]interface{}, anchoPiv)

	if nuevaLista.Len() > 0 {
		for x := 0; x < anchoPiv; x++ {
			s[x] = i.agregarValores(nuevaLista)
		}
	} else {
		for x := 0; x < anchoPiv; x++ {
			s[x] = 50
		}
	}

	return s

}
