package main

import (

	parse2	"EJEMPLO2/analizador/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"EJEMPLO2/utilidades"
	"fmt"
)

func main(){

	errores := &utilidades.CustomErrorListener{}
	
	// lexer
	entrada := antlr.NewInputStream("4+5 |")

	lexer := parse2.NewCalLexer(entrada)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errores)

	streamTokens := antlr.NewCommonTokenStream(lexer,antlr.TokenDefaultChannel)



	// parser
	erroresParser := &utilidades.CustomErrorListener{}
	pars := parse2.NewCalculadora(streamTokens)
	pars.RemoveErrorListeners()
	pars.AddErrorListener(erroresParser)


	pars.BuildParseTrees = true

	arbol := pars.Inicio()

	var visitor *utilidades.VisitorPersonal = utilidades.NewVisitorPersonal()

	antlr.ParseTreeWalkerDefault.Walk(visitor,arbol)


	fmt.Printf("%v",errores.Errors)
	fmt.Printf("%v",erroresParser.Errors)
	fmt.Printf("El valor es %v: ",visitor.Valor)

}