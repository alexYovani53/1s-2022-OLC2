package analizador

import (
	"fmt"
	"strings"
)

type Generador struct {
	temporales int
	etiquetas  int
}

var GeneradorGlobal Generador = Generador{0, 0}

func (g *Generador) ObtenerTemporal() string {

	temporaln := "t" + fmt.Sprint(g.temporales)
	g.temporales = g.temporales + 1
	return temporaln
}

func (g *Generador) ObtenerEtiqueta() string {

	temporalEtiqueta := "L" + fmt.Sprint(g.etiquetas)
	g.etiquetas = g.etiquetas + 1
	return temporalEtiqueta
}

func (g *Generador) Encabezado() string {
	encabezado := `#include <stdio.h> 

		float Heap[100000]; //Estructura heap 
		float Stack[100000]; //Estructura stack
		int SP=0; //Puntero al stack 
		int HP=0; //Puntero al heap`

	if g.temporales > 0 {

		fmt.Println("llego ")

		encabezado += "\nfloat "
		for i := 0; i < g.temporales; i++ {
			if i%15 == 0 && i > 0 {
				encabezado += "\n"
			}
			encabezado += "t" + fmt.Sprint(i)
			if i < g.temporales-1 {
				encabezado += ", "
			}
		}

		encabezado += "; \n\n"
	}

	return encabezado
}

func (g *Generador) Reiniciar() {
	g.temporales = 0
	g.etiquetas = 0
}

func (g *Generador) Tabular(cadena string) string {

	stringnuevo := ""
	for _, elemento := range strings.Split(cadena, "\n") {
		stringnuevo += "\t" + elemento + "\n"
	}
	return stringnuevo
}

func (g *Generador) TabularLinea(cadena string, tabs int) string {
	stringNuevo := ""

	for i := 0; i < tabs; i++ {
		stringNuevo = "\t" + stringNuevo
	}
	return stringNuevo
}
