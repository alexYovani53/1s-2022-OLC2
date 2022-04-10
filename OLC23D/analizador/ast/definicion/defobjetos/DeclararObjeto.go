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

// Persona  a,b = new Persana

func (dec DeclararObjeto) Get3D(ent *entorno.Entorno) string {

	if dec.ListaVars.Len() > 1 {
		return ""
	}
	expresionRetorno := dec.ValorInicializacion.Obtener3D(ent)

	if expresionRetorno.Tipo != entorno.OBJETO {
		return ""
	}

	if !ent.ExisteClase(dec.TipoVariables) {
		fmt.Printf("No existe TIPO CLASE")
		return ""
	}

	OBJETO_SIMBOLO := expresionRetorno.Valor.(Simbolos.Objecto)

	if dec.TipoVariables != OBJETO_SIMBOLO.NombrePlantilla {
		return ""
	}

	for i := 0; i < dec.ListaVars.Len(); i++ {

		varDeclarar := dec.ListaVars.GetValue(0).(expresion.Identificador)

		if ent.ExisteSimbolo(varDeclarar.Identificador) {
			fmt.Printf("Errror, variable %s ya declarada", varDeclarar.Identificador)
		} else {

			ent.AgregarSimbolo(varDeclarar.Identificador, OBJETO_SIMBOLO)
			ent.Tamanio = ent.Tamanio + 1
		}
	}

	return expresionRetorno.Codigo
}
