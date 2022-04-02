package control

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/expresion"
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

func (ifInstr IfInstruccion) Get3D(ent *entorno.Entorno) string {

	ETIQUETA_SALIDA := analizador.GeneradorGlobal.ObtenerEtiqueta()

	RESULTADO_FINAL := entorno.Result3D{}

	expresionOperacion := ifInstr.Condicion.(expresion.Operacion)
	expresionOperacion.EtiquetaVerdadera = analizador.GeneradorGlobal.ObtenerEtiqueta()
	expresionOperacion.EtiquetaFalsa = analizador.GeneradorGlobal.ObtenerEtiqueta()

	resultadoCondicion := ifInstr.Condicion.(expresion.Operacion).Obtener3D(ent)

	RESULTADO_FINAL.Codigo += "/*****************INSTRUCCIÓN IF *****************/\n"
	RESULTADO_FINAL.Codigo += resultadoCondicion.Codigo
	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: \n", resultadoCondicion.EtiquetaV)

	RESULTADO_FINAL.Codigo += ifInstr.generarCodigoInstrucciones(ifInstr.ListaInstruccionesPrincipal, ent)

	RESULTADO_FINAL.Codigo += fmt.Sprintf("goto %s;\n", ETIQUETA_SALIDA)

	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: \n", resultadoCondicion.EtiquetaF)

	if ifInstr.ListaIfElse != nil {

		for i := 0; i < ifInstr.ListaIfElse.Len(); i++ {

			IFNUEVO := ifInstr.ListaInstruccionesElse.GetValue(i).(IfInstruccion)

			NUEVO_CONDICION := IFNUEVO.Condicion.(expresion.Operacion)
			NUEVO_CONDICION.EtiquetaVerdadera = analizador.GeneradorGlobal.ObtenerEtiqueta()
			NUEVO_CONDICION.EtiquetaFalsa = analizador.GeneradorGlobal.ObtenerEtiqueta()

			resultadoNuevaCondicion := expresionOperacion.Obtener3D(ent)

			RESULTADO_FINAL.Codigo += "\n\n"
			RESULTADO_FINAL.Codigo += resultadoCondicion.Codigo

			RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: \n", resultadoNuevaCondicion.EtiquetaV)

			RESULTADO_FINAL.Codigo += ifInstr.generarCodigoInstrucciones(IFNUEVO.ListaInstruccionesPrincipal, ent)

			RESULTADO_FINAL.Codigo += fmt.Sprintf("goto %s;\n", ETIQUETA_SALIDA)
			RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: \n", resultadoNuevaCondicion.EtiquetaF)

		}

	}

	if ifInstr.ListaInstruccionesElse != nil {
		if ifInstr.ListaInstruccionesElse.Len() > 0 {

			RESULTADO_FINAL.Codigo += "/*INSTRUCCIONES ELSE */\n"
			RESULTADO_FINAL.Codigo += ifInstr.generarCodigoInstrucciones(ifInstr.ListaInstruccionesElse, ent)
		}
	}

	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: \n", ETIQUETA_SALIDA)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("/**********************FIN IF *****************/")

	return RESULTADO_FINAL.Codigo
}

func (this IfInstruccion) generarCodigoInstrucciones(lista *arrayList.List, ent *entorno.Entorno) string {

	CODIGO_SALIDA := ""

	for i := 0; i < lista.Len(); i++ {
		INSTR := lista.GetValue(i)
		CODIGO_SALIDA += INSTR.(interfaces.Instruccion).Get3D(ent)
	}

	return CODIGO_SALIDA
}
