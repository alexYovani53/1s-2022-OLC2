package definicion

import (
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Referencia"
	"fmt"
	"github.com/colegno/arraylist"
	"reflect"
)

var tipoDef = [5][5]entorno.TipoDato{
	{entorno.INTEGER, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

type Declaracion struct {
	ValorInicializacion interfaces.Expresion
	TipoVariables       entorno.TipoDato
	ListaVars           *arraylist.List
	Referencia          bool
	EntornoRef          *entorno.Entorno
}

func NewDeclaracion(listaVars *arraylist.List, tipoVariables entorno.TipoDato) *Declaracion {
	return &Declaracion{
		TipoVariables: tipoVariables,
		ListaVars:     listaVars,
	}
}

func NewDeclaracionParametro(listaVars *arraylist.List, tipoVariables entorno.TipoDato, referencia bool) *Declaracion {
	return &Declaracion{
		TipoVariables: tipoVariables,
		ListaVars:     listaVars,
		Referencia:    referencia,
	}
}

func NewDeclaracionInicializacion(listaVars *arraylist.List, tipoVariables entorno.TipoDato, valInicial interfaces.Expresion) *Declaracion {
	return &Declaracion{
		TipoVariables:       tipoVariables,
		ListaVars:           listaVars,
		ValorInicializacion: valInicial,
	}
}

func (dec *Declaracion) Ejecutar(ent entorno.Entorno) interface{} {

	/*

		int a,b,c,d;

	*/
	if dec.esInicializado() {

		if dec.ListaVars.Len() > 1 {
			return nil
		}

		retornoExpresion := dec.ValorInicializacion.ObtenerValor(ent)

		tipoExpresion := retornoExpresion.Tipo
		tipoDeclaracion := dec.TipoVariables

		tipoResultante := tipoDef[tipoDeclaracion][tipoExpresion]

		if tipoResultante == entorno.NULL {
			return nil
		}

		for i := 0; i < dec.ListaVars.Len(); i++ {

			varDeclarar := dec.ListaVars.GetValue(i).(expresion.Identificador)

			if ent.ExisteSimbolo(varDeclarar.Identificador) {
				fmt.Printf("Errror, variable %s ya declarada", varDeclarar.Identificador)
			} else {

				simboloTabala := entorno.NewSimboloIdentificadorValor(
					0,
					0,
					varDeclarar.Identificador,
					retornoExpresion.Valor,
					tipoResultante)

				if dec.Referencia {

					if reflect.TypeOf(dec.ValorInicializacion) == reflect.TypeOf(expresion.Primitivo{}) {
						fmt.Printf("No se puede pasar una referencia a un primitivo")
						return nil
					}

					//      llama( ID,  ID.ID , ID[0] )
					//
					//      ID = VALOR
					// 		ID.ID = VALOR
					// 		ID[0] = VALOR

					simboloTabala.EsReferencia = true
					simboloTabala.Referencia = Referencia.ValorRef{
						Entorno: dec.EntornoRef,
						ID:      dec.ValorInicializacion,
					}
				}

				ent.AgregarSimbolo(varDeclarar.Identificador, simboloTabala)
			}

		}

	}

	return nil
}

func (dec *Declaracion) esInicializado() bool {
	return dec.ValorInicializacion != nil
}
