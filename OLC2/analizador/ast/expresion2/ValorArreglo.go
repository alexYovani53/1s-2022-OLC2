package expresion2

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type ValorArreglo struct {
	Expresiones *arrayList.List
}

func NewValorArreglo(expresiones *arrayList.List) ValorArreglo {
	return ValorArreglo{Expresiones: expresiones}
}

func (v ValorArreglo) ObtenerValor(ent entorno.Entorno) entorno.ValorType {
	data, tipo := v.ObtenerData(ent)

	return entorno.ValorType{Valor: data, Tipo: tipo}
}

func (v ValorArreglo) ObtenerData(ent entorno.Entorno) (interface{}, entorno.TipoDato) {
	tipo := entorno.NULL
	data := arrayList.New()

	for i := 0; i < v.Expresiones.Len(); i++ {
		Expr := v.Expresiones.GetValue(i).(interfaces.Expresion)
		ValorExpr := Expr.ObtenerValor(ent)

		if i == 0 {
			tipo = ValorExpr.Tipo
			data.Add(ValorExpr)
		} else {
			if tipo != ValorExpr.Tipo {
				fmt.Println("Error1")
				return nil, entorno.NULL
			}
			data.Add(ValorExpr)
		}

	}

	ListaIntDimensiones := arrayList.New()
	ListaIntDimensiones.Add(data.Len())

	TipoDatos := entorno.NULL

	s := make([]interface{}, data.Len())

	for i := 0; i < data.Len(); i++ {

		dato := data.GetValue(i)
		valorDato := dato.(entorno.ValorType)

		if valorDato.Tipo == entorno.NULL {
			fmt.Println("Error2")
			return nil, entorno.NULL
		}
		/*
			{ 5 , 6, 7 , 8}
		*/
		if valorDato.Tipo != entorno.ARREGLO {
			if i == 0 {
				TipoDatos = valorDato.Tipo
			} else {
				if TipoDatos != valorDato.Tipo {
					fmt.Println("Error4")
					return nil, entorno.NULL
				}
			}
			s[i] = valorDato.Valor
			continue
		}

		/*
			DIMENSION ARREGLO
		*/

		valorObjeto := valorDato.Valor.(entorno.ValorType)
		ObjectoArreglo := valorObjeto.Valor.(Simbolos.ObjetoArray)

		if i == 0 {
			TipoDatos = valorObjeto.Tipo
			ListaIntDimensiones.AddAll(ObjectoArreglo.ListaIntDimensiones.ToArray())
		} else {
			if TipoDatos != valorObjeto.Tipo {

				fmt.Println("Error4")
				return nil, entorno.NULL
			}
		}

		s[i] = ObjectoArreglo.Valores

	}

	objeto := Simbolos.NewObjetoArray("", entorno.INTEGER, s, ListaIntDimensiones)
	objetoVal := entorno.ValorType{
		Valor: objeto,
		Tipo:  TipoDatos,
	}

	return objetoVal, entorno.ARREGLO

}

func (val ValorArreglo) dimensionArreglo(dato entorno.TipoDato) bool {
	switch dato {
	case entorno.INTEGER:
	case entorno.BOOLEAN:
	case entorno.STRING:
	case entorno.FLOAT:
		return false
	}
	return true
}
