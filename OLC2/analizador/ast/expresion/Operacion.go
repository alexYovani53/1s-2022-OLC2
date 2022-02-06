package expresion

import (
	interfaces2 "OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
)

var suma_dominante = [5][5]interfaces2.TipoDato{
	{interfaces2.INTEGER, interfaces2.FLOAT, interfaces2.STRING, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.FLOAT, interfaces2.FLOAT, interfaces2.STRING, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.STRING, interfaces2.STRING, interfaces2.STRING, interfaces2.STRING, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.STRING, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
}

var multi_division_dominante = [5][5]interfaces2.TipoDato{
	{interfaces2.INTEGER, interfaces2.FLOAT, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.FLOAT, interfaces2.FLOAT, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
}
var resta_dominante = [5][5]interfaces2.TipoDato{
	{interfaces2.INTEGER, interfaces2.FLOAT, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.FLOAT, interfaces2.FLOAT, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
	{interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL, interfaces2.NULL},
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

func (p Operacion) ObtenerValor(entorno entorno.Entorno) interfaces2.RetornoType {

	var retornoIzq interfaces2.RetornoType
	var retornoDer interfaces2.RetornoType

	if p.Unario == true {
		retornoIzq = p.Op1.ObtenerValor(entorno)
	} else {
		retornoIzq = p.Op1.ObtenerValor(entorno)
		retornoDer = p.Op2.ObtenerValor(entorno)
	}

	var dominante interfaces2.TipoDato

	switch p.Operador {
	case "+":
		{

			dominante = suma_dominante[retornoIzq.Tipo][retornoDer.Tipo]

			if dominante == interfaces2.INTEGER {

				fmt.Println(retornoIzq.Tipo)
				fmt.Println(retornoDer.Tipo)

				return interfaces2.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(int) + retornoDer.Valor.(int)}

			} else if dominante == interfaces2.FLOAT {
				return interfaces2.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(float64) + retornoDer.Valor.(float64)}

			} else if dominante == interfaces2.STRING {
				r1 := fmt.Sprintf("%v", retornoIzq.Valor)
				r2 := fmt.Sprintf("%v", retornoDer.Valor)

				return interfaces2.RetornoType{Tipo: dominante, Valor: r1 + r2}
			}

		}

	case "*":
		{
			dominante = multi_division_dominante[retornoIzq.Tipo][retornoDer.Tipo]

			if dominante == interfaces2.INTEGER {
				return interfaces2.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(int) * retornoDer.Valor.(int)}

			} else if dominante == interfaces2.FLOAT {
				return interfaces2.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(float64) * retornoDer.Valor.(float64)}

			} else if dominante == interfaces2.NULL {
				return interfaces2.RetornoType{Tipo: dominante, Valor: nil}
			}

		}
	}

	return interfaces2.RetornoType{Tipo: interfaces2.NULL, Valor: nil}
}
