package expresion

import (
	"OLC2/analizador/entorno"
	"reflect"
)

type Identificador struct {
	Identificador string
}

func NewIdentificador(identificador string) Identificador {
	return Identificador{Identificador: identificador}
}

func (ide Identificador) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	var encontrado bool = ent.ExisteSimbolo(ide.Identificador)

	if encontrado == false {
		return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
	}

	simbo := ent.ObtenerSimbolo(ide.Identificador)

	if (reflect.TypeOf(simbo) == reflect.TypeOf(entorno.Simbolo{})) {
		dato := simbo.(entorno.Simbolo)
		return entorno.ValorType{Valor: dato.Valor, Tipo: dato.Tipo}
	}

	return entorno.ValorType{Valor: -1, Tipo: entorno.NULL}

}

func (ide Identificador) ObtenerReferencia(ent entorno.Entorno) entorno.ValorType {

	var encontrado bool = ent.ExisteSimbolo(ide.Identificador)

	if encontrado == false {
		return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
	}

	simbo := ent.ObtenerSimboloRef(ide.Identificador)

	if simbo == nil {
		return entorno.ValorType{Valor: -1, Tipo: entorno.NULL}
	}

	return entorno.ValorType{Valor: simbo, Tipo: entorno.NULL}

}
