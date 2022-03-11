package Simbolos

import "OLC2/analizador/entorno"

type Objecto struct {
	entorno.Simbolo
	EntornoPropio   entorno.Entorno
	NombrePlantilla string
}

func NewObjecto(identificador string, entornoPropio entorno.Entorno, nombrePlantilla string) Objecto {

	simb := entorno.NewSimboloObjeto(0, 0, identificador, entorno.OBJETO)

	return Objecto{
		Simbolo:         simb,
		EntornoPropio:   entornoPropio,
		NombrePlantilla: nombrePlantilla,
	}
}
