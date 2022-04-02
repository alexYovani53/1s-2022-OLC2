package expresion

import (
	"OLC2/analizador"
	"OLC2/analizador/entorno"
	"fmt"
)

type Primitivo struct {
	Valor interface{}
	Tipo  entorno.TipoDato
}

func (p Primitivo) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	resultadoNuevo := entorno.Result3D{}

	if p.Tipo == entorno.BOOLEAN {

		valor := ""
		if (p.Valor.(bool) == true) {
			valor = "1"
		} else {
			valor = "0"
		}

		resultadoNuevo.Codigo = ""
		resultadoNuevo.Temporal = fmt.Sprint(valor)
		resultadoNuevo.Tipo = p.Tipo

	} else if p.Tipo == entorno.INTEGER {

		resultadoNuevo.Codigo = ""
		resultadoNuevo.Temporal = fmt.Sprint(p.Valor)
		resultadoNuevo.Tipo = p.Tipo

	} else if p.Tipo == entorno.STRING {

		temporal := analizador.GeneradorGlobal.ObtenerTemporal()

		resultadoNuevo.Temporal = temporal
		resultadoNuevo.Tipo = entorno.STRING

		resultadoNuevo.Codigo += "/* Almacenamiento de una cadena */\n"
		resultadoNuevo.Codigo += temporal + " = HP;\n"

		stringPivote := fmt.Sprint(p.Valor)

		for _, caracter := range stringPivote {

			resultadoNuevo.Codigo += "Heap[HP] = " + fmt.Sprint(caracter) + ";" + " //letra " + string(caracter) + "\n"
			resultadoNuevo.Codigo += "HP = HP + 1; \n"
		}

		resultadoNuevo.Codigo += "Heap[HP] = 0;\n "
		resultadoNuevo.Codigo += "HP = HP + 1;\n"
		resultadoNuevo.Codigo += "/*Fin de cadena*/"

		return resultadoNuevo
	}

	return resultadoNuevo
}

func NewPrimitivo(val interface{}, tipo entorno.TipoDato) Primitivo {
	e := Primitivo{val, tipo}
	return e
}
