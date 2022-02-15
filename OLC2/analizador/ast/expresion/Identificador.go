package expresion

import (
	"OLC2/analizador/entorno"
)

type Identificador struct {
	Identificador string
}

func NewIdentificador(identificador string) Identificador {
	return Identificador{Identificador: identificador}
}

func (ide Identificador) ObtenerValor(ent entorno.Entorno) entorno.RetornoType {

	var encontrado bool = ent.ExisteSimbolo(ide.Identificador)

	if encontrado == false {
		return entorno.RetornoType{Valor: nil, Tipo: entorno.NULL}
	}

	simbo := ent.ObtenerSimbolo(ide.Identificador)

	return entorno.RetornoType{Valor: simbo.Valor, Tipo: simbo.Tipo}

}
