// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by Calc.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcListener) ExitStart(ctx *StartContext) {}

// EnterPrueba is called when production prueba is entered.
func (s *BaseCalcListener) EnterPrueba(ctx *PruebaContext) {}

// ExitPrueba is called when production prueba is exited.
func (s *BaseCalcListener) ExitPrueba(ctx *PruebaContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseCalcListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseCalcListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseCalcListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseCalcListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BaseCalcListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BaseCalcListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterListides is called when production listides is entered.
func (s *BaseCalcListener) EnterListides(ctx *ListidesContext) {}

// ExitListides is called when production listides is exited.
func (s *BaseCalcListener) ExitListides(ctx *ListidesContext) {}

// EnterTiposvars is called when production tiposvars is entered.
func (s *BaseCalcListener) EnterTiposvars(ctx *TiposvarsContext) {}

// ExitTiposvars is called when production tiposvars is exited.
func (s *BaseCalcListener) ExitTiposvars(ctx *TiposvarsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCalcListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCalcListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpr_rel is called when production expr_rel is entered.
func (s *BaseCalcListener) EnterExpr_rel(ctx *Expr_relContext) {}

// ExitExpr_rel is called when production expr_rel is exited.
func (s *BaseCalcListener) ExitExpr_rel(ctx *Expr_relContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseCalcListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseCalcListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterExpr_valor is called when production expr_valor is entered.
func (s *BaseCalcListener) EnterExpr_valor(ctx *Expr_valorContext) {}

// ExitExpr_valor is called when production expr_valor is exited.
func (s *BaseCalcListener) ExitExpr_valor(ctx *Expr_valorContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseCalcListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseCalcListener) ExitPrimitivo(ctx *PrimitivoContext) {}
