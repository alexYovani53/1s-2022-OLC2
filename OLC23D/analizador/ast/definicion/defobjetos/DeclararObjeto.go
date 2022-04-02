package defobjetos

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"github.com/colegno/arraylist"
)

type DeclararObjeto struct {
	ValorInicializacion interfaces.Expresion
	TipoVariables       string
	ListaVars           *arraylist.List
}

func NewDeclararObjeto(tipoVariables string, listaVars *arraylist.List, valor interfaces.Expresion) DeclararObjeto {
	return DeclararObjeto{TipoVariables: tipoVariables, ListaVars: listaVars, ValorInicializacion: valor}
}

func (dec DeclararObjeto) Get3D(ent *entorno.Entorno) string {

	return ""
}
