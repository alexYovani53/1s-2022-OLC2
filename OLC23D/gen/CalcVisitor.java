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

import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link Calc}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface CalcVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link Calc#start}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStart(Calc.StartContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#listaFunciones}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitListaFunciones(Calc.ListaFuncionesContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#funcionItem}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFuncionItem(Calc.FuncionItemContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#modaccess}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitModaccess(Calc.ModaccessContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#parametros}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParametros(Calc.ParametrosContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#funcmain}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFuncmain(Calc.FuncmainContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#instrucciones}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInstrucciones(Calc.InstruccionesContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#instruccion}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInstruccion(Calc.InstruccionContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#if_instr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIf_instr(Calc.If_instrContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#listaelseif}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitListaelseif(Calc.ListaelseifContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#else_if}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitElse_if(Calc.Else_ifContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#bloque}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBloque(Calc.BloqueContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#consola}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConsola(Calc.ConsolaContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#llamada}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLlamada(Calc.LlamadaContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#listaExpresiones}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitListaExpresiones(Calc.ListaExpresionesContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#declaracionIni}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeclaracionIni(Calc.DeclaracionIniContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#declaracion}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeclaracion(Calc.DeclaracionContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#listides}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitListides(Calc.ListidesContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#tiposvars}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTiposvars(Calc.TiposvarsContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression(Calc.ExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#expr_rel}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpr_rel(Calc.Expr_relContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#expr_arit}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpr_arit(Calc.Expr_aritContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#expr_valor}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpr_valor(Calc.Expr_valorContext ctx);
	/**
	 * Visit a parse tree produced by {@link Calc#primitivo}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimitivo(Calc.PrimitivoContext ctx);
}