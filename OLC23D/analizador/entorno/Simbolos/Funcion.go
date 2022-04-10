package Simbolos

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/definicion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
	"github.com/colegno/arraylist"
	"reflect"
	"strings"
)

var retornoVal = [6][56]entorno.TipoDato{
	{entorno.INTEGER, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.VOID},
}

type Funcion struct {
	entorno.Simbolo
	ListaParamsDecl    *arraylist.List
	ListaInstrucciones *arraylist.List

	Generado bool
}

func NewFuncion(nombre string, listaParams *arraylist.List, listaInstrucciones *arraylist.List, tipo entorno.TipoDato) Funcion {
	funcSimbolo := entorno.NewSimboloFuncion(0, 0, nombre, tipo, listaParams)

	return Funcion{
		ListaInstrucciones: listaInstrucciones,
		ListaParamsDecl:    listaParams,
		Simbolo:            funcSimbolo,
	}
}

func (f Funcion) EjecutarParametros(ent *entorno.Entorno, expresiones *arraylist.List, TEMPORAL_ENTORNO string, entRef *entorno.Entorno) (bool, string) {

	CODIGO_SALIDA := ""

	declaraciones := f.ListaParamsDecl.Clone()

	if declaraciones.Len() != expresiones.Len() {
		fmt.Println("Error en variables")
		return false, ""
	}

	for i := 0; i < declaraciones.Len(); i++ {

		pivoteDec := declaraciones.GetValue(i).(*definicion.Declaracion)
		pivoteExpr := expresiones.GetValue(i).(interfaces.Expresion)

		if pivoteDec.Referencia {

			pivoteDec.EntornoRef = entRef
			pivoteDec.ValorInicializacion3D = pivoteExpr.Obtener3DRef(entRef)

		} else {
			pivoteDec.ValorInicializacion3D = pivoteExpr.Obtener3D(entRef)
		}

		pivoteDec.DecFuncion = true
		pivoteDec.TemporalEntorno = TEMPORAL_ENTORNO

		CODIGO_SALIDA += pivoteDec.Get3D(ent)
	}

	return true, CODIGO_SALIDA
}

func (f Funcion) Get3D(ent *entorno.Entorno) string {

	etiquetaRetorno := analizador.GeneradorGlobal.ObtenerEtiqueta()

	ENTORNO_FUNCION := entorno.NewEntorno("FUNCION", ent)

	CODIGO_FUNCION := ""
	CODIGO_FUNCION += fmt.Sprintf("void %s(){ \n", f.Identificador)

	f.generarParametros(&ENTORNO_FUNCION)

	for i := 0; i < f.ListaInstrucciones.Len(); i++ {

		INSTR := f.ListaInstrucciones.GetValue(i)

		codigoINSTR := INSTR.(interfaces.Instruccion).Get3D(&ENTORNO_FUNCION)

		CODIGO_FUNCION += analizador.GeneradorGlobal.Tabular(codigoINSTR)
	}

	CODIGO_FUNCION += analizador.GeneradorGlobal.TabularLinea(etiquetaRetorno+":\n", 1)
	CODIGO_FUNCION += analizador.GeneradorGlobal.TabularLinea("return\n", 1)

	PRUEBA := string(CODIGO_FUNCION)

	CODIGO_FUNCION += strings.ReplaceAll(PRUEBA, "SALIR_DE_FUNCION_:D", etiquetaRetorno)

	CODIGO_FUNCION += "}"

	return CODIGO_FUNCION
}

func (this Funcion) generarParametros(ent *entorno.Entorno) {

	ENTORNO := entorno.NewEntorno("", nil)

	ENTORNO.Tamanio = 1

	if this.ListaParamsDecl.Len() > 0 {

		analizador.GeneradorGlobal.Bloquear = true

		for i := 0; i < this.ListaParamsDecl.Len(); i++ {

			pivodeDec := this.ListaParamsDecl.GetValue(i)

			if reflect.TypeOf(pivodeDec) == reflect.TypeOf(definicion.NewDeclaracion(nil, entorno.NULL)) {

				DECLARACION := pivodeDec.(*definicion.Declaracion)

				if DECLARACION.Referencia {
					DECLARACION.Temporal_Ref = analizador.GeneradorGlobal.ObtenerTemporalForzado()
				}

				DECLARACION.Get3D(&ENTORNO)

			}
		}

		analizador.GeneradorGlobal.Bloquear = false

	}

	for ID, VARIABLE := range ENTORNO.Tabla {
		ent.AgregarSimbolo(ID, VARIABLE)
	}
	ent.Tamanio = ENTORNO.Tamanio

}

/*
fn is_even(a:u8) -> bool {
    if a % 2 == 0 {
        return true
    }
    return false // el valor de retorno en todos los otros casos
}

*/
