// Code generated from Calculadora.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calculadora

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalculadoraListener is a complete listener for a parse tree produced by Calculadora.
type BaseCalculadoraListener struct{}

var _ CalculadoraListener = &BaseCalculadoraListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalculadoraListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalculadoraListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalculadoraListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalculadoraListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInicio is called when production inicio is entered.
func (s *BaseCalculadoraListener) EnterInicio(ctx *InicioContext) {}

// ExitInicio is called when production inicio is exited.
func (s *BaseCalculadoraListener) ExitInicio(ctx *InicioContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseCalculadoraListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseCalculadoraListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseCalculadoraListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseCalculadoraListener) ExitPrimitivo(ctx *PrimitivoContext) {}
