package entorno

import (
	"strings"
)

type Entorno struct {
	Nombre      string
	EntAnterior *Entorno
	Tabla       map[string]interface{}
}

func NewEntorno(nombre string, entAnterior *Entorno) Entorno {

	en := Entorno{Nombre: nombre, EntAnterior: entAnterior}
	en.Tabla = make(map[string]interface{})
	return en
}

func (ent *Entorno) ExisteSimbolo(identificador string) bool {

	ideFinal := strings.ToLower(identificador)

	for entActual := ent; entActual != nil; entActual = entActual.EntAnterior {

		for key, _ := range entActual.Tabla {
			if key == ideFinal {
				return true
			}
		}

	}
	return false
}

func (ent *Entorno) AgregarSimbolo(identificador string, simbolo Simbolo) {
	ideFinal := strings.ToLower(identificador)

	ent.Tabla[ideFinal] = simbolo
}

func (ent *Entorno) ObtenerSimbolo(identificador string) Simbolo {

	ideFinal := strings.ToLower(identificador)

	for entActual := ent; entActual != nil; entActual = entActual.EntAnterior {

		for key, simboloElement := range entActual.Tabla {
			if key == ideFinal {
				return simboloElement.(Simbolo)
			}
		}

	}

	var simboloNil Simbolo
	return simboloNil
}

func (ent *Entorno) CambiarValor(identificador string, simboloNuevo Simbolo) {

	ideFinal := strings.ToLower(identificador)

	for entActual := ent; entActual != nil; entActual = entActual.EntAnterior {

		for key, _ := range entActual.Tabla {
			if key == ideFinal {
				ent.Tabla[ideFinal] = simboloNuevo
				return
			}
		}
	}
}
