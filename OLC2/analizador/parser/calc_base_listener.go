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

// EnterListaClases is called when production listaClases is entered.
func (s *BaseCalcListener) EnterListaClases(ctx *ListaClasesContext) {}

// ExitListaClases is called when production listaClases is exited.
func (s *BaseCalcListener) ExitListaClases(ctx *ListaClasesContext) {}

// EnterClases is called when production clases is entered.
func (s *BaseCalcListener) EnterClases(ctx *ClasesContext) {}

// ExitClases is called when production clases is exited.
func (s *BaseCalcListener) ExitClases(ctx *ClasesContext) {}

// EnterCuerpoClase is called when production cuerpoClase is entered.
func (s *BaseCalcListener) EnterCuerpoClase(ctx *CuerpoClaseContext) {}

// ExitCuerpoClase is called when production cuerpoClase is exited.
func (s *BaseCalcListener) ExitCuerpoClase(ctx *CuerpoClaseContext) {}

// EnterContenidoClase is called when production contenidoClase is entered.
func (s *BaseCalcListener) EnterContenidoClase(ctx *ContenidoClaseContext) {}

// ExitContenidoClase is called when production contenidoClase is exited.
func (s *BaseCalcListener) ExitContenidoClase(ctx *ContenidoClaseContext) {}

// EnterItemClase is called when production itemClase is entered.
func (s *BaseCalcListener) EnterItemClase(ctx *ItemClaseContext) {}

// ExitItemClase is called when production itemClase is exited.
func (s *BaseCalcListener) ExitItemClase(ctx *ItemClaseContext) {}

// EnterFuncionItem is called when production funcionItem is entered.
func (s *BaseCalcListener) EnterFuncionItem(ctx *FuncionItemContext) {}

// ExitFuncionItem is called when production funcionItem is exited.
func (s *BaseCalcListener) ExitFuncionItem(ctx *FuncionItemContext) {}

// EnterModaccess is called when production modaccess is entered.
func (s *BaseCalcListener) EnterModaccess(ctx *ModaccessContext) {}

// ExitModaccess is called when production modaccess is exited.
func (s *BaseCalcListener) ExitModaccess(ctx *ModaccessContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseCalcListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseCalcListener) ExitParametros(ctx *ParametrosContext) {}

// EnterParametro is called when production parametro is entered.
func (s *BaseCalcListener) EnterParametro(ctx *ParametroContext) {}

// ExitParametro is called when production parametro is exited.
func (s *BaseCalcListener) ExitParametro(ctx *ParametroContext) {}

// EnterFuncmain is called when production funcmain is entered.
func (s *BaseCalcListener) EnterFuncmain(ctx *FuncmainContext) {}

// ExitFuncmain is called when production funcmain is exited.
func (s *BaseCalcListener) ExitFuncmain(ctx *FuncmainContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseCalcListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseCalcListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseCalcListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseCalcListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterDec_objeto is called when production dec_objeto is entered.
func (s *BaseCalcListener) EnterDec_objeto(ctx *Dec_objetoContext) {}

// ExitDec_objeto is called when production dec_objeto is exited.
func (s *BaseCalcListener) ExitDec_objeto(ctx *Dec_objetoContext) {}

// EnterDec_arr is called when production dec_arr is entered.
func (s *BaseCalcListener) EnterDec_arr(ctx *Dec_arrContext) {}

// ExitDec_arr is called when production dec_arr is exited.
func (s *BaseCalcListener) ExitDec_arr(ctx *Dec_arrContext) {}

// EnterDimensiones is called when production dimensiones is entered.
func (s *BaseCalcListener) EnterDimensiones(ctx *DimensionesContext) {}

// ExitDimensiones is called when production dimensiones is exited.
func (s *BaseCalcListener) ExitDimensiones(ctx *DimensionesContext) {}

// EnterDimension is called when production dimension is entered.
func (s *BaseCalcListener) EnterDimension(ctx *DimensionContext) {}

// ExitDimension is called when production dimension is exited.
func (s *BaseCalcListener) ExitDimension(ctx *DimensionContext) {}

// EnterListanchos is called when production listanchos is entered.
func (s *BaseCalcListener) EnterListanchos(ctx *ListanchosContext) {}

// ExitListanchos is called when production listanchos is exited.
func (s *BaseCalcListener) ExitListanchos(ctx *ListanchosContext) {}

// EnterAncho is called when production ancho is entered.
func (s *BaseCalcListener) EnterAncho(ctx *AnchoContext) {}

// ExitAncho is called when production ancho is exited.
func (s *BaseCalcListener) ExitAncho(ctx *AnchoContext) {}

// EnterIf_instr is called when production if_instr is entered.
func (s *BaseCalcListener) EnterIf_instr(ctx *If_instrContext) {}

// ExitIf_instr is called when production if_instr is exited.
func (s *BaseCalcListener) ExitIf_instr(ctx *If_instrContext) {}

// EnterListaelseif is called when production listaelseif is entered.
func (s *BaseCalcListener) EnterListaelseif(ctx *ListaelseifContext) {}

// ExitListaelseif is called when production listaelseif is exited.
func (s *BaseCalcListener) ExitListaelseif(ctx *ListaelseifContext) {}

// EnterElse_if is called when production else_if is entered.
func (s *BaseCalcListener) EnterElse_if(ctx *Else_ifContext) {}

// ExitElse_if is called when production else_if is exited.
func (s *BaseCalcListener) ExitElse_if(ctx *Else_ifContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseCalcListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseCalcListener) ExitBloque(ctx *BloqueContext) {}

// EnterConsola is called when production consola is entered.
func (s *BaseCalcListener) EnterConsola(ctx *ConsolaContext) {}

// ExitConsola is called when production consola is exited.
func (s *BaseCalcListener) ExitConsola(ctx *ConsolaContext) {}

// EnterLlamada is called when production llamada is entered.
func (s *BaseCalcListener) EnterLlamada(ctx *LlamadaContext) {}

// ExitLlamada is called when production llamada is exited.
func (s *BaseCalcListener) ExitLlamada(ctx *LlamadaContext) {}

// EnterListaExpresiones is called when production listaExpresiones is entered.
func (s *BaseCalcListener) EnterListaExpresiones(ctx *ListaExpresionesContext) {}

// ExitListaExpresiones is called when production listaExpresiones is exited.
func (s *BaseCalcListener) ExitListaExpresiones(ctx *ListaExpresionesContext) {}

// EnterDeclaracionIni is called when production declaracionIni is entered.
func (s *BaseCalcListener) EnterDeclaracionIni(ctx *DeclaracionIniContext) {}

// ExitDeclaracionIni is called when production declaracionIni is exited.
func (s *BaseCalcListener) ExitDeclaracionIni(ctx *DeclaracionIniContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BaseCalcListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BaseCalcListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterRetorno is called when production retorno is entered.
func (s *BaseCalcListener) EnterRetorno(ctx *RetornoContext) {}

// ExitRetorno is called when production retorno is exited.
func (s *BaseCalcListener) ExitRetorno(ctx *RetornoContext) {}

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

// EnterArraydata is called when production arraydata is entered.
func (s *BaseCalcListener) EnterArraydata(ctx *ArraydataContext) {}

// ExitArraydata is called when production arraydata is exited.
func (s *BaseCalcListener) ExitArraydata(ctx *ArraydataContext) {}

// EnterInstancia is called when production instancia is entered.
func (s *BaseCalcListener) EnterInstancia(ctx *InstanciaContext) {}

// ExitInstancia is called when production instancia is exited.
func (s *BaseCalcListener) ExitInstancia(ctx *InstanciaContext) {}

// EnterInstanciaClase is called when production instanciaClase is entered.
func (s *BaseCalcListener) EnterInstanciaClase(ctx *InstanciaClaseContext) {}

// ExitInstanciaClase is called when production instanciaClase is exited.
func (s *BaseCalcListener) ExitInstanciaClase(ctx *InstanciaClaseContext) {}

// EnterAccesoarr is called when production accesoarr is entered.
func (s *BaseCalcListener) EnterAccesoarr(ctx *AccesoarrContext) {}

// ExitAccesoarr is called when production accesoarr is exited.
func (s *BaseCalcListener) ExitAccesoarr(ctx *AccesoarrContext) {}

// EnterAccesoObjeto is called when production accesoObjeto is entered.
func (s *BaseCalcListener) EnterAccesoObjeto(ctx *AccesoObjetoContext) {}

// ExitAccesoObjeto is called when production accesoObjeto is exited.
func (s *BaseCalcListener) ExitAccesoObjeto(ctx *AccesoObjetoContext) {}

// EnterListaAccesos is called when production listaAccesos is entered.
func (s *BaseCalcListener) EnterListaAccesos(ctx *ListaAccesosContext) {}

// ExitListaAccesos is called when production listaAccesos is exited.
func (s *BaseCalcListener) ExitListaAccesos(ctx *ListaAccesosContext) {}

// EnterAcceso is called when production acceso is entered.
func (s *BaseCalcListener) EnterAcceso(ctx *AccesoContext) {}

// ExitAcceso is called when production acceso is exited.
func (s *BaseCalcListener) ExitAcceso(ctx *AccesoContext) {}

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
