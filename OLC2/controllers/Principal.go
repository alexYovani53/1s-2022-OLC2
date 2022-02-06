package controllers

import (
	parser2 "OLC2/analizador/parser"
	"OLC2/utilidades"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"net/http"
)

type Solicitud struct {
	Text string `json:"text"`
}

func Inicio() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]interface{}{"ok": "Exitoso"})
	}
}

func Data() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		var solicitu Solicitud

		if err := json.NewDecoder(r.Body).Decode(&solicitu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Status": http.StatusBadRequest, "Message": "Error"})
			return
		}

		errores := &utilidades.CustomErrorListener{}

		//is := antlr.NewInputStream("{system.out.println( \"\nresultado: \"+( 10*1)+ \"\n\") ;system.out.println( \"\nresultado: \"+( 10*1)+ \"\n\") ;}")

		is := antlr.NewInputStream(solicitu.Text)

		// Creación de lexer
		lexer := parser2.NewCalcLexer(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errores)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Creando el parser
		p := parser2.NewCalc(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(errores)

		p.BuildParseTrees = true

		// Obteniendo la raiz
		tree := p.Start()

		fmt.Printf("\nErrores este punto %v", errores.Errors)

		// Listener para recorrer el arbol
		var listener *utilidades.TreeShapeListener = utilidades.NewTreeShapeListener()

		if len(errores.Errors) == 0 {
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}

		fmt.Println(listener.Cadena)

		json.NewEncoder(w).Encode(map[string]interface{}{"val": listener.Cadena})
	}
}
