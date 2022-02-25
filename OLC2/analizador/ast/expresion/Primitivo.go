package expresion

import (
	"OLC2/analizador/entorno"
)

type Primitivo struct {
	Valor interface{}
	Tipo  entorno.TipoDato
}

func (p Primitivo) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	return entorno.ValorType{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NewPrimitivo(val interface{}, tipo entorno.TipoDato) Primitivo {
	e := Primitivo{val, tipo}
	return e
}
