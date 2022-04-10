package expresion

import (
	"OLC2/analizador"
	"OLC2/analizador/entorno"
	"fmt"
	"strings"
)

type Identificador struct {
	Identificador     string
	ObtenerReferencia bool
}

func NewIdentificador(identificador string) Identificador {
	return Identificador{Identificador: identificador}
}

func (this Identificador) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	if this.ObtenerReferencia {
		return this.ObtenerDireccion(ent)

	} else {
		return this.ObtenerValor(ent)
	}

	return entorno.Result3D{}
}

func (this Identificador) ObtenerValor(ent *entorno.Entorno) entorno.Result3D {

	RESULTADO_FINAL := entorno.Result3D{}

	temporal1 := analizador.GeneradorGlobal.ObtenerTemporal()

	RESULTADO_FINAL.Codigo += fmt.Sprintf("/* BUSCANDO UN IDENTIFICADOR  >>> %s <<<*/ \n", this.Identificador)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = SP; \n", temporal1)

	ID_BUSCADO := strings.ToLower(this.Identificador)

	for entActual := ent; entActual != nil; entActual = entActual.EntAnterior {

		for key, simboloElement := range entActual.Tabla {
			if key == ID_BUSCADO {

				SIMBOLO := simboloElement.(entorno.Simbolo)

				temporal2 := analizador.GeneradorGlobal.ObtenerTemporal()
				temporal3 := analizador.GeneradorGlobal.ObtenerTemporal()

				RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = %s + %d; \n", temporal2, temporal1, SIMBOLO.Direccion)
				RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Stack[(int)%s]; \n", temporal3, temporal2)

				RESULTADO_FINAL.Codigo += fmt.Sprintf("/* IDENTIFICADOR ENCONTRADO*/\n")

				if SIMBOLO.EsReferencia {

					temporal4 := analizador.GeneradorGlobal.ObtenerTemporal()

					if SIMBOLO.Temporal_REF != "" {

						etiqueta1 := analizador.GeneradorGlobal.ObtenerEtiqueta()
						etiqueta2 := analizador.GeneradorGlobal.ObtenerEtiqueta()

						RESULTADO_FINAL.Codigo += "/* COMPROBAR UBICACION DE REFERENCIA */\n"
						RESULTADO_FINAL.Codigo += fmt.Sprintf("if( %s == 1 ) goto %s;\n", SIMBOLO.Temporal_REF, etiqueta1)
						RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Stack[(int)%s]; /* VARIABLE REFERENCIADA, VALOR FINAL*/\n", temporal4, temporal3)
						RESULTADO_FINAL.Codigo += fmt.Sprintf("goto %s;\n", etiqueta2)
						RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: /* ETIQUETA ES IGUAL A 1 --> ESTA EN HEAP*/ \n", etiqueta1)
						RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Heap[(int)%s]; /* VARIABLE REFERENCIADA, VALOR FINAL*/\n", temporal4, temporal3)
						RESULTADO_FINAL.Codigo += fmt.Sprintf("%s: /* ETIQUETA DE SALIDA*/ \n", etiqueta2)

						RESULTADO_FINAL.Temporal = temporal4
					} else {
						RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Stack[(int)%s]; /* VARIABLE REFERENCIADA, VALOR FINAL*/\n", temporal4, temporal3)
						RESULTADO_FINAL.Temporal = temporal4
					}

				} else {
					RESULTADO_FINAL.Temporal = temporal3
				}

				RESULTADO_FINAL.Tipo = SIMBOLO.Tipo
				return RESULTADO_FINAL
			}
		}

		if entActual.EntAnterior != nil {
			RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = %s - %d;\n /* RETROCEDEMOS EN EL TAMAÑO DEL ENTORNO ACTUAL*/\n", temporal1, temporal1, entActual.Tamanio)
		}

	}

	return RESULTADO_FINAL
}
func (this Identificador) ObtenerDireccion(ent *entorno.Entorno) entorno.Result3D {

	RESULTADO_FINAL := entorno.Result3D{}

	temporal1 := analizador.GeneradorGlobal.ObtenerTemporal()

	RESULTADO_FINAL.Codigo += fmt.Sprintf("/* BUSCANDO UN IDENTIFICADOR  >>> %s <<<*/ \n", this.Identificador)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = SP; \n", temporal1)

	ID_BUSCADO := strings.ToLower(this.Identificador)

	for entActual := ent; entActual != nil; entActual = entActual.EntAnterior {

		for key, simboloElement := range entActual.Tabla {
			if key == ID_BUSCADO {

				SIMBOLO := simboloElement.(entorno.Simbolo)

				temporal2 := analizador.GeneradorGlobal.ObtenerTemporal()
				RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = %s + %d; \n", temporal2, temporal1, SIMBOLO.Direccion)
				RESULTADO_FINAL.Codigo += fmt.Sprintf("/* IDENTIFICADOR ENCONTRADO*/\n")

				RESULTADO_FINAL.Temporal = temporal2
				RESULTADO_FINAL.Tipo = SIMBOLO.Tipo
				return RESULTADO_FINAL
			}
		}

		if entActual.EntAnterior != nil {
			RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = %s - %d;\n /* RETROCEDEMOS EN EL TAMAÑO DEL ENTORNO ACTUAL*/\n", temporal1, temporal1, entActual.Tamanio)
		}

	}

	return RESULTADO_FINAL
}

func (this Identificador) Obtener3DRef(ent *entorno.Entorno) entorno.Result3D {

	this.ObtenerReferencia = true
	return this.Obtener3D(ent)
}
