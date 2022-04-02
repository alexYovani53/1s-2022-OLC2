// Generated from C:/Users/Yovani/Desktop/OLC2/analizador\Calc.g4 by ANTLR 4.9.2


    import "OLC2/analizador/ast"
    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/funbasica"
    import "OLC2/analizador/ast/definicion"
    import "OLC2/analizador/ast/control"
    import "OLC2/analizador/entorno"
    import "OLC2/analizador/entorno/Simbolos"
    import arrayList "github.com/colegno/arraylist"

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link Calc}.
 */
public interface CalcListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link Calc#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(Calc.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(Calc.StartContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#listaFunciones}.
	 * @param ctx the parse tree
	 */
	void enterListaFunciones(Calc.ListaFuncionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listaFunciones}.
	 * @param ctx the parse tree
	 */
	void exitListaFunciones(Calc.ListaFuncionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#funcionItem}.
	 * @param ctx the parse tree
	 */
	void enterFuncionItem(Calc.FuncionItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#funcionItem}.
	 * @param ctx the parse tree
	 */
	void exitFuncionItem(Calc.FuncionItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#modaccess}.
	 * @param ctx the parse tree
	 */
	void enterModaccess(Calc.ModaccessContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#modaccess}.
	 * @param ctx the parse tree
	 */
	void exitModaccess(Calc.ModaccessContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(Calc.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(Calc.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#funcmain}.
	 * @param ctx the parse tree
	 */
	void enterFuncmain(Calc.FuncmainContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#funcmain}.
	 * @param ctx the parse tree
	 */
	void exitFuncmain(Calc.FuncmainContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterInstrucciones(Calc.InstruccionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitInstrucciones(Calc.InstruccionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#instruccion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion(Calc.InstruccionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#instruccion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion(Calc.InstruccionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#if_instr}.
	 * @param ctx the parse tree
	 */
	void enterIf_instr(Calc.If_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#if_instr}.
	 * @param ctx the parse tree
	 */
	void exitIf_instr(Calc.If_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#listaelseif}.
	 * @param ctx the parse tree
	 */
	void enterListaelseif(Calc.ListaelseifContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listaelseif}.
	 * @param ctx the parse tree
	 */
	void exitListaelseif(Calc.ListaelseifContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#else_if}.
	 * @param ctx the parse tree
	 */
	void enterElse_if(Calc.Else_ifContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#else_if}.
	 * @param ctx the parse tree
	 */
	void exitElse_if(Calc.Else_ifContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#bloque}.
	 * @param ctx the parse tree
	 */
	void enterBloque(Calc.BloqueContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#bloque}.
	 * @param ctx the parse tree
	 */
	void exitBloque(Calc.BloqueContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#consola}.
	 * @param ctx the parse tree
	 */
	void enterConsola(Calc.ConsolaContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#consola}.
	 * @param ctx the parse tree
	 */
	void exitConsola(Calc.ConsolaContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#llamada}.
	 * @param ctx the parse tree
	 */
	void enterLlamada(Calc.LlamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#llamada}.
	 * @param ctx the parse tree
	 */
	void exitLlamada(Calc.LlamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#listaExpresiones}.
	 * @param ctx the parse tree
	 */
	void enterListaExpresiones(Calc.ListaExpresionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listaExpresiones}.
	 * @param ctx the parse tree
	 */
	void exitListaExpresiones(Calc.ListaExpresionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#declaracionIni}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionIni(Calc.DeclaracionIniContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#declaracionIni}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionIni(Calc.DeclaracionIniContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion(Calc.DeclaracionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion(Calc.DeclaracionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#listides}.
	 * @param ctx the parse tree
	 */
	void enterListides(Calc.ListidesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listides}.
	 * @param ctx the parse tree
	 */
	void exitListides(Calc.ListidesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#tiposvars}.
	 * @param ctx the parse tree
	 */
	void enterTiposvars(Calc.TiposvarsContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#tiposvars}.
	 * @param ctx the parse tree
	 */
	void exitTiposvars(Calc.TiposvarsContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(Calc.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(Calc.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#expr_rel}.
	 * @param ctx the parse tree
	 */
	void enterExpr_rel(Calc.Expr_relContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#expr_rel}.
	 * @param ctx the parse tree
	 */
	void exitExpr_rel(Calc.Expr_relContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#expr_arit}.
	 * @param ctx the parse tree
	 */
	void enterExpr_arit(Calc.Expr_aritContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#expr_arit}.
	 * @param ctx the parse tree
	 */
	void exitExpr_arit(Calc.Expr_aritContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#expr_valor}.
	 * @param ctx the parse tree
	 */
	void enterExpr_valor(Calc.Expr_valorContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#expr_valor}.
	 * @param ctx the parse tree
	 */
	void exitExpr_valor(Calc.Expr_valorContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#primitivo}.
	 * @param ctx the parse tree
	 */
	void enterPrimitivo(Calc.PrimitivoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#primitivo}.
	 * @param ctx the parse tree
	 */
	void exitPrimitivo(Calc.PrimitivoContext ctx);
}