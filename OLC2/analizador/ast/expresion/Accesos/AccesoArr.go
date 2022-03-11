package Accesos

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type AccessoArr struct {
	Identificador string
	Dimensiones   *arrayList.List
}

func NewAccessoArr(identificador string, dimensiones *arrayList.List) AccessoArr {
	return AccessoArr{Identificador: identificador, Dimensiones: dimensiones}
}

func (a AccessoArr) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	existe := ent.ExisteSimbolo(a.Identificador)

	if !existe {
		fmt.Println("no existe el arreglo")
		return entorno.ValorType{Valor: 0, Tipo: entorno.NULL}
	}

	simbol := ent.ObtenerSimbolo(a.Identificador)
	if reflect.TypeOf(simbol) != reflect.TypeOf(Simbolos.ObjetoArray{}) {

		fmt.Println("no es un arreglo")
		return entorno.ValorType{Valor: 0, Tipo: entorno.NULL}
	}

	objetoArr := simbol.(Simbolos.ObjetoArray)

	if objetoArr.ListaIntDimensiones.Len() != a.Dimensiones.Len() {

		fmt.Println("Dimensiones invalidas")
		return entorno.ValorType{Valor: 0, Tipo: entorno.NULL}
	}

	dimensiones := a.ObtenerIntDimensiones(ent)

	val := objetoArr.ObtenerValor(dimensiones, 0, objetoArr.Valores)

	return entorno.ValorType{Valor: val, Tipo: objetoArr.Tipo}
}

func (a AccessoArr) ObtenerIntDimensiones(ent entorno.Entorno) *arrayList.List {

	listaDimensiones := arrayList.New()

	for x := 0; x < a.Dimensiones.Len(); x++ {
		dime := a.Dimensiones.GetValue(x)

		valorRet := dime.(interfaces.Expresion).ObtenerValor(ent)

		if valorRet.Tipo != entorno.INTEGER {
			return nil
		}

		listaDimensiones.Add(valorRet.Valor)
	}

	return listaDimensiones

}
