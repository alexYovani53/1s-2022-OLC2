package transferenciaFlujo

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type Return struct {
	Tipo   entorno.TipoDato
	Salida interfaces.Expresion
}

func (r Return) Ejecutar(ent entorno.Entorno) interface{} {

	if r.Tipo == entorno.VOID {
		return entorno.ValorType{Tipo: entorno.VOID, Valor: 0}
	}

	retExpresion := r.Salida.ObtenerValor(ent)

	return retExpresion

}

func NewReturn(tipo entorno.TipoDato, salida interfaces.Expresion) Return {

	if salida != nil {
		return Return{
			Tipo:   tipo,
			Salida: salida,
		}
	}

	return Return{
		Tipo:   tipo,
		Salida: nil,
	}
}
