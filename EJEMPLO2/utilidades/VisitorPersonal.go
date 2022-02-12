package utilidades

import (

	parser2 "EJEMPLO2/analizador/parser"
	"fmt"
	"encoding/json"

)

type VisitorPersonal  struct {
	*parser2.BaseCalculadoraListener
	Valor *ArrayList 
}

func NewVisitorPersonal() *VisitorPersonal {
	return new(VisitorPersonal)
}

func (this *VisitorPersonal) ExitInicio (ctx *parser2.InicioContext){

	this.Valor = ctx.GetLista()

	data, err := json.MarshalIndent(ctx.GetSalida(), "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)

	fmt.Println("Ingreso")

}