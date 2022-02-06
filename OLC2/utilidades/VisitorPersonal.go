package utilidades

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	parser2 "OLC2/analizador/parser"
	"fmt"
)

type TreeShapeListener struct {
	*parser2.BaseCalcListener
	Cadena string
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser2.StartContext) {
	result := ctx.GetLista()
	for i := 0; i < result.Len(); i++ {

		r := result.GetValue(i)
		if r != nil {
			val := result.GetValue(i).(interfaces.Instruccion).Ejecutar(entorno.NewEntorno("hl", nil))
			this.Cadena += fmt.Sprintf("%v", val)
		}

	}

	/*data, err := json.MarshalIndent(result.ToArray(), "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)
	fmt.Printf("%v", result)*/

}
