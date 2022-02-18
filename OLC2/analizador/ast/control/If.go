package control

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type IfInstruccion struct {
	Condicion                   interfaces.Expresion
	ListaInstruccionesPrincipal *arrayList.List
	ListaIfElse                 *arrayList.List
	ListaInstruccionesElse      *arrayList.List

}

func NewIfInstruccion(condicion interfaces.Expresion, listaInstruccionesPrincipal *arrayList.List, listaIfElse *arrayList.List, listaInstruccionesElse *arrayList.List) IfInstruccion {
	return IfInstruccion{
		Condicion:                   condicion,
		ListaInstruccionesPrincipal: listaInstruccionesPrincipal,
		ListaIfElse:                 listaIfElse,
		ListaInstruccionesElse:      listaInstruccionesElse,
	}
}

func (i IfInstruccion) Ejecutar(ent entorno.Entorno) interface{} {

	retornoCondicionPrincipal := i.Condicion.ObtenerValor(ent)

	fmt.Printf(" \n valor %v \n", i)

	if retornoCondicionPrincipal.Tipo != entorno.BOOLEAN {
		analizador.Consola += "\ntipo no bool\n"
		return nil
	}

	if retornoCondicionPrincipal.Valor.(bool) == true {

		entornoNuevoIf := entorno.NewEntorno("IF", &ent)

		for j := 0; j < i.ListaInstruccionesPrincipal.Len(); j++ {

			instr := i.ListaInstruccionesPrincipal.GetValue(j).(interfaces.Instruccion)

			instr.Ejecutar(entornoNuevoIf)
		}

	} else {

		entornoNuevoIf := entorno.NewEntorno("IF", &ent)
		
		for j := 0; j < i.ListaInstruccionesElse.Len(); j++ {

			instr := i.ListaInstruccionesElse.GetValue(j).(interfaces.Instruccion)

			instr.Ejecutar(entornoNuevoIf)
		}

	}

	return nil
}
