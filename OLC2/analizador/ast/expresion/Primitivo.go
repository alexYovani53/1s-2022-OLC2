package expresion

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type Primitivo struct {
	Valor interface{}
	Tipo  interfaces.TipoDato
}

func (p Primitivo) ObtenerValor(entorno entorno.Entorno) interfaces.RetornoType {

	return interfaces.RetornoType{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NewPrimitivo(val interface{}, tipo interfaces.TipoDato) Primitivo {
	e := Primitivo{val, tipo}
	return e
}
