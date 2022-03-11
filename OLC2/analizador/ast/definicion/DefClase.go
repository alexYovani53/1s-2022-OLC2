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

func NewDefClase(id string, instrucciones *arrayList.List) DefClase {

	return DefClase{Id: id, Instrucciones: instrucciones}
}

func (d DefClase) Ejecutar(ent entorno.Entorno) interface{} {

	clasePlantilla := entorno.NewClase(d.Id, d.Instrucciones)

	existeClase := ent.ExisteClase(clasePlantilla.Id)

	if !existeClase {

		ent.AgregarClase(clasePlantilla.Id, clasePlantilla)
		return nil

	}

	fmt.Printf("ERROR, LA CLASE %s YA EXISTE", clasePlantilla.Id)

	return nil
}
