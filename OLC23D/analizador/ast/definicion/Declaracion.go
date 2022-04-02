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

func (dec *Declaracion) Get3D(ent *entorno.Entorno) string {

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
		CODIGO_SALIDA += temporal + " = SP + " + fmt.Sprint(posicionRelativa) + ";\n"
		CODIGO_SALIDA += "Stack[(int)" + temporal + "] = " + resultadoExpr.Temporal

		simbolo := entorno.NewSimboloIdentificador(0, 0, IDE.Identificador)
		simbolo.Direccion = posicionRelativa

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
