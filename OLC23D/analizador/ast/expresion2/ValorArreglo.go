package expresion2

import (
	"OLC2/analizador/entorno"
	arrayList "github.com/colegno/arraylist"
)

type ValorArreglo struct {
	Expresiones *arrayList.List
}

func NewValorArreglo(expresiones *arrayList.List) ValorArreglo {
	return ValorArreglo{Expresiones: expresiones}
}

func (v ValorArreglo) Obtener3D(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}

func (v ValorArreglo) ObtenerData(ent *entorno.Entorno) (interface{}, entorno.TipoDato) {
	return entorno.Result3D{}, entorno.NULL

}

func (val ValorArreglo) dimensionArreglo(dato entorno.TipoDato) bool {
	switch dato {
	case entorno.INTEGER:
	case entorno.BOOLEAN:
	case entorno.STRING:
	case entorno.FLOAT:
		return false
	}
	return true
}

func (this ValorArreglo) Obtener3DRef(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}
