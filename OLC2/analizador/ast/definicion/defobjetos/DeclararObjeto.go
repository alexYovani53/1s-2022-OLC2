package defobjetos

import (
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	"github.com/colegno/arraylist"
)

type DeclararObjeto struct {
	ValorInicializacion interfaces.Expresion
	TipoVariables       string
	ListaVars           *arraylist.List
}

func NewDeclararObjeto(tipoVariables string, listaVars *arraylist.List, valor interfaces.Expresion) DeclararObjeto {
	return DeclararObjeto{TipoVariables: tipoVariables, ListaVars: listaVars, ValorInicializacion: valor}
}

func (dec DeclararObjeto) Ejecutar(ent entorno.Entorno) interface{} {

	/*

		int a,b,c,d;

	*/

	if dec.ListaVars.Len() > 1 {
		return nil
	}

	retornoExpresion := dec.ValorInicializacion.ObtenerValor(ent)

	tipoExpresion := retornoExpresion.Tipo
	tipoDeclaracion := entorno.OBJETO

	if tipoDeclaracion != tipoExpresion {
		return nil
	}

	if !ent.ExisteClase(dec.TipoVariables) {
		fmt.Printf("No existe TIPO CLASE")
		return nil
	}

	OBJETO_SIMBOLO := retornoExpresion.Valor.(Simbolos.Objecto)

	if dec.TipoVariables != OBJETO_SIMBOLO.NombrePlantilla {
		return nil
	}

	for i := 0; i < dec.ListaVars.Len(); i++ {

		varDeclarar := dec.ListaVars.GetValue(i).(expresion.Identificador)

		if ent.ExisteSimbolo(varDeclarar.Identificador) {
			fmt.Printf("Errror, variable %s ya declarada", varDeclarar.Identificador)
		} else {

			ent.AgregarSimbolo(varDeclarar.Identificador, OBJETO_SIMBOLO)
		}

	}

	return nil
}
