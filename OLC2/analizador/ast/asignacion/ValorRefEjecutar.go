package asignacion

import (
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/ast/expresion/Accesos"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Referencia"
	"fmt"
	"reflect"
)

type ValorRefEjecutar struct {
	referencia interface{}
}

func NewValorRefEjecutar(referencia interface{}) ValorRefEjecutar {
	return ValorRefEjecutar{referencia: referencia}
}

func (v ValorRefEjecutar) Ejecutar(ent entorno.Entorno, valor interface{}) interface{} {

	valorRef := v.referencia.(Referencia.ValorRef)

	// un identificador puede ser referencia de un simbolo primitivo, objeto o arreglo

	if reflect.TypeOf(valorRef.ID) == reflect.TypeOf(expresion.Identificador{}) {

		asignacion := NewAsignacionValor(valorRef.ID.(expresion.Identificador).Identificador, valor)
		asignacion.Ejecutar(*valorRef.Entorno)

	} else if reflect.TypeOf(valorRef.ID) == reflect.TypeOf(Accesos.AccesoObjeto{}) {

		lista := valorRef.ID.(Accesos.AccesoObjeto)

		asignaObjeto := NewAsignacionObjetoValor(lista.ListaAccesos, valor)
		asignaObjeto.Ejecutar(*valorRef.Entorno)
	}

	fmt.Printf("%v", valorRef)

	return nil
}
