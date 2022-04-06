package definicion

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
	"github.com/colegno/arraylist"
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

	CambiarEntorno  bool
	TemporalEntorno string
}

func NewDeclaracion(listaVars *arraylist.List, tipoVariables entorno.TipoDato) *Declaracion {
	return &Declaracion{
		TipoVariables:   tipoVariables,
		ListaVars:       listaVars,
		CambiarEntorno:  false,
		TemporalEntorno: "",
	}
}

func NewDeclaracionParametro(listaVars *arraylist.List, tipoVariables entorno.TipoDato, referencia bool) *Declaracion {
	return &Declaracion{
		TipoVariables:   tipoVariables,
		ListaVars:       listaVars,
		Referencia:      referencia,
		CambiarEntorno:  false,
		TemporalEntorno: "",
	}
}

func NewDeclaracionInicializacion(listaVars *arraylist.List, tipoVariables entorno.TipoDato, valInicial interfaces.Expresion) *Declaracion {
	return &Declaracion{
		TipoVariables:       tipoVariables,
		ListaVars:           listaVars,
		ValorInicializacion: valInicial,
		CambiarEntorno:      false,
		TemporalEntorno:     "",
	}
}

func (dec *Declaracion) Get3D(ent *entorno.Entorno) string {

	PUNTERO_STACK_O_HEAP := "SP"
	STACK_O_HEAP := "Stack"

	if dec.CambiarEntorno {
		PUNTERO_STACK_O_HEAP = dec.TemporalEntorno
		STACK_O_HEAP = "Heap"
	}

	CODIGO_SALIDA := "\n\n\n"
	puntero_ambito := analizador.GeneradorGlobal.ObtenerTemporal()

	CODIGO_SALIDA += puntero_ambito + " = SP; \n"

	if dec.esInicializado() {

		IDE := dec.ListaVars.GetValue(0).(expresion.Identificador)

		resultadoExpr := dec.ValorInicializacion.Obtener3D(ent)

		if resultadoExpr.Tipo != dec.TipoVariables {
			return ""
		}

		CODIGO_SALIDA += resultadoExpr.Codigo

		posicionRelativa := ent.Tamanio

		temporal := analizador.GeneradorGlobal.ObtenerTemporal()

		CODIGO_SALIDA += "/* DECLARACIÓN DE VARIABLE INICIALIZADA*/\n"
		//CODIGO_SALIDA += temporal + " = SP + " + fmt.Sprint(posicionRelativa) + ";\n"

		//  TEMPORAL =  SP | HP  + ent.tamano
		CODIGO_SALIDA += fmt.Sprintf("%s = %s + %d;\n", temporal, PUNTERO_STACK_O_HEAP, posicionRelativa)

		//CODIGO_SALIDA += "Stack[(int)" + temporal + "] = " + resultadoExpr.Temporal
		CODIGO_SALIDA += fmt.Sprintf("%s[(int) %s]= %s;\n", STACK_O_HEAP, temporal, resultadoExpr.Temporal)

		simbolo := entorno.NewSimboloIdentificadorValor(0, 0, IDE.Identificador, dec.TipoVariables)
		simbolo.Direccion = posicionRelativa

		ent.AgregarSimbolo(simbolo.Identificador, simbolo)
		ent.Tamanio = ent.Tamanio + 1

	} else {

		// VARIABLE NO INICIALIZADA, VALOR POR DEFECTO

	}

	return CODIGO_SALIDA
}

func (dec *Declaracion) esInicializado() bool {
	return dec.ValorInicializacion != nil
}

func (dec *Declaracion) valorPorDefecto(tipo entorno.TipoDato) entorno.Result3D {

	resultado3D := entorno.Result3D{}

	if tipo == entorno.INTEGER {
		resultado3D.Codigo = ""
		resultado3D.Tipo = entorno.INTEGER
		resultado3D.Temporal = "0"
	}

	return resultado3D
}
