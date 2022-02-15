package utilidades

import (
	"OLC2/analizador/ast"
	parser2 "OLC2/analizador/parser"
)

type TreeShapeListener struct {
	*parser2.BaseCalcListener
	Ast ast.Ast
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser2.StartContext) {
	this.Ast = ctx.GetAst()

	/*en := entorno.NewEntorno("hl", nil)

	for i := 0; i < result.Len(); i++ {

		r := result.GetValue(i)
		if r != nil {
			val := result.GetValue(i).(interfaces.Instruccion).Ejecutar(en)
		}

	}*/

	/*data, err := json.MarshalIndent(result.ToArray(), "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)
	fmt.Printf("%v", result)*/

}
