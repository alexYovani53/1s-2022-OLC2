package asignacion

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	arrayList "github.com/colegno/arraylist"
)

type AsignacionObjeto struct {
	ListaAccesos *arrayList.List
	Expr         interfaces.Expresion
	Valor        interface{}
}

func NewAsignacionObjeto(listaAccesos *arrayList.List, expr interfaces.Expresion) AsignacionObjeto {
	return AsignacionObjeto{ListaAccesos: listaAccesos, Expr: expr}
}

func NewAsignacionObjetoValor(listaAccesos *arrayList.List, valor interface{}) AsignacionObjeto {
	return AsignacionObjeto{ListaAccesos: listaAccesos, Valor: valor}
}

func (a AsignacionObjeto) get3D(ent *entorno.Entorno) string {

	return ""
}

func (a AsignacionObjeto) CambiarValorRecursivo(ent *entorno.Entorno, ListaAccesos *arrayList.List, OBJETO Simbolos.Objecto) entorno.Result3D {

	return entorno.Result3D{}
}
