package funbasica

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
)

type Imprimir struct {
	Expresiones interfaces.Expresion
}

func NewImprimir(val interfaces.Expresion) Imprimir {
	e := Imprimir{val}
	return e
}

func (this Imprimir) Get3D(ent *entorno.Entorno) string {

	resultadoExpr := this.Expresiones.Obtener3D(ent)

	CODIGO_SALIDA := ""

	if resultadoExpr.Tipo == entorno.NULL {
		return ""
	}

	if resultadoExpr.Tipo == entorno.INTEGER {

		CODIGO_SALIDA += "/* IMPRIMIENDO INTEGER*/\n"
		CODIGO_SALIDA += fmt.Sprintf("\"printf(\"%d\", (int)%s\"); \n", resultadoExpr.Temporal)

	} else if resultadoExpr.Tipo == entorno.FLOAT {

		CODIGO_SALIDA += "/* IMPRIMIENDO INTEGER*/\n"
		CODIGO_SALIDA += fmt.Sprintf("\"printf(\"%f\", (float)%s\"); \n", resultadoExpr.Temporal)

	} else if resultadoExpr.Tipo == entorno.STRING {

		temporal1 := analizador.GeneradorGlobal.ObtenerTemporal()

		CARACTER := analizador.GeneradorGlobal.ObtenerTemporal()
		etiquetaCiclo := analizador.GeneradorGlobal.ObtenerEtiqueta()
		etiquetaChar := analizador.GeneradorGlobal.ObtenerEtiqueta()
		etiquetaAumento := analizador.GeneradorGlobal.ObtenerEtiqueta()
		etiquetaSalida := analizador.GeneradorGlobal.ObtenerEtiqueta()

		CODIGO_SALIDA += resultadoExpr.Codigo
		CODIGO_SALIDA += fmt.Sprintf("%s = %s; /*capturando direccion en heap*/\n", temporal1, resultadoExpr.Temporal)

		CODIGO_SALIDA += fmt.Sprintf("%s: \n", etiquetaCiclo)
		CODIGO_SALIDA += fmt.Sprintf("	%s = Heap[(int)%s]; /*tomando caracter*/\n", CARACTER, temporal1)

		CODIGO_SALIDA += fmt.Sprintf("		if(%s != -1) goto %s; \n", CARACTER, etiquetaChar)
		CODIGO_SALIDA += fmt.Sprintf("			%s = %s + 1; \n", temporal1, temporal1)
		CODIGO_SALIDA += fmt.Sprintf("			%s = Heap[(int)%s]; /*tomando caracter*/\n", CARACTER, temporal1)
		CODIGO_SALIDA += "printf(\"%d\", (char) " + CARACTER + ");\n"
		CODIGO_SALIDA += fmt.Sprintf("			goto %s; \n", etiquetaAumento)

		CODIGO_SALIDA += fmt.Sprintf("		%s: \n", etiquetaChar)
		CODIGO_SALIDA += fmt.Sprintf("		if(%s == 0 ) goto %s; \n", CARACTER, etiquetaSalida)
		CODIGO_SALIDA += "printf(\"%c\", (char) " + CARACTER + ");\n"

		CODIGO_SALIDA += fmt.Sprintf("			%s: \n", etiquetaAumento)
		CODIGO_SALIDA += fmt.Sprintf("			%s = %s + 1; \n", temporal1, temporal1)

		CODIGO_SALIDA += fmt.Sprintf("			goto %s; \n", etiquetaCiclo)

		CODIGO_SALIDA += fmt.Sprintf("%s: \n", etiquetaSalida)

	}

	return CODIGO_SALIDA
}
