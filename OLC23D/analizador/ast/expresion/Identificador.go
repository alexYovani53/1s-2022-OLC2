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

func (ide Identificador) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	return entorno.Result3D{}
}

func (ide Identificador) ObtenerReferencia(ent *entorno.Entorno) entorno.Result3D {

	return entorno.Result3D{}

}
