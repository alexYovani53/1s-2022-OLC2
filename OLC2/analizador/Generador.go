package analizador

type Generador struct {
	temporales int
	etiquetas  int
}

func (g Generador) obtenerTemporal() string {

	temporaln := "t" + string(g.temporales)
	g.temporales++
	return temporaln
}

func (g Generador) obtenerEtiqueta() string {

	temporaln := "t" + string(g.temporales)
	g.temporales++
	return temporaln
}

func (g Generador) encabezado() string {
	encabezado := `#include <stdio.h> 

		float Heap[100000]; //Estructura heap 
		float Stack[100000]; //Estructura stack 
		int SP=0; //Puntero al stack  
		int HP=0; //Puntero al heap  
		float`

	return encabezado
}
