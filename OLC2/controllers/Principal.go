package controllers

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/definicion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	parser2 "OLC2/analizador/parser"
	"OLC2/utilidades"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"net/http"
	"reflect"
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

		analizador.Consola = ""
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

		AST := listener.Ast

		ENTORNO_GLOBAL := entorno.NewEntorno("GLOBAL", nil)

		for i := 0; i < AST.ListaInstrucciones.Len(); i++ {

			r := AST.ListaInstrucciones.GetValue(i)
			if r != nil {
				if reflect.TypeOf(r) == reflect.TypeOf(definicion.DefClase{}) {
					def_CLASE := r.(definicion.DefClase)
					def_CLASE.Ejecutar(ENTORNO_GLOBAL)
				}
			}

		}

		salir := false

		for _, element := range ENTORNO_GLOBAL.TablaClases {

			pivotClass := element.(*entorno.Clase)

			if salir {
				break
			}

			for j := 0; j < pivotClass.Instrucciones.Len(); j++ {

				buscar := pivotClass.Instrucciones.GetValue(j)

				if (reflect.TypeOf(buscar) == reflect.TypeOf(Simbolos.Funcion{})) {

					if buscar.(Simbolos.Funcion).Identificador == "main" {

						CREAR_MAIN(*pivotClass, ENTORNO_GLOBAL)

						MAIN_FN := buscar.(Simbolos.Funcion)
						MAIN_FN.Ejecutar(ENTORNO_GLOBAL)

						salir = true
						break
					}
				}
			}

		}

		json.NewEncoder(w).Encode(map[string]interface{}{"val": analizador.Consola})
	}
}

func CREAR_MAIN(MAIN entorno.Clase, ent entorno.Entorno) {
	for x := 0; x < MAIN.Instrucciones.Len(); x++ {

		r := MAIN.Instrucciones.GetValue(x)

		if r != nil {
			if reflect.TypeOf(r) == reflect.TypeOf(Simbolos.Funcion{}) {
				func_ := r.(Simbolos.Funcion)

				if !ent.ExisteFuncion(func_.Identificador) {
					ent.AgregarFuncion(func_.Identificador, func_)
				} else {
					//ERROR
				}

			} else {
				r.(interfaces.Instruccion).Ejecutar(ent)
			}

		}

	}

}