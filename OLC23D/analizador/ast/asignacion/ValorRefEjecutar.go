package asignacion

import (
	"OLC2/analizador/entorno"
)

type ValorRefEjecutar struct {
	referencia interface{}
}

func NewValorRefEjecutar(referencia interface{}) ValorRefEjecutar {
	return ValorRefEjecutar{referencia: referencia}
}

func (a ValorRefEjecutar) get3D(ent *entorno.Entorno) string {

	return ""
}
