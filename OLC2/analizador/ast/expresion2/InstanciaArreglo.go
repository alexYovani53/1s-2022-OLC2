package expresion2

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"encoding/json"
	"fmt"
	"github.com/colegno/arraylist"
)

type InstanciaArreglo struct {
	Tipo        entorno.TipoDato
	Dimensiones *arraylist.List
}

func NewInstanciaArreglo(tipo entorno.TipoDato, dimensiones *arraylist.List) InstanciaArreglo {
	return InstanciaArreglo{Tipo: tipo, Dimensiones: dimensiones}
}

func (i InstanciaArreglo) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	ListaIntDimensiones := i.ObtenerIntDimensiones(ent)

	valors := i.agregarValores(ListaIntDimensiones)

	data, err := json.MarshalIndent(valors, "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Printf(" data %v ", stringEsQuery)

	objeto := Simbolos.NewObjetoArray("", entorno.INTEGER, valors, ListaIntDimensiones)

	objetoVal := entorno.ValorType{
		Valor: objeto,
		Tipo:  entorno.INTEGER,
	}

	return entorno.ValorType{Valor: objetoVal, Tipo: entorno.ARREGLO}
}

func (i InstanciaArreglo) ObtenerIntDimensiones(ent entorno.Entorno) *arraylist.List {

	listaDimensiones := arraylist.New()

	for x := 0; x < i.Dimensiones.Len(); x++ {
		dime := i.Dimensiones.GetValue(x)

		valorRet := dime.(interfaces.Expresion).ObtenerValor(ent)

		if valorRet.Tipo != entorno.INTEGER {
			return nil
		}

		listaDimensiones.Add(valorRet.Valor)
	}

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
			s[x] = -1
		}
	}

	return s

}
