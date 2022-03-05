package Simbolos

import (
	"OLC2/analizador/entorno"
	"fmt"
	"github.com/colegno/arraylist"
	"reflect"
)

type ObjetoArray struct {
	entorno.Simbolo
	Valores             []interface{}
	ListaIntDimensiones *arraylist.List
}

func NewObjetoArray(identificador string, Tipo entorno.TipoDato, valores []interface{}, ListaIntDimensiones *arraylist.List) ObjetoArray {

	simb := entorno.NewSimboloArreglo(0, 0, identificador, Tipo)
	return ObjetoArray{Simbolo: simb, Valores: valores, ListaIntDimensiones: ListaIntDimensiones}
}

func (Ob ObjetoArray) ObtenerValor(lista *arraylist.List, indiceNivel int, valores []interface{}) interface{} {

	listaClon := lista.Clone()

	/*
		                           0   1  2
			Ent -> Un arreglo  arr[5][6][7]

			arr[2][3][4]
	*/

	indicePiv := listaClon.GetValue(0).(int)
	tamaNivel := Ob.ListaIntDimensiones.GetValue(indiceNivel).(int)
	listaClon.RemoveAtIndex(0)

	fmt.Printf("%v -> %v", indicePiv, indiceNivel)

	if listaClon.Len() > 0 {

		if indicePiv > tamaNivel-1 {
			fmt.Println("Error1")
			return nil
		} else {

			subArray := valores[indicePiv]

			if reflect.TypeOf(subArray) != reflect.TypeOf([]interface{}{}) {

				fmt.Println("Error2")
				return nil
			} else {
				return Ob.ObtenerValor(listaClon, indiceNivel+1, subArray.([]interface{}))
			}

		}

	} else {
		if indicePiv > tamaNivel-1 {

			fmt.Println("Error3")
			return nil
		}
		return valores[indicePiv]
	}

}
