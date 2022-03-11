package defarreglos

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	"reflect"
)

type DeclaracionArray struct {
	Tamano        int
	Identificador string
	Inicializar   interfaces.Expresion
	Tipo          entorno.TipoDato
}

func NewDeclaracionArray(tamano int, identificador string, inicializar interfaces.Expresion, tipo entorno.TipoDato) DeclaracionArray {
	return DeclaracionArray{
		Tamano:        tamano,
		Identificador: identificador,
		Inicializar:   inicializar,
		Tipo:          tipo,
	}
}

func (d DeclaracionArray) Ejecutar(ent entorno.Entorno) interface{} {

	valorDec := d.Inicializar.ObtenerValor(ent)

	if valorDec.Tipo != entorno.VOID {
		fmt.Printf("%v ", valorDec)
	}

	if valorDec.Tipo != entorno.ARREGLO {
		fmt.Println("Err1")
		return nil
	}

	if reflect.TypeOf(valorDec.Valor) != reflect.TypeOf(entorno.ValorType{}) {
		fmt.Println("Err2")
		return nil
	}

	Objeto := valorDec.Valor.(entorno.ValorType)

	if Objeto.Tipo != d.Tipo {

		fmt.Println("Err3")
		return nil
	}

	if Objeto.Valor.(Simbolos.ObjetoArray).ListaIntDimensiones.Len() > d.Tamano {

		fmt.Printf("Errror, variable %s no posible declarar", d.Identificador)
		return nil
	}

	if ent.ExisteSimbolo(d.Identificador) {
		fmt.Printf("Errror, variable %s ya declarada", d.Identificador)
	} else {

		simbol := Objeto.Valor.(Simbolos.ObjetoArray)
		simbol.Identificador = d.Identificador

		ent.AgregarSimbolo(d.Identificador, simbol)
	}

	fmt.Printf("%v ", ent.Tabla)

	return nil

}
