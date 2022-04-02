package transferenciaFlujo

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type Return struct {
	Tipo   entorno.TipoDato
	Salida interfaces.Expresion
}

func (r Return) Get3D(ent *entorno.Entorno) string {
	return ""
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
