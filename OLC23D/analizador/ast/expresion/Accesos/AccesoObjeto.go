package Accesos

import (
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	arrayList "github.com/colegno/arraylist"
)

type AccesoObjeto struct {
	ListaAccesos *arrayList.List
}

func NewAccesoObjeto(listaAccesos *arrayList.List) AccesoObjeto {
	return AccesoObjeto{ListaAccesos: listaAccesos}
}

func (a AccesoObjeto) Obtener3D(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}

func (a AccesoObjeto) ObtenerValorRecursivo(ent *entorno.Entorno, ListaAccesos *arrayList.List, OBJETO Simbolos.Objecto) entorno.Result3D {

	return entorno.Result3D{}
}
