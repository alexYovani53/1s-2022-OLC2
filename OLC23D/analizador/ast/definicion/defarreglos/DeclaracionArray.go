package defarreglos

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type DeclaracionArray struct {
	Tamano        int
	Identificador string
	Inicializar   interfaces.Expresion
	Tipo          entorno.TipoDato
}

func NewDeclaracionArray(tamano int, identificador string, inicializar interfaces.Expresion, tipo entorno.TipoDato) DeclaracionArray {
	return DeclaracionArray{
		Tamano:        tamano,
		Identificador: identificador,
		Inicializar:   inicializar,
		Tipo:          tipo,
	}
}

func (d DeclaracionArray) Get3D(ent *entorno.Entorno) string {
	return ""

}
