package expresion

import (
	interfaces2 "OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
	"reflect"
	"strconv"
)

var suma_dominante = [5][5]entorno.TipoDato{
	{entorno.INTEGER, entorno.FLOAT, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.STRING, entorno.STRING, entorno.STRING, entorno.STRING, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var multi_division_dominante = [5][5]entorno.TipoDato{
	{entorno.INTEGER, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}
var resta_dominante = [5][5]entorno.TipoDato{
	{entorno.INTEGER, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

type Operacion struct {
	Op1      interfaces2.Expresion
	Operador string
	Op2      interfaces2.Expresion
	Unario   bool
}

func NewOperacion(Op1 interfaces2.Expresion, Operador string, Op2 interfaces2.Expresion, unario bool) Operacion {

	e := Operacion{Op1, Operador, Op2, unario}
	return e
}

func (p Operacion) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	var retornoIzq entorno.ValorType
	var retornoDer entorno.ValorType

	if p.Unario == true {
		retornoIzq = p.Op1.ObtenerValor(ent)
	} else {

		if reflect.TypeOf(p.Op1).Name() == "Identificador" {
			existeIzquierdo := ent.ExisteSimbolo(p.Op1.(Identificador).Identificador)
			if !existeIzquierdo {
				return entorno.ValorType{Tipo: entorno.NULL, Valor: nil}
			}
		}
		if reflect.TypeOf(p.Op2).Name() == "Identificador" {
			existeDerecho := ent.ExisteSimbolo(p.Op2.(Identificador).Identificador)
			if !existeDerecho {
				return entorno.ValorType{Tipo: entorno.NULL, Valor: nil}
			}
		}

		retornoIzq = p.Op1.ObtenerValor(ent)
		retornoDer = p.Op2.ObtenerValor(ent)
	}

	var dominante entorno.TipoDato

	switch p.Operador {
	case "+":
		{

			fmt.Printf("%v -> %v", retornoIzq.Tipo, retornoDer.Tipo)

			dominante = suma_dominante[retornoIzq.Tipo][retornoDer.Tipo]

			if dominante == entorno.INTEGER {

				/*

					nuevaVariable :=   variable.(instrucciones.Imprimir)

				*/

				return entorno.ValorType{Tipo: dominante, Valor: retornoIzq.Valor.(int) + retornoDer.Valor.(int)}

			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoDer.Valor), 64)
				return entorno.ValorType{Tipo: dominante, Valor: val1 + val2}

			} else if dominante == entorno.STRING {

				r1 := fmt.Sprintf("%v", retornoIzq.Valor)
				r2 := fmt.Sprintf("%v", retornoDer.Valor)

				return entorno.ValorType{Tipo: dominante, Valor: r1 + r2}
			}

		}

	case "*":
		{
			dominante = multi_division_dominante[retornoIzq.Tipo][retornoDer.Tipo]

			if dominante == entorno.INTEGER {
				return entorno.ValorType{Tipo: dominante, Valor: retornoIzq.Valor.(int) * retornoDer.Valor.(int)}

			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoDer.Valor), 64)
				return entorno.ValorType{Tipo: dominante, Valor: val1 * val2}

			} else if dominante == entorno.NULL {
				return entorno.ValorType{Tipo: dominante, Valor: nil}
			}

		}
	case "-":
		{
			if p.Unario {

				if retornoIzq.Tipo != entorno.INTEGER && retornoIzq.Tipo != entorno.FLOAT {
					return entorno.ValorType{Tipo: entorno.NULL, Valor: nil}
				}

				if retornoIzq.Tipo == entorno.INTEGER {
					return entorno.ValorType{Tipo: retornoIzq.Tipo, Valor: -1 * retornoIzq.Valor.(int)}
				} else if retornoIzq.Tipo == entorno.FLOAT {
					return entorno.ValorType{Tipo: retornoIzq.Tipo, Valor: -1 * retornoIzq.Valor.(float64)}
				}

			} else {
				dominante = resta_dominante[retornoIzq.Tipo][retornoDer.Tipo]

				if dominante == entorno.INTEGER {

					fmt.Println(retornoIzq.Tipo)
					fmt.Println(retornoDer.Tipo)

					return entorno.ValorType{Tipo: dominante, Valor: retornoIzq.Valor.(int) - retornoDer.Valor.(int)}

				} else if dominante == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoDer.Valor), 64)
					return entorno.ValorType{Tipo: dominante, Valor: val1 - val2}

				} else if dominante == entorno.NULL {
					return entorno.ValorType{Tipo: dominante, Valor: nil}
				}
			}
		}
	}

	return entorno.ValorType{Tipo: entorno.NULL, Valor: nil}
}
