// Generated from Calc.g4 by ANTLR 4.7.2


    import "OLC2/analizador/ast"
    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/expresion/Accesos"
    import "OLC2/analizador/ast/funbasica"
    import "OLC2/analizador/ast/definicion"
    import "OLC2/analizador/ast/definicion/defobjetos"
    import "OLC2/analizador/ast/definicion/defarreglos"
    import "OLC2/analizador/ast/control"
    import "OLC2/analizador/ast/expresion2"
    import "OLC2/analizador/ast/transferenciaFlujo"
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
	 * Enter a parse tree produced by {@link Calc#listaClases}.
	 * @param ctx the parse tree
	 */
	void enterListaClases(Calc.ListaClasesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listaClases}.
	 * @param ctx the parse tree
	 */
	void exitListaClases(Calc.ListaClasesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#clases}.
	 * @param ctx the parse tree
	 */
	void enterClases(Calc.ClasesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#clases}.
	 * @param ctx the parse tree
	 */
	void exitClases(Calc.ClasesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#cuerpoClase}.
	 * @param ctx the parse tree
	 */
	void enterCuerpoClase(Calc.CuerpoClaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#cuerpoClase}.
	 * @param ctx the parse tree
	 */
	void exitCuerpoClase(Calc.CuerpoClaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#contenidoClase}.
	 * @param ctx the parse tree
	 */
	void enterContenidoClase(Calc.ContenidoClaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#contenidoClase}.
	 * @param ctx the parse tree
	 */
	void exitContenidoClase(Calc.ContenidoClaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#itemClase}.
	 * @param ctx the parse tree
	 */
	void enterItemClase(Calc.ItemClaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#itemClase}.
	 * @param ctx the parse tree
	 */
	void exitItemClase(Calc.ItemClaseContext ctx);
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
	 * Enter a parse tree produced by {@link Calc#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParametro(Calc.ParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParametro(Calc.ParametroContext ctx);
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
	 * Enter a parse tree produced by {@link Calc#dec_objeto}.
	 * @param ctx the parse tree
	 */
	void enterDec_objeto(Calc.Dec_objetoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#dec_objeto}.
	 * @param ctx the parse tree
	 */
	void exitDec_objeto(Calc.Dec_objetoContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#dec_arr}.
	 * @param ctx the parse tree
	 */
	void enterDec_arr(Calc.Dec_arrContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#dec_arr}.
	 * @param ctx the parse tree
	 */
	void exitDec_arr(Calc.Dec_arrContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#dimensiones}.
	 * @param ctx the parse tree
	 */
	void enterDimensiones(Calc.DimensionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#dimensiones}.
	 * @param ctx the parse tree
	 */
	void exitDimensiones(Calc.DimensionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#dimension}.
	 * @param ctx the parse tree
	 */
	void enterDimension(Calc.DimensionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#dimension}.
	 * @param ctx the parse tree
	 */
	void exitDimension(Calc.DimensionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#listanchos}.
	 * @param ctx the parse tree
	 */
	void enterListanchos(Calc.ListanchosContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listanchos}.
	 * @param ctx the parse tree
	 */
	void exitListanchos(Calc.ListanchosContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#ancho}.
	 * @param ctx the parse tree
	 */
	void enterAncho(Calc.AnchoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#ancho}.
	 * @param ctx the parse tree
	 */
	void exitAncho(Calc.AnchoContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion(Calc.AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion(Calc.AsignacionContext ctx);
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
	 * Enter a parse tree produced by {@link Calc#retorno}.
	 * @param ctx the parse tree
	 */
	void enterRetorno(Calc.RetornoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#retorno}.
	 * @param ctx the parse tree
	 */
	void exitRetorno(Calc.RetornoContext ctx);
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
	 * Enter a parse tree produced by {@link Calc#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpresion(Calc.ExpresionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpresion(Calc.ExpresionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#arraydata}.
	 * @param ctx the parse tree
	 */
	void enterArraydata(Calc.ArraydataContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#arraydata}.
	 * @param ctx the parse tree
	 */
	void exitArraydata(Calc.ArraydataContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#instancia}.
	 * @param ctx the parse tree
	 */
	void enterInstancia(Calc.InstanciaContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#instancia}.
	 * @param ctx the parse tree
	 */
	void exitInstancia(Calc.InstanciaContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#instanciaClase}.
	 * @param ctx the parse tree
	 */
	void enterInstanciaClase(Calc.InstanciaClaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#instanciaClase}.
	 * @param ctx the parse tree
	 */
	void exitInstanciaClase(Calc.InstanciaClaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#accesoarr}.
	 * @param ctx the parse tree
	 */
	void enterAccesoarr(Calc.AccesoarrContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#accesoarr}.
	 * @param ctx the parse tree
	 */
	void exitAccesoarr(Calc.AccesoarrContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#accesoObjeto}.
	 * @param ctx the parse tree
	 */
	void enterAccesoObjeto(Calc.AccesoObjetoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#accesoObjeto}.
	 * @param ctx the parse tree
	 */
	void exitAccesoObjeto(Calc.AccesoObjetoContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#listaAccesos}.
	 * @param ctx the parse tree
	 */
	void enterListaAccesos(Calc.ListaAccesosContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#listaAccesos}.
	 * @param ctx the parse tree
	 */
	void exitListaAccesos(Calc.ListaAccesosContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#acceso}.
	 * @param ctx the parse tree
	 */
	void enterAcceso(Calc.AccesoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#acceso}.
	 * @param ctx the parse tree
	 */
	void exitAcceso(Calc.AccesoContext ctx);
	/**
	 * Enter a parse tree produced by {@link Calc#expr_log}.
	 * @param ctx the parse tree
	 */
	void enterExpr_log(Calc.Expr_logContext ctx);
	/**
	 * Exit a parse tree produced by {@link Calc#expr_log}.
	 * @param ctx the parse tree
	 */
	void exitExpr_log(Calc.Expr_logContext ctx);
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