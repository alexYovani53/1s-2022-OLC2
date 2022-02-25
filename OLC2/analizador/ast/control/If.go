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

		if i.ListaIfElse != nil {
			for _, elseIf_instruccion := range i.ListaIfElse.ToArray() {

				nuevoIf := elseIf_instruccion.(IfInstruccion)

				retornoCondicionNuevoIf := nuevoIf.Condicion.ObtenerValor(ent)

				if retornoCondicionNuevoIf.Tipo != entorno.BOOLEAN {
					return nil
				}

				if retornoCondicionNuevoIf.Valor.(bool) {

					entornoNuevoElseIf := entorno.NewEntorno("Else if", &ent)

					for j := 0; j < nuevoIf.ListaInstruccionesPrincipal.Len(); j++ {

						instr := nuevoIf.ListaInstruccionesPrincipal.GetValue(j).(interfaces.Instruccion)

						instr.Ejecutar(entornoNuevoElseIf)
					}

					return nil
				}

			}
		}

		if i.ListaInstruccionesElse != nil {

			entornoElseFinal := entorno.NewEntorno("entorno Else final", &ent)

			for j := 0; j < i.ListaInstruccionesElse.Len(); j++ {

				instr := i.ListaInstruccionesElse.GetValue(j).(interfaces.Instruccion)

				instr.Ejecutar(entornoElseFinal)
			}

		}

	}

	return nil
}
