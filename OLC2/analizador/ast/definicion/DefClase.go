package definicion

import (
	"OLC2/analizador/entorno"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type DefClase struct {
	Id            string          `json:'Id'`
	Instrucciones *arrayList.List `json:'Instruccion'`
}

type s struct {
	a int
}

func NewDefClase(id string, instrucciones *arrayList.List) DefClase {

	return DefClase{Id: id, Instrucciones: instrucciones}
}

func (d DefClase) Ejecutar(ent entorno.Entorno) interface{} {

	classPlantilla := entorno.NewClase(d.Id, d.Instrucciones)

	existeClase := ent.ExisteClase(classPlantilla.Id)

	if !existeClase {
		ent.AgregarClase(classPlantilla.Id, classPlantilla)
		return nil
	}

	fmt.Printf("ERROR, LA CLASE %s YA EXISTE", classPlantilla.Id)

	return nil
}
