package ast

import arrayList "github.com/colegno/arraylist"

type Ast struct {
	ListaInstrucciones *arrayList.List
}

func NewAst(lista *arrayList.List) Ast {
	ast := Ast{ListaInstrucciones: lista}
	return ast
}
