package Referencia

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type ValorRef struct {
	Expresion   interfaces.Expresion
	ValorNuevo  interfaces.Expresion
	ValorObjeto interface{}
}

func NewValorRef(expresion interfaces.Expresion) ValorRef {
	return ValorRef{Expresion: expresion, ValorNuevo: nil, ValorObjeto: nil}
}

func (v ValorRef) Ejecutar(ent entorno.Entorno) interface{} {

	return nil
}
