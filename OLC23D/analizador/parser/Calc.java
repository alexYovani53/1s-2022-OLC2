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

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Calc extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LP=1, RP=2, L_LLAVE=3, R_LLAVE=4, L_CORCH=5, R_CORCH=6, OUT=7, PRINTLN=8, 
		IF_TOK=9, ELSE=10, ARGS=11, CLASS=12, NEW_=13, MAIN=14, PRIVATE=15, PUBLIC=16, 
		STATIC=17, STRINGARGS=18, RETURN_P=19, INTTYPE=20, FLOATTYPE=21, STRINGTYPE=22, 
		VOIDTYPE=23, SYSTEM=24, BOOLTYPE=25, PUNTO=26, COMA=27, PTCOMA=28, AND=29, 
		OR=30, NOT=31, IGUAL=32, DIFERENTE=33, MAYORIGUAL=34, MENORIGUAL=35, MAYOR=36, 
		MENOR=37, MUL=38, DIV=39, ADD=40, SUB=41, NUMBER=42, FLOAT=43, STRING=44, 
		TRUE=45, FALSE=46, ID=47, WHITESPACE=48;
	public static final int
		RULE_start = 0, RULE_listaClases = 1, RULE_clases = 2, RULE_cuerpoClase = 3, 
		RULE_contenidoClase = 4, RULE_itemClase = 5, RULE_funcionItem = 6, RULE_modaccess = 7, 
		RULE_parametros = 8, RULE_parametro = 9, RULE_funcmain = 10, RULE_instrucciones = 11, 
		RULE_instruccion = 12, RULE_dec_objeto = 13, RULE_dec_arr = 14, RULE_dimensiones = 15, 
		RULE_dimension = 16, RULE_listanchos = 17, RULE_ancho = 18, RULE_asignacion = 19, 
		RULE_if_instr = 20, RULE_listaelseif = 21, RULE_else_if = 22, RULE_bloque = 23, 
		RULE_consola = 24, RULE_llamada = 25, RULE_listaExpresiones = 26, RULE_declaracionIni = 27, 
		RULE_declaracion = 28, RULE_retorno = 29, RULE_listides = 30, RULE_tiposvars = 31, 
		RULE_expresion = 32, RULE_arraydata = 33, RULE_instancia = 34, RULE_instanciaClase = 35, 
		RULE_accesoarr = 36, RULE_accesoObjeto = 37, RULE_listaAccesos = 38, RULE_acceso = 39, 
		RULE_expr_log = 40, RULE_expr_rel = 41, RULE_expr_arit = 42, RULE_expr_valor = 43, 
		RULE_primitivo = 44;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "listaClases", "clases", "cuerpoClase", "contenidoClase", "itemClase", 
			"funcionItem", "modaccess", "parametros", "parametro", "funcmain", "instrucciones", 
			"instruccion", "dec_objeto", "dec_arr", "dimensiones", "dimension", "listanchos", 
			"ancho", "asignacion", "if_instr", "listaelseif", "else_if", "bloque", 
			"consola", "llamada", "listaExpresiones", "declaracionIni", "declaracion", 
			"retorno", "listides", "tiposvars", "expresion", "arraydata", "instancia", 
			"instanciaClase", "accesoarr", "accesoObjeto", "listaAccesos", "acceso", 
			"expr_log", "expr_rel", "expr_arit", "expr_valor", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'['", "']'", "'out'", "'println'", 
			"'if'", "'else'", "'args'", "'class'", "'new'", "'main'", "'private'", 
			"'public'", "'static'", "'String'", "'return'", "'int'", "'float'", "'string'", 
			"'void'", "'System'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'", 
			"'!'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", 
			"'-'", null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LP", "RP", "L_LLAVE", "R_LLAVE", "L_CORCH", "R_CORCH", "OUT", 
			"PRINTLN", "IF_TOK", "ELSE", "ARGS", "CLASS", "NEW_", "MAIN", "PRIVATE", 
			"PUBLIC", "STATIC", "STRINGARGS", "RETURN_P", "INTTYPE", "FLOATTYPE", 
			"STRINGTYPE", "VOIDTYPE", "SYSTEM", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", 
			"AND", "OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", 
			"MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", 
			"TRUE", "FALSE", "ID", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Calc.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }




	public Calc(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public ast.Ast ast;
		public ListaClasesContext listaClases;
		public ListaClasesContext listaClases() {
			return getRuleContext(ListaClasesContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterStart(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitStart(this);
		}
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			((StartContext)_localctx).listaClases = listaClases(0);
			 _localctx.ast = ast.NewAst( ((StartContext)_localctx).listaClases.lista)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListaClasesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ListaClasesContext SUBLISTA;
		public ClasesContext clases;
		public ClasesContext clases() {
			return getRuleContext(ClasesContext.class,0);
		}
		public ListaClasesContext listaClases() {
			return getRuleContext(ListaClasesContext.class,0);
		}
		public ListaClasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaClases; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListaClases(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListaClases(this);
		}
	}

	public final ListaClasesContext listaClases() throws RecognitionException {
		return listaClases(0);
	}

	private ListaClasesContext listaClases(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListaClasesContext _localctx = new ListaClasesContext(_ctx, _parentState);
		ListaClasesContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_listaClases, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(94);
			((ListaClasesContext)_localctx).clases = clases();
			 _localctx.lista.Add( ((ListaClasesContext)_localctx).clases.instr ) 
			}
			_ctx.stop = _input.LT(-1);
			setState(103);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListaClasesContext(_parentctx, _parentState);
					_localctx.SUBLISTA = _prevctx;
					_localctx.SUBLISTA = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listaClases);
					setState(97);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(98);
					((ListaClasesContext)_localctx).clases = clases();

					                                                          ((ListaClasesContext)_localctx).SUBLISTA.lista.Add( ((ListaClasesContext)_localctx).clases.instr)
					                                                          _localctx.lista =  ((ListaClasesContext)_localctx).SUBLISTA.lista
					                                                      
					}
					} 
				}
				setState(105);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ClasesContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Token ID;
		public CuerpoClaseContext cuerpoClase;
		public TerminalNode CLASS() { return getToken(Calc.CLASS, 0); }
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public CuerpoClaseContext cuerpoClase() {
			return getRuleContext(CuerpoClaseContext.class,0);
		}
		public ClasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_clases; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterClases(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitClases(this);
		}
	}

	public final ClasesContext clases() throws RecognitionException {
		ClasesContext _localctx = new ClasesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_clases);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			match(CLASS);
			setState(107);
			((ClasesContext)_localctx).ID = match(ID);
			setState(108);
			((ClasesContext)_localctx).cuerpoClase = cuerpoClase();
			_localctx.instr = definicion.NewDefClase((((ClasesContext)_localctx).ID!=null?((ClasesContext)_localctx).ID.getText():null), ((ClasesContext)_localctx).cuerpoClase.lista)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CuerpoClaseContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ContenidoClaseContext contenidoClase;
		public TerminalNode L_LLAVE() { return getToken(Calc.L_LLAVE, 0); }
		public ContenidoClaseContext contenidoClase() {
			return getRuleContext(ContenidoClaseContext.class,0);
		}
		public TerminalNode R_LLAVE() { return getToken(Calc.R_LLAVE, 0); }
		public CuerpoClaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cuerpoClase; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterCuerpoClase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitCuerpoClase(this);
		}
	}

	public final CuerpoClaseContext cuerpoClase() throws RecognitionException {
		CuerpoClaseContext _localctx = new CuerpoClaseContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_cuerpoClase);
		try {
			setState(119);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				match(L_LLAVE);
				setState(112);
				((CuerpoClaseContext)_localctx).contenidoClase = contenidoClase(0);
				setState(113);
				match(R_LLAVE);
				_localctx.lista = ((CuerpoClaseContext)_localctx).contenidoClase.lista
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(116);
				match(L_LLAVE);
				setState(117);
				match(R_LLAVE);
				_localctx.lista = arrayList.New()
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContenidoClaseContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ContenidoClaseContext SUBLISTA;
		public ItemClaseContext itemClase;
		public ItemClaseContext itemClase() {
			return getRuleContext(ItemClaseContext.class,0);
		}
		public ContenidoClaseContext contenidoClase() {
			return getRuleContext(ContenidoClaseContext.class,0);
		}
		public ContenidoClaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contenidoClase; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterContenidoClase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitContenidoClase(this);
		}
	}

	public final ContenidoClaseContext contenidoClase() throws RecognitionException {
		return contenidoClase(0);
	}

	private ContenidoClaseContext contenidoClase(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ContenidoClaseContext _localctx = new ContenidoClaseContext(_ctx, _parentState);
		ContenidoClaseContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_contenidoClase, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(122);
			((ContenidoClaseContext)_localctx).itemClase = itemClase();

			                                                            _localctx.lista.Add( ((ContenidoClaseContext)_localctx).itemClase.instr )
			                                                        
			}
			_ctx.stop = _input.LT(-1);
			setState(131);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ContenidoClaseContext(_parentctx, _parentState);
					_localctx.SUBLISTA = _prevctx;
					_localctx.SUBLISTA = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_contenidoClase);
					setState(125);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(126);
					((ContenidoClaseContext)_localctx).itemClase = itemClase();

					                                                                      ((ContenidoClaseContext)_localctx).SUBLISTA.lista.Add( ((ContenidoClaseContext)_localctx).itemClase.instr )
					                                                                      _localctx.lista = ((ContenidoClaseContext)_localctx).SUBLISTA.lista
					                                                                  
					}
					} 
				}
				setState(133);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ItemClaseContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public FuncionItemContext funcionItem;
		public DeclaracionIniContext declaracionIni;
		public DeclaracionContext declaracion;
		public FuncionItemContext funcionItem() {
			return getRuleContext(FuncionItemContext.class,0);
		}
		public DeclaracionIniContext declaracionIni() {
			return getRuleContext(DeclaracionIniContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Calc.PTCOMA, 0); }
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public ItemClaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_itemClase; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterItemClase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitItemClase(this);
		}
	}

	public final ItemClaseContext itemClase() throws RecognitionException {
		ItemClaseContext _localctx = new ItemClaseContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_itemClase);
		try {
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(134);
				((ItemClaseContext)_localctx).funcionItem = funcionItem();
				_localctx.instr = ((ItemClaseContext)_localctx).funcionItem.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				((ItemClaseContext)_localctx).declaracionIni = declaracionIni();
				setState(138);
				match(PTCOMA);
				_localctx.instr = ((ItemClaseContext)_localctx).declaracionIni.instr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(141);
				((ItemClaseContext)_localctx).declaracion = declaracion();
				setState(142);
				match(PTCOMA);
				_localctx.instr = ((ItemClaseContext)_localctx).declaracion.instr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncionItemContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public FuncmainContext funcmain;
		public Token ID;
		public BloqueContext bloque;
		public TiposvarsContext tiposvars;
		public ParametrosContext parametros;
		public FuncmainContext funcmain() {
			return getRuleContext(FuncmainContext.class,0);
		}
		public ModaccessContext modaccess() {
			return getRuleContext(ModaccessContext.class,0);
		}
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public FuncionItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionItem; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterFuncionItem(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitFuncionItem(this);
		}
	}

	public final FuncionItemContext funcionItem() throws RecognitionException {
		FuncionItemContext _localctx = new FuncionItemContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_funcionItem);
		 listaParams :=  arrayList.New() 
		try {
			setState(167);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(147);
				((FuncionItemContext)_localctx).funcmain = funcmain();
				_localctx.instr =  ((FuncionItemContext)_localctx).funcmain.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(150);
				modaccess();
				setState(151);
				tiposvars();
				setState(152);
				((FuncionItemContext)_localctx).ID = match(ID);
				setState(153);
				match(LP);
				setState(154);
				match(RP);
				setState(155);
				((FuncionItemContext)_localctx).bloque = bloque();
				 _localctx.instr = Simbolos.NewFuncion((((FuncionItemContext)_localctx).ID!=null?((FuncionItemContext)_localctx).ID.getText():null),listaParams,((FuncionItemContext)_localctx).bloque.lista,entorno.VOID)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(158);
				modaccess();
				setState(159);
				((FuncionItemContext)_localctx).tiposvars = tiposvars();
				setState(160);
				((FuncionItemContext)_localctx).ID = match(ID);
				setState(161);
				match(LP);
				setState(162);
				((FuncionItemContext)_localctx).parametros = parametros(0);
				setState(163);
				match(RP);
				setState(164);
				((FuncionItemContext)_localctx).bloque = bloque();
				 _localctx.instr = Simbolos.NewFuncion((((FuncionItemContext)_localctx).ID!=null?((FuncionItemContext)_localctx).ID.getText():null),((FuncionItemContext)_localctx).parametros.lista,((FuncionItemContext)_localctx).bloque.lista,((FuncionItemContext)_localctx).tiposvars.tipo)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ModaccessContext extends ParserRuleContext {
		public entorno.TipoModAccess modAccess;
		public TerminalNode PUBLIC() { return getToken(Calc.PUBLIC, 0); }
		public TerminalNode PRIVATE() { return getToken(Calc.PRIVATE, 0); }
		public ModaccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_modaccess; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterModaccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitModaccess(this);
		}
	}

	public final ModaccessContext modaccess() throws RecognitionException {
		ModaccessContext _localctx = new ModaccessContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_modaccess);
		try {
			setState(174);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PUBLIC:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				match(PUBLIC);
				 _localctx.modAccess = entorno.PUBLIC
				}
				break;
			case PRIVATE:
				enterOuterAlt(_localctx, 2);
				{
				setState(171);
				match(PRIVATE);
				 _localctx.modAccess = entorno.PRIVATE
				}
				break;
			case INTTYPE:
			case FLOATTYPE:
			case STRINGTYPE:
			case VOIDTYPE:
			case BOOLTYPE:
				enterOuterAlt(_localctx, 3);
				{
				 _localctx.modAccess = entorno.PRIVATE
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametrosContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ParametrosContext sublista;
		public ParametroContext parametro;
		public ParametroContext parametro() {
			return getRuleContext(ParametroContext.class,0);
		}
		public TerminalNode COMA() { return getToken(Calc.COMA, 0); }
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterParametros(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitParametros(this);
		}
	}

	public final ParametrosContext parametros() throws RecognitionException {
		return parametros(0);
	}

	private ParametrosContext parametros(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ParametrosContext _localctx = new ParametrosContext(_ctx, _parentState);
		ParametrosContext _prevctx = _localctx;
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_parametros, _p);

		_localctx.lista =  arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(177);
			((ParametrosContext)_localctx).parametro = parametro();

			                                                                    _localctx.lista.Add( ((ParametrosContext)_localctx).parametro.instr)
			                                                                 
			}
			_ctx.stop = _input.LT(-1);
			setState(187);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParametrosContext(_parentctx, _parentState);
					_localctx.sublista = _prevctx;
					_localctx.sublista = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_parametros);
					setState(180);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(181);
					match(COMA);
					setState(182);
					((ParametrosContext)_localctx).parametro = parametro();

					                                                                              ((ParametrosContext)_localctx).sublista.lista.Add( ((ParametrosContext)_localctx).parametro.instr )
					                                                                              _localctx.lista =  ((ParametrosContext)_localctx).sublista.lista
					                                                                           
					}
					} 
				}
				setState(189);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParametroContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public TiposvarsContext tiposvars;
		public Token ID;
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode MUL() { return getToken(Calc.MUL, 0); }
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterParametro(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitParametro(this);
		}
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_parametro);
		try {
			setState(199);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTTYPE:
			case FLOATTYPE:
			case STRINGTYPE:
			case VOIDTYPE:
			case BOOLTYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				((ParametroContext)_localctx).tiposvars = tiposvars();
				setState(191);
				((ParametroContext)_localctx).ID = match(ID);

				                                                                    listaIdes := arrayList.New()
				                                                                    listaIdes.Add(expresion.NewIdentificador((((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getText():null)))
				                                                                    _localctx.instr = definicion.NewDeclaracionParametro(listaIdes, ((ParametroContext)_localctx).tiposvars.tipo,false)
				                                                                
				}
				break;
			case MUL:
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				match(MUL);
				setState(195);
				((ParametroContext)_localctx).tiposvars = tiposvars();
				setState(196);
				((ParametroContext)_localctx).ID = match(ID);

				                                                                    listaIdes := arrayList.New()
				                                                                    listaIdes.Add(expresion.NewIdentificador((((ParametroContext)_localctx).ID!=null?((ParametroContext)_localctx).ID.getText():null)))
				                                                                    _localctx.instr = definicion.NewDeclaracionParametro(listaIdes, ((ParametroContext)_localctx).tiposvars.tipo,true)
				                                                                
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncmainContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public BloqueContext bloque;
		public TerminalNode PUBLIC() { return getToken(Calc.PUBLIC, 0); }
		public TerminalNode STATIC() { return getToken(Calc.STATIC, 0); }
		public TerminalNode VOIDTYPE() { return getToken(Calc.VOIDTYPE, 0); }
		public TerminalNode MAIN() { return getToken(Calc.MAIN, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public TerminalNode STRINGARGS() { return getToken(Calc.STRINGARGS, 0); }
		public TerminalNode ARGS() { return getToken(Calc.ARGS, 0); }
		public TerminalNode L_CORCH() { return getToken(Calc.L_CORCH, 0); }
		public TerminalNode R_CORCH() { return getToken(Calc.R_CORCH, 0); }
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public FuncmainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcmain; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterFuncmain(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitFuncmain(this);
		}
	}

	public final FuncmainContext funcmain() throws RecognitionException {
		FuncmainContext _localctx = new FuncmainContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_funcmain);
		 listaParams:= arrayList.New() 
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			match(PUBLIC);
			setState(202);
			match(STATIC);
			setState(203);
			match(VOIDTYPE);
			setState(204);
			match(MAIN);
			setState(205);
			match(LP);
			setState(206);
			match(STRINGARGS);
			setState(207);
			match(ARGS);
			setState(208);
			match(L_CORCH);
			setState(209);
			match(R_CORCH);
			setState(210);
			match(RP);
			setState(211);
			((FuncmainContext)_localctx).bloque = bloque();
			 _localctx.instr = Simbolos.NewFuncion("main",listaParams,((FuncmainContext)_localctx).bloque.lista,entorno.VOID)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterInstrucciones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitInstrucciones(this);
		}
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_instrucciones);
		 _localctx.lista =  arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(214);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(217); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF_TOK) | (1L << RETURN_P) | (1L << INTTYPE) | (1L << FLOATTYPE) | (1L << STRINGTYPE) | (1L << VOIDTYPE) | (1L << SYSTEM) | (1L << BOOLTYPE) | (1L << ID))) != 0) );

			                                                                    listInt := localctx.(*InstruccionesContext).GetE()
			                                                                        for _, e := range listInt {
			                                                                          _localctx.lista.Add(e.GetInstr())
			                                                                        }
			                                                                    fmt.Printf("tipo %T",localctx.(*InstruccionesContext).GetE())
			                                                                
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public If_instrContext if_instr;
		public ConsolaContext consola;
		public DeclaracionIniContext declaracionIni;
		public DeclaracionContext declaracion;
		public LlamadaContext llamada;
		public RetornoContext retorno;
		public Dec_arrContext dec_arr;
		public Dec_objetoContext dec_objeto;
		public AsignacionContext asignacion;
		public If_instrContext if_instr() {
			return getRuleContext(If_instrContext.class,0);
		}
		public ConsolaContext consola() {
			return getRuleContext(ConsolaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Calc.PTCOMA, 0); }
		public DeclaracionIniContext declaracionIni() {
			return getRuleContext(DeclaracionIniContext.class,0);
		}
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public LlamadaContext llamada() {
			return getRuleContext(LlamadaContext.class,0);
		}
		public RetornoContext retorno() {
			return getRuleContext(RetornoContext.class,0);
		}
		public Dec_arrContext dec_arr() {
			return getRuleContext(Dec_arrContext.class,0);
		}
		public Dec_objetoContext dec_objeto() {
			return getRuleContext(Dec_objetoContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterInstruccion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitInstruccion(this);
		}
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_instruccion);
		try {
			setState(256);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(221);
				((InstruccionContext)_localctx).if_instr = if_instr();
				_localctx.instr = ((InstruccionContext)_localctx).if_instr.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(224);
				((InstruccionContext)_localctx).consola = consola();
				setState(225);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).consola.instr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(228);
				((InstruccionContext)_localctx).declaracionIni = declaracionIni();
				setState(229);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).declaracionIni.instr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(232);
				((InstruccionContext)_localctx).declaracion = declaracion();
				setState(233);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).declaracion.instr
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(236);
				((InstruccionContext)_localctx).llamada = llamada();
				setState(237);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).llamada.instr
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(240);
				((InstruccionContext)_localctx).retorno = retorno();
				setState(241);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).retorno.instr
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(244);
				((InstruccionContext)_localctx).dec_arr = dec_arr();
				setState(245);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).dec_arr.instr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(248);
				((InstruccionContext)_localctx).dec_objeto = dec_objeto();
				setState(249);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).dec_objeto.instr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(252);
				((InstruccionContext)_localctx).asignacion = asignacion();
				setState(253);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).asignacion.instr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dec_objetoContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Token ID;
		public ListidesContext listides;
		public ExpresionContext expresion;
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public ListidesContext listides() {
			return getRuleContext(ListidesContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Calc.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Dec_objetoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_objeto; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterDec_objeto(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitDec_objeto(this);
		}
	}

	public final Dec_objetoContext dec_objeto() throws RecognitionException {
		Dec_objetoContext _localctx = new Dec_objetoContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_dec_objeto);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			((Dec_objetoContext)_localctx).ID = match(ID);
			setState(259);
			((Dec_objetoContext)_localctx).listides = listides(0);
			setState(260);
			match(IGUAL);
			setState(261);
			((Dec_objetoContext)_localctx).expresion = expresion();
			_localctx.instr = defobjetos.NewDeclararObjeto( (((Dec_objetoContext)_localctx).ID!=null?((Dec_objetoContext)_localctx).ID.getText():null), ((Dec_objetoContext)_localctx).listides.lista, ((Dec_objetoContext)_localctx).expresion.expr)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dec_arrContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public TiposvarsContext tiposvars;
		public DimensionesContext dimensiones;
		public Token ID;
		public ExpresionContext expresion;
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public DimensionesContext dimensiones() {
			return getRuleContext(DimensionesContext.class,0);
		}
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Calc.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Dec_arrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_arr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterDec_arr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitDec_arr(this);
		}
	}

	public final Dec_arrContext dec_arr() throws RecognitionException {
		Dec_arrContext _localctx = new Dec_arrContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_dec_arr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(264);
			((Dec_arrContext)_localctx).tiposvars = tiposvars();
			setState(265);
			((Dec_arrContext)_localctx).dimensiones = dimensiones(0);
			setState(266);
			((Dec_arrContext)_localctx).ID = match(ID);
			setState(267);
			match(IGUAL);
			setState(268);
			((Dec_arrContext)_localctx).expresion = expresion();
			_localctx.instr = defarreglos.NewDeclaracionArray(((Dec_arrContext)_localctx).dimensiones.tam,(((Dec_arrContext)_localctx).ID!=null?((Dec_arrContext)_localctx).ID.getText():null),((Dec_arrContext)_localctx).expresion.expr,((Dec_arrContext)_localctx).tiposvars.tipo)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DimensionesContext extends ParserRuleContext {
		public int tam;
		public DimensionesContext tamano;
		public DimensionContext dimension() {
			return getRuleContext(DimensionContext.class,0);
		}
		public DimensionesContext dimensiones() {
			return getRuleContext(DimensionesContext.class,0);
		}
		public DimensionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dimensiones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterDimensiones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitDimensiones(this);
		}
	}

	public final DimensionesContext dimensiones() throws RecognitionException {
		return dimensiones(0);
	}

	private DimensionesContext dimensiones(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DimensionesContext _localctx = new DimensionesContext(_ctx, _parentState);
		DimensionesContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_dimensiones, _p);
		 _localctx.tam = 0
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(272);
			dimension();
			_localctx.tam = 1
			}
			_ctx.stop = _input.LT(-1);
			setState(281);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new DimensionesContext(_parentctx, _parentState);
					_localctx.tamano = _prevctx;
					_localctx.tamano = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_dimensiones);
					setState(275);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(276);
					dimension();


					                                                                              _localctx.tam = ((DimensionesContext)_localctx).tamano.tam + 1
					                                                                           
					}
					} 
				}
				setState(283);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DimensionContext extends ParserRuleContext {
		public TerminalNode L_CORCH() { return getToken(Calc.L_CORCH, 0); }
		public TerminalNode R_CORCH() { return getToken(Calc.R_CORCH, 0); }
		public DimensionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dimension; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterDimension(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitDimension(this);
		}
	}

	public final DimensionContext dimension() throws RecognitionException {
		DimensionContext _localctx = new DimensionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_dimension);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			match(L_CORCH);
			setState(285);
			match(R_CORCH);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListanchosContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ListanchosContext sublist;
		public AnchoContext ancho;
		public AnchoContext ancho() {
			return getRuleContext(AnchoContext.class,0);
		}
		public ListanchosContext listanchos() {
			return getRuleContext(ListanchosContext.class,0);
		}
		public ListanchosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listanchos; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListanchos(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListanchos(this);
		}
	}

	public final ListanchosContext listanchos() throws RecognitionException {
		return listanchos(0);
	}

	private ListanchosContext listanchos(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListanchosContext _localctx = new ListanchosContext(_ctx, _parentState);
		ListanchosContext _prevctx = _localctx;
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_listanchos, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(288);
			((ListanchosContext)_localctx).ancho = ancho();
			_localctx.lista.Add(((ListanchosContext)_localctx).ancho.expr)
			}
			_ctx.stop = _input.LT(-1);
			setState(297);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListanchosContext(_parentctx, _parentState);
					_localctx.sublist = _prevctx;
					_localctx.sublist = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listanchos);
					setState(291);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(292);
					((ListanchosContext)_localctx).ancho = ancho();

					                                                                                          ((ListanchosContext)_localctx).sublist.lista.Add(((ListanchosContext)_localctx).ancho.expr)
					                                                                                          _localctx.lista = ((ListanchosContext)_localctx).sublist.lista
					                                                                                      
					}
					} 
				}
				setState(299);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AnchoContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public ExpresionContext expresion;
		public TerminalNode L_CORCH() { return getToken(Calc.L_CORCH, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode R_CORCH() { return getToken(Calc.R_CORCH, 0); }
		public AnchoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ancho; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterAncho(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitAncho(this);
		}
	}

	public final AnchoContext ancho() throws RecognitionException {
		AnchoContext _localctx = new AnchoContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_ancho);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(300);
			match(L_CORCH);
			setState(301);
			((AnchoContext)_localctx).expresion = expresion();
			setState(302);
			match(R_CORCH);
			_localctx.expr = ((AnchoContext)_localctx).expresion.expr
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AsignacionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Token ID;
		public ExpresionContext expresion;
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Calc.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterAsignacion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitAsignacion(this);
		}
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(305);
			((AsignacionContext)_localctx).ID = match(ID);
			setState(306);
			match(IGUAL);
			setState(307);
			((AsignacionContext)_localctx).expresion = expresion();
			_localctx.instr = definicion.NewAsignacion((((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getText():null), ((AsignacionContext)_localctx).expresion.expr)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class If_instrContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpresionContext expresion;
		public BloqueContext bloque;
		public BloqueContext bprincipal;
		public BloqueContext belse;
		public ListaelseifContext listaelseif;
		public TerminalNode IF_TOK() { return getToken(Calc.IF_TOK, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public List<BloqueContext> bloque() {
			return getRuleContexts(BloqueContext.class);
		}
		public BloqueContext bloque(int i) {
			return getRuleContext(BloqueContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(Calc.ELSE, 0); }
		public ListaelseifContext listaelseif() {
			return getRuleContext(ListaelseifContext.class,0);
		}
		public If_instrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_instr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterIf_instr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitIf_instr(this);
		}
	}

	public final If_instrContext if_instr() throws RecognitionException {
		If_instrContext _localctx = new If_instrContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_if_instr);
		try {
			setState(336);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(310);
				match(IF_TOK);
				setState(311);
				match(LP);
				setState(312);
				((If_instrContext)_localctx).expresion = expresion();
				setState(313);
				match(RP);
				setState(314);
				((If_instrContext)_localctx).bloque = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expresion.expr,((If_instrContext)_localctx).bloque.lista,nil,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(317);
				match(IF_TOK);
				setState(318);
				match(LP);
				setState(319);
				((If_instrContext)_localctx).expresion = expresion();
				setState(320);
				match(RP);
				setState(321);
				((If_instrContext)_localctx).bprincipal = bloque();
				setState(322);
				match(ELSE);
				setState(323);
				((If_instrContext)_localctx).belse = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expresion.expr,((If_instrContext)_localctx).bprincipal.lista,nil,((If_instrContext)_localctx).belse.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(326);
				match(IF_TOK);
				setState(327);
				match(LP);
				setState(328);
				((If_instrContext)_localctx).expresion = expresion();
				setState(329);
				match(RP);
				setState(330);
				((If_instrContext)_localctx).bprincipal = bloque();
				setState(331);
				((If_instrContext)_localctx).listaelseif = listaelseif();
				setState(332);
				match(ELSE);
				setState(333);
				((If_instrContext)_localctx).belse = bloque();

				        _localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expresion.expr,((If_instrContext)_localctx).bprincipal.lista,((If_instrContext)_localctx).listaelseif.lista,((If_instrContext)_localctx).belse.lista)
				    
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListaelseifContext extends ParserRuleContext {
		public *arrayList.List lista;
		public Else_ifContext else_if;
		public List<Else_ifContext> list = new ArrayList<Else_ifContext>();
		public List<Else_ifContext> else_if() {
			return getRuleContexts(Else_ifContext.class);
		}
		public Else_ifContext else_if(int i) {
			return getRuleContext(Else_ifContext.class,i);
		}
		public ListaelseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaelseif; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListaelseif(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListaelseif(this);
		}
	}

	public final ListaelseifContext listaelseif() throws RecognitionException {
		ListaelseifContext _localctx = new ListaelseifContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(339); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(338);
					((ListaelseifContext)_localctx).else_if = else_if();
					((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).else_if);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(341); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			                                                                            listInt := localctx.(*ListaelseifContext).GetList()
			                                                                            for _, e := range listInt {
			                                                                                _localctx.lista.Add(e.GetInstr())
			                                                                            }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Else_ifContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpresionContext expresion;
		public BloqueContext bloque;
		public TerminalNode ELSE() { return getToken(Calc.ELSE, 0); }
		public TerminalNode IF_TOK() { return getToken(Calc.IF_TOK, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public Else_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_if; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterElse_if(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitElse_if(this);
		}
	}

	public final Else_ifContext else_if() throws RecognitionException {
		Else_ifContext _localctx = new Else_ifContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_else_if);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			match(ELSE);
			setState(346);
			match(IF_TOK);
			setState(347);
			match(LP);
			setState(348);
			((Else_ifContext)_localctx).expresion = expresion();
			setState(349);
			match(RP);
			setState(350);
			((Else_ifContext)_localctx).bloque = bloque();
			_localctx.instr = control.NewIfInstruccion(((Else_ifContext)_localctx).expresion.expr,((Else_ifContext)_localctx).bloque.lista,nil,nil)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BloqueContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public TerminalNode L_LLAVE() { return getToken(Calc.L_LLAVE, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode R_LLAVE() { return getToken(Calc.R_LLAVE, 0); }
		public BloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloque; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterBloque(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitBloque(this);
		}
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_bloque);
		try {
			setState(361);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(353);
				match(L_LLAVE);
				setState(354);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(355);
				match(R_LLAVE);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(358);
				match(L_LLAVE);
				setState(359);
				match(R_LLAVE);
				_localctx.lista = arrayList.New()
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConsolaContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpresionContext expresion;
		public TerminalNode SYSTEM() { return getToken(Calc.SYSTEM, 0); }
		public List<TerminalNode> PUNTO() { return getTokens(Calc.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(Calc.PUNTO, i);
		}
		public TerminalNode OUT() { return getToken(Calc.OUT, 0); }
		public TerminalNode PRINTLN() { return getToken(Calc.PRINTLN, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public ConsolaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_consola; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterConsola(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitConsola(this);
		}
	}

	public final ConsolaContext consola() throws RecognitionException {
		ConsolaContext _localctx = new ConsolaContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_consola);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(SYSTEM);
			setState(364);
			match(PUNTO);
			setState(365);
			match(OUT);
			setState(366);
			match(PUNTO);
			setState(367);
			match(PRINTLN);
			setState(368);
			match(LP);
			setState(369);
			((ConsolaContext)_localctx).expresion = expresion();
			setState(370);
			match(RP);
			_localctx.instr = funbasica.NewImprimir(((ConsolaContext)_localctx).expresion.expr)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LlamadaContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public interfaces.Expresion expr;
		public Token ID;
		public ListaExpresionesContext listaExpresiones;
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public ListaExpresionesContext listaExpresiones() {
			return getRuleContext(ListaExpresionesContext.class,0);
		}
		public LlamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamada; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterLlamada(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitLlamada(this);
		}
	}

	public final LlamadaContext llamada() throws RecognitionException {
		LlamadaContext _localctx = new LlamadaContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_llamada);
		try {
			setState(383);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(373);
				((LlamadaContext)_localctx).ID = match(ID);
				setState(374);
				match(LP);
				setState(375);
				match(RP);

				                                                                        _localctx.instr = expresion2.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), arrayList.New())
				                                                                        _localctx.expr = expresion2.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), arrayList.New())
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(377);
				((LlamadaContext)_localctx).ID = match(ID);
				setState(378);
				match(LP);
				setState(379);
				((LlamadaContext)_localctx).listaExpresiones = listaExpresiones(0);
				setState(380);
				match(RP);

				                                                                        _localctx.instr = expresion2.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), ((LlamadaContext)_localctx).listaExpresiones.lista)
				                                                                        _localctx.expr = expresion2.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), ((LlamadaContext)_localctx).listaExpresiones.lista)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListaExpresionesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ListaExpresionesContext LISTA;
		public ExpresionContext expresion;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode COMA() { return getToken(Calc.COMA, 0); }
		public ListaExpresionesContext listaExpresiones() {
			return getRuleContext(ListaExpresionesContext.class,0);
		}
		public ListaExpresionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaExpresiones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListaExpresiones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListaExpresiones(this);
		}
	}

	public final ListaExpresionesContext listaExpresiones() throws RecognitionException {
		return listaExpresiones(0);
	}

	private ListaExpresionesContext listaExpresiones(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListaExpresionesContext _localctx = new ListaExpresionesContext(_ctx, _parentState);
		ListaExpresionesContext _prevctx = _localctx;
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_listaExpresiones, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(386);
			((ListaExpresionesContext)_localctx).expresion = expresion();

			                                                                        _localctx.lista.Add( ((ListaExpresionesContext)_localctx).expresion.expr )
			                                                                     
			}
			_ctx.stop = _input.LT(-1);
			setState(396);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListaExpresionesContext(_parentctx, _parentState);
					_localctx.LISTA = _prevctx;
					_localctx.LISTA = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listaExpresiones);
					setState(389);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(390);
					match(COMA);
					setState(391);
					((ListaExpresionesContext)_localctx).expresion = expresion();

					                                                                                  ((ListaExpresionesContext)_localctx).LISTA.lista.Add( ((ListaExpresionesContext)_localctx).expresion.expr )
					                                                                                  _localctx.lista =  ((ListaExpresionesContext)_localctx).LISTA.lista
					                                                                               
					}
					} 
				}
				setState(398);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DeclaracionIniContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public TiposvarsContext tiposvars;
		public ListidesContext listides;
		public ExpresionContext expresion;
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public ListidesContext listides() {
			return getRuleContext(ListidesContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Calc.IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclaracionIniContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracionIni; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterDeclaracionIni(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitDeclaracionIni(this);
		}
	}

	public final DeclaracionIniContext declaracionIni() throws RecognitionException {
		DeclaracionIniContext _localctx = new DeclaracionIniContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_declaracionIni);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(399);
			((DeclaracionIniContext)_localctx).tiposvars = tiposvars();
			setState(400);
			((DeclaracionIniContext)_localctx).listides = listides(0);
			setState(401);
			match(IGUAL);
			setState(402);
			((DeclaracionIniContext)_localctx).expresion = expresion();

			                                                                        _localctx.instr = definicion.NewDeclaracionInicializacion(((DeclaracionIniContext)_localctx).listides.lista,((DeclaracionIniContext)_localctx).tiposvars.tipo,((DeclaracionIniContext)_localctx).expresion.expr)
			                                                                     
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclaracionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public TiposvarsContext tiposvars;
		public ListidesContext listides;
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public ListidesContext listides() {
			return getRuleContext(ListidesContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterDeclaracion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitDeclaracion(this);
		}
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			((DeclaracionContext)_localctx).tiposvars = tiposvars();
			setState(406);
			((DeclaracionContext)_localctx).listides = listides(0);

			                                                                        _localctx.instr = definicion.NewDeclaracion(((DeclaracionContext)_localctx).listides.lista,((DeclaracionContext)_localctx).tiposvars.tipo)
			                                                                    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RetornoContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpresionContext expresion;
		public TerminalNode RETURN_P() { return getToken(Calc.RETURN_P, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public RetornoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_retorno; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterRetorno(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitRetorno(this);
		}
	}

	public final RetornoContext retorno() throws RecognitionException {
		RetornoContext _localctx = new RetornoContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_retorno);
		try {
			setState(415);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(409);
				match(RETURN_P);
				 _localctx.instr = transferenciaFlujo.NewReturn(entorno.VOID,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(411);
				match(RETURN_P);
				setState(412);
				((RetornoContext)_localctx).expresion = expresion();
				 _localctx.instr = transferenciaFlujo.NewReturn(entorno.NULL,((RetornoContext)_localctx).expresion.expr)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListidesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ListidesContext sub;
		public Token ID;
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode COMA() { return getToken(Calc.COMA, 0); }
		public ListidesContext listides() {
			return getRuleContext(ListidesContext.class,0);
		}
		public ListidesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listides; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListides(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListides(this);
		}
	}

	public final ListidesContext listides() throws RecognitionException {
		return listides(0);
	}

	private ListidesContext listides(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListidesContext _localctx = new ListidesContext(_ctx, _parentState);
		ListidesContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_listides, _p);
		  _localctx.lista =  arrayList.New() 
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(418);
			((ListidesContext)_localctx).ID = match(ID);
			_localctx.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
			}
			_ctx.stop = _input.LT(-1);
			setState(427);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListidesContext(_parentctx, _parentState);
					_localctx.sub = _prevctx;
					_localctx.sub = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listides);
					setState(421);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(422);
					match(COMA);
					setState(423);
					((ListidesContext)_localctx).ID = match(ID);

					                                                                              ((ListidesContext)_localctx).sub.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
					                                                                              _localctx.lista = ((ListidesContext)_localctx).sub.lista
					                                                                          
					}
					} 
				}
				setState(429);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class TiposvarsContext extends ParserRuleContext {
		public entorno.TipoDato tipo;
		public TerminalNode INTTYPE() { return getToken(Calc.INTTYPE, 0); }
		public TerminalNode STRINGTYPE() { return getToken(Calc.STRINGTYPE, 0); }
		public TerminalNode FLOATTYPE() { return getToken(Calc.FLOATTYPE, 0); }
		public TerminalNode BOOLTYPE() { return getToken(Calc.BOOLTYPE, 0); }
		public TerminalNode VOIDTYPE() { return getToken(Calc.VOIDTYPE, 0); }
		public TiposvarsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tiposvars; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterTiposvars(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitTiposvars(this);
		}
	}

	public final TiposvarsContext tiposvars() throws RecognitionException {
		TiposvarsContext _localctx = new TiposvarsContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_tiposvars);
		try {
			setState(440);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTTYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(430);
				match(INTTYPE);
				_localctx.tipo = entorno.INTEGER
				}
				break;
			case STRINGTYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(432);
				match(STRINGTYPE);
				_localctx.tipo = entorno.STRING
				}
				break;
			case FLOATTYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(434);
				match(FLOATTYPE);
				_localctx.tipo = entorno.FLOAT
				}
				break;
			case BOOLTYPE:
				enterOuterAlt(_localctx, 4);
				{
				setState(436);
				match(BOOLTYPE);
				_localctx.tipo = entorno.BOOLEAN
				}
				break;
			case VOIDTYPE:
				enterOuterAlt(_localctx, 5);
				{
				setState(438);
				match(VOIDTYPE);
				_localctx.tipo = entorno.VOID
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpresionContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Expr_logContext expr_log;
		public Expr_relContext expr_rel;
		public Expr_aritContext expr_arit;
		public InstanciaContext instancia;
		public ArraydataContext arraydata;
		public InstanciaClaseContext instanciaClase;
		public Expr_logContext expr_log() {
			return getRuleContext(Expr_logContext.class,0);
		}
		public Expr_relContext expr_rel() {
			return getRuleContext(Expr_relContext.class,0);
		}
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public InstanciaContext instancia() {
			return getRuleContext(InstanciaContext.class,0);
		}
		public ArraydataContext arraydata() {
			return getRuleContext(ArraydataContext.class,0);
		}
		public InstanciaClaseContext instanciaClase() {
			return getRuleContext(InstanciaClaseContext.class,0);
		}
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterExpresion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitExpresion(this);
		}
	}

	public final ExpresionContext expresion() throws RecognitionException {
		ExpresionContext _localctx = new ExpresionContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_expresion);
		try {
			setState(460);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				((ExpresionContext)_localctx).expr_log = expr_log(0);
				_localctx.expr = ((ExpresionContext)_localctx).expr_log.expr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(445);
				((ExpresionContext)_localctx).expr_rel = expr_rel(0);
				_localctx.expr = ((ExpresionContext)_localctx).expr_rel.expr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(448);
				((ExpresionContext)_localctx).expr_arit = expr_arit(0);
				_localctx.expr = ((ExpresionContext)_localctx).expr_arit.expr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(451);
				((ExpresionContext)_localctx).instancia = instancia();
				_localctx.expr = ((ExpresionContext)_localctx).instancia.expr
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(454);
				((ExpresionContext)_localctx).arraydata = arraydata();
				_localctx.expr = ((ExpresionContext)_localctx).arraydata.expr
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(457);
				((ExpresionContext)_localctx).instanciaClase = instanciaClase();
				_localctx.expr = ((ExpresionContext)_localctx).instanciaClase.expr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArraydataContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public ListaExpresionesContext listaExpresiones;
		public TerminalNode L_LLAVE() { return getToken(Calc.L_LLAVE, 0); }
		public ListaExpresionesContext listaExpresiones() {
			return getRuleContext(ListaExpresionesContext.class,0);
		}
		public TerminalNode R_LLAVE() { return getToken(Calc.R_LLAVE, 0); }
		public ArraydataContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arraydata; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterArraydata(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitArraydata(this);
		}
	}

	public final ArraydataContext arraydata() throws RecognitionException {
		ArraydataContext _localctx = new ArraydataContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_arraydata);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(462);
			match(L_LLAVE);
			setState(463);
			((ArraydataContext)_localctx).listaExpresiones = listaExpresiones(0);
			setState(464);
			match(R_LLAVE);
			_localctx.expr = expresion2.NewValorArreglo(((ArraydataContext)_localctx).listaExpresiones.lista)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstanciaContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public TiposvarsContext tiposvars;
		public ListanchosContext listanchos;
		public TerminalNode NEW_() { return getToken(Calc.NEW_, 0); }
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public ListanchosContext listanchos() {
			return getRuleContext(ListanchosContext.class,0);
		}
		public InstanciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instancia; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterInstancia(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitInstancia(this);
		}
	}

	public final InstanciaContext instancia() throws RecognitionException {
		InstanciaContext _localctx = new InstanciaContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_instancia);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(467);
			match(NEW_);
			setState(468);
			((InstanciaContext)_localctx).tiposvars = tiposvars();
			setState(469);
			((InstanciaContext)_localctx).listanchos = listanchos(0);
			_localctx.expr = expresion2.NewInstanciaArreglo(((InstanciaContext)_localctx).tiposvars.tipo, ((InstanciaContext)_localctx).listanchos.lista )
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstanciaClaseContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Token ID;
		public TerminalNode NEW_() { return getToken(Calc.NEW_, 0); }
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public InstanciaClaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instanciaClase; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterInstanciaClase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitInstanciaClase(this);
		}
	}

	public final InstanciaClaseContext instanciaClase() throws RecognitionException {
		InstanciaClaseContext _localctx = new InstanciaClaseContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_instanciaClase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(472);
			match(NEW_);
			setState(473);
			((InstanciaClaseContext)_localctx).ID = match(ID);
			setState(474);
			match(LP);
			setState(475);
			match(RP);
			_localctx.expr = expresion2.NewInstanciaObjeto((((InstanciaClaseContext)_localctx).ID!=null?((InstanciaClaseContext)_localctx).ID.getText():null) )
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AccesoarrContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Token ID;
		public ListanchosContext listanchos;
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public ListanchosContext listanchos() {
			return getRuleContext(ListanchosContext.class,0);
		}
		public AccesoarrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accesoarr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterAccesoarr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitAccesoarr(this);
		}
	}

	public final AccesoarrContext accesoarr() throws RecognitionException {
		AccesoarrContext _localctx = new AccesoarrContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_accesoarr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(478);
			((AccesoarrContext)_localctx).ID = match(ID);
			setState(479);
			((AccesoarrContext)_localctx).listanchos = listanchos(0);
			_localctx.expr = Accesos.NewAccessoArr((((AccesoarrContext)_localctx).ID!=null?((AccesoarrContext)_localctx).ID.getText():null),((AccesoarrContext)_localctx).listanchos.lista)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AccesoObjetoContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public ListaAccesosContext listaAccesos;
		public ListaAccesosContext listaAccesos() {
			return getRuleContext(ListaAccesosContext.class,0);
		}
		public AccesoObjetoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accesoObjeto; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterAccesoObjeto(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitAccesoObjeto(this);
		}
	}

	public final AccesoObjetoContext accesoObjeto() throws RecognitionException {
		AccesoObjetoContext _localctx = new AccesoObjetoContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_accesoObjeto);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(482);
			((AccesoObjetoContext)_localctx).listaAccesos = listaAccesos(0);
			_localctx.expr = Accesos.NewAccesoObjeto( ((AccesoObjetoContext)_localctx).listaAccesos.lista)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListaAccesosContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ListaAccesosContext SUBLISTA;
		public AccesoContext acceso;
		public AccesoContext acceso() {
			return getRuleContext(AccesoContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(Calc.PUNTO, 0); }
		public ListaAccesosContext listaAccesos() {
			return getRuleContext(ListaAccesosContext.class,0);
		}
		public ListaAccesosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaAccesos; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListaAccesos(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListaAccesos(this);
		}
	}

	public final ListaAccesosContext listaAccesos() throws RecognitionException {
		return listaAccesos(0);
	}

	private ListaAccesosContext listaAccesos(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListaAccesosContext _localctx = new ListaAccesosContext(_ctx, _parentState);
		ListaAccesosContext _prevctx = _localctx;
		int _startState = 76;
		enterRecursionRule(_localctx, 76, RULE_listaAccesos, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(486);
			((ListaAccesosContext)_localctx).acceso = acceso();
			   _localctx.lista.Add(((ListaAccesosContext)_localctx).acceso.expr)
			}
			_ctx.stop = _input.LT(-1);
			setState(496);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListaAccesosContext(_parentctx, _parentState);
					_localctx.SUBLISTA = _prevctx;
					_localctx.SUBLISTA = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listaAccesos);
					setState(489);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(490);
					match(PUNTO);
					setState(491);
					((ListaAccesosContext)_localctx).acceso = acceso();

					                                                              ((ListaAccesosContext)_localctx).SUBLISTA.lista.Add( ((ListaAccesosContext)_localctx).acceso.expr)
					                                                              _localctx.lista = ((ListaAccesosContext)_localctx).SUBLISTA.lista
					                                                          
					}
					} 
				}
				setState(498);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AccesoContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Token ID;
		public AccesoarrContext accesoarr;
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public AccesoarrContext accesoarr() {
			return getRuleContext(AccesoarrContext.class,0);
		}
		public AccesoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_acceso; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterAcceso(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitAcceso(this);
		}
	}

	public final AccesoContext acceso() throws RecognitionException {
		AccesoContext _localctx = new AccesoContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_acceso);
		try {
			setState(504);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(499);
				((AccesoContext)_localctx).ID = match(ID);
				 _localctx.expr = expresion.NewIdentificador((((AccesoContext)_localctx).ID!=null?((AccesoContext)_localctx).ID.getText():null))
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(501);
				((AccesoContext)_localctx).accesoarr = accesoarr();
				 _localctx.expr = ((AccesoContext)_localctx).accesoarr.expr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_logContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Expr_logContext opIz;
		public Expr_relContext expr_rel;
		public Expr_logContext opDe;
		public Expr_relContext expr_rel() {
			return getRuleContext(Expr_relContext.class,0);
		}
		public TerminalNode AND() { return getToken(Calc.AND, 0); }
		public List<Expr_logContext> expr_log() {
			return getRuleContexts(Expr_logContext.class);
		}
		public Expr_logContext expr_log(int i) {
			return getRuleContext(Expr_logContext.class,i);
		}
		public TerminalNode OR() { return getToken(Calc.OR, 0); }
		public Expr_logContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_log; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterExpr_log(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitExpr_log(this);
		}
	}

	public final Expr_logContext expr_log() throws RecognitionException {
		return expr_log(0);
	}

	private Expr_logContext expr_log(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_logContext _localctx = new Expr_logContext(_ctx, _parentState);
		Expr_logContext _prevctx = _localctx;
		int _startState = 80;
		enterRecursionRule(_localctx, 80, RULE_expr_log, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(507);
			((Expr_logContext)_localctx).expr_rel = expr_rel(0);
			_localctx.expr = ((Expr_logContext)_localctx).expr_rel.expr
			}
			_ctx.stop = _input.LT(-1);
			setState(522);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(520);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_logContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_log);
						setState(510);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(511);
						match(AND);
						setState(512);
						((Expr_logContext)_localctx).opDe = expr_log(4);
						_localctx.expr = expresion.NewOperacion(((Expr_logContext)_localctx).opIz.expr,"&&",((Expr_logContext)_localctx).opDe.expr,false)
						}
						break;
					case 2:
						{
						_localctx = new Expr_logContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_log);
						setState(515);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(516);
						match(OR);
						setState(517);
						((Expr_logContext)_localctx).opDe = expr_log(3);
						_localctx.expr = expresion.NewOperacion(((Expr_logContext)_localctx).opIz.expr,"||",((Expr_logContext)_localctx).opDe.expr,false)
						}
						break;
					}
					} 
				}
				setState(524);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expr_relContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Expr_relContext opIz;
		public Expr_aritContext expr_arit;
		public Token op;
		public Expr_relContext opDe;
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public List<Expr_relContext> expr_rel() {
			return getRuleContexts(Expr_relContext.class);
		}
		public Expr_relContext expr_rel(int i) {
			return getRuleContext(Expr_relContext.class,i);
		}
		public TerminalNode MAYORIGUAL() { return getToken(Calc.MAYORIGUAL, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Calc.MENORIGUAL, 0); }
		public TerminalNode MENOR() { return getToken(Calc.MENOR, 0); }
		public TerminalNode MAYOR() { return getToken(Calc.MAYOR, 0); }
		public Expr_relContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_rel; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterExpr_rel(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitExpr_rel(this);
		}
	}

	public final Expr_relContext expr_rel() throws RecognitionException {
		return expr_rel(0);
	}

	private Expr_relContext expr_rel(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_relContext _localctx = new Expr_relContext(_ctx, _parentState);
		Expr_relContext _prevctx = _localctx;
		int _startState = 82;
		enterRecursionRule(_localctx, 82, RULE_expr_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(526);
			((Expr_relContext)_localctx).expr_arit = expr_arit(0);
			_localctx.expr = ((Expr_relContext)_localctx).expr_arit.expr
			}
			_ctx.stop = _input.LT(-1);
			setState(536);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Expr_relContext(_parentctx, _parentState);
					_localctx.opIz = _prevctx;
					_localctx.opIz = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expr_rel);
					setState(529);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(530);
					((Expr_relContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << MAYOR) | (1L << MENOR))) != 0)) ) {
						((Expr_relContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(531);
					((Expr_relContext)_localctx).opDe = expr_rel(3);
					_localctx.expr = expresion.NewOperacion(((Expr_relContext)_localctx).opIz.expr,(((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null),((Expr_relContext)_localctx).opDe.expr,false)
					}
					} 
				}
				setState(538);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expr_aritContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Expr_aritContext opIz;
		public ExpresionContext opU;
		public ExpresionContext expresion;
		public Expr_valorContext expr_valor;
		public Token op;
		public Expr_aritContext opDe;
		public TerminalNode SUB() { return getToken(Calc.SUB, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Expr_valorContext expr_valor() {
			return getRuleContext(Expr_valorContext.class,0);
		}
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public TerminalNode RP() { return getToken(Calc.RP, 0); }
		public List<Expr_aritContext> expr_arit() {
			return getRuleContexts(Expr_aritContext.class);
		}
		public Expr_aritContext expr_arit(int i) {
			return getRuleContext(Expr_aritContext.class,i);
		}
		public TerminalNode MUL() { return getToken(Calc.MUL, 0); }
		public TerminalNode DIV() { return getToken(Calc.DIV, 0); }
		public TerminalNode ADD() { return getToken(Calc.ADD, 0); }
		public Expr_aritContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_arit; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterExpr_arit(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitExpr_arit(this);
		}
	}

	public final Expr_aritContext expr_arit() throws RecognitionException {
		return expr_arit(0);
	}

	private Expr_aritContext expr_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_aritContext _localctx = new Expr_aritContext(_ctx, _parentState);
		Expr_aritContext _prevctx = _localctx;
		int _startState = 84;
		enterRecursionRule(_localctx, 84, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(552);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				setState(540);
				match(SUB);
				setState(541);
				((Expr_aritContext)_localctx).opU = ((Expr_aritContext)_localctx).expresion = expresion();
				_localctx.expr = expresion.NewOperacion(((Expr_aritContext)_localctx).opU.expr,"-",nil,true)
				}
				break;
			case NUMBER:
			case FLOAT:
			case STRING:
			case TRUE:
			case FALSE:
			case ID:
				{
				setState(544);
				((Expr_aritContext)_localctx).expr_valor = expr_valor();
				_localctx.expr = ((Expr_aritContext)_localctx).expr_valor.expr
				}
				break;
			case LP:
				{
				setState(547);
				match(LP);
				setState(548);
				((Expr_aritContext)_localctx).expresion = expresion();
				setState(549);
				match(RP);
				_localctx.expr = ((Expr_aritContext)_localctx).expresion.expr
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(566);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(564);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(554);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(555);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(556);
						((Expr_aritContext)_localctx).opDe = expr_arit(5);
						_localctx.expr = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.expr,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.expr,false)
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(559);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(560);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(561);
						((Expr_aritContext)_localctx).opDe = expr_arit(4);
						_localctx.expr = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.expr,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.expr,false)
						}
						break;
					}
					} 
				}
				setState(568);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expr_valorContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public PrimitivoContext primitivo;
		public LlamadaContext llamada;
		public AccesoarrContext accesoarr;
		public AccesoObjetoContext accesoObjeto;
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public LlamadaContext llamada() {
			return getRuleContext(LlamadaContext.class,0);
		}
		public AccesoarrContext accesoarr() {
			return getRuleContext(AccesoarrContext.class,0);
		}
		public AccesoObjetoContext accesoObjeto() {
			return getRuleContext(AccesoObjetoContext.class,0);
		}
		public Expr_valorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_valor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterExpr_valor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitExpr_valor(this);
		}
	}

	public final Expr_valorContext expr_valor() throws RecognitionException {
		Expr_valorContext _localctx = new Expr_valorContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_expr_valor);
		try {
			setState(581);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(569);
				((Expr_valorContext)_localctx).primitivo = primitivo();
				_localctx.expr = ((Expr_valorContext)_localctx).primitivo.expr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(572);
				((Expr_valorContext)_localctx).llamada = llamada();
				_localctx.expr = ((Expr_valorContext)_localctx).llamada.expr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(575);
				((Expr_valorContext)_localctx).accesoarr = accesoarr();
				_localctx.expr = ((Expr_valorContext)_localctx).accesoarr.expr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(578);
				((Expr_valorContext)_localctx).accesoObjeto = accesoObjeto();
				_localctx.expr = ((Expr_valorContext)_localctx).accesoObjeto.expr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expresion expr;
		public Token NUMBER;
		public Token FLOAT;
		public Token STRING;
		public Token ID;
		public TerminalNode NUMBER() { return getToken(Calc.NUMBER, 0); }
		public TerminalNode FLOAT() { return getToken(Calc.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(Calc.STRING, 0); }
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public TerminalNode TRUE() { return getToken(Calc.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Calc.FALSE, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterPrimitivo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitPrimitivo(this);
		}
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_primitivo);
		try {
			setState(595);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(583);
				((PrimitivoContext)_localctx).NUMBER = match(NUMBER);

				                                                                    num,err := strconv.Atoi((((PrimitivoContext)_localctx).NUMBER!=null?((PrimitivoContext)_localctx).NUMBER.getText():null))
				                                                                    if err!= nil{
				                                                                        fmt.Println(err)
				                                                                    }
				                                                                    _localctx.expr = expresion.NewPrimitivo (num,entorno.INTEGER)
				                                                                
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(585);
				((PrimitivoContext)_localctx).FLOAT = match(FLOAT);

				                                                                     num,err := strconv.ParseFloat((((PrimitivoContext)_localctx).FLOAT!=null?((PrimitivoContext)_localctx).FLOAT.getText():null),64)
				                                                                     if err!= nil{
				                                                                         fmt.Println(err)
				                                                                     }
				                                                                     _localctx.expr = expresion.NewPrimitivo (num,entorno.FLOAT)
				                                                                
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(587);
				((PrimitivoContext)_localctx).STRING = match(STRING);

				                                                                    str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				                                                                    _localctx.expr = expresion.NewPrimitivo(str,entorno.STRING)
				                                                                
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(589);
				((PrimitivoContext)_localctx).ID = match(ID);
				 _localctx.expr = expresion.NewIdentificador((((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getText():null))
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(591);
				match(TRUE);
				 _localctx.expr = expresion.NewPrimitivo(true,entorno.BOOLEAN)
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 6);
				{
				setState(593);
				match(FALSE);
				 _localctx.expr = expresion.NewPrimitivo(false,entorno.BOOLEAN)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return listaClases_sempred((ListaClasesContext)_localctx, predIndex);
		case 4:
			return contenidoClase_sempred((ContenidoClaseContext)_localctx, predIndex);
		case 8:
			return parametros_sempred((ParametrosContext)_localctx, predIndex);
		case 15:
			return dimensiones_sempred((DimensionesContext)_localctx, predIndex);
		case 17:
			return listanchos_sempred((ListanchosContext)_localctx, predIndex);
		case 26:
			return listaExpresiones_sempred((ListaExpresionesContext)_localctx, predIndex);
		case 30:
			return listides_sempred((ListidesContext)_localctx, predIndex);
		case 38:
			return listaAccesos_sempred((ListaAccesosContext)_localctx, predIndex);
		case 40:
			return expr_log_sempred((Expr_logContext)_localctx, predIndex);
		case 41:
			return expr_rel_sempred((Expr_relContext)_localctx, predIndex);
		case 42:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listaClases_sempred(ListaClasesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean contenidoClase_sempred(ContenidoClaseContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean parametros_sempred(ParametrosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean dimensiones_sempred(DimensionesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listanchos_sempred(ListanchosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listaExpresiones_sempred(ListaExpresionesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listides_sempred(ListidesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listaAccesos_sempred(ListaAccesosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_log_sempred(Expr_logContext _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return precpred(_ctx, 3);
		case 9:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_rel_sempred(Expr_relContext _localctx, int predIndex) {
		switch (predIndex) {
		case 10:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return precpred(_ctx, 4);
		case 12:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\62\u0258\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\7\3h\n\3"+
		"\f\3\16\3k\13\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5"+
		"\5z\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6\u0084\n\6\f\6\16\6\u0087\13"+
		"\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u0094\n\7\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\5\b\u00aa\n\b\3\t\3\t\3\t\3\t\3\t\5\t\u00b1\n\t\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\7\n\u00bc\n\n\f\n\16\n\u00bf\13\n\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\5\13\u00ca\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\r\6\r\u00da\n\r\r\r\16\r\u00db\3\r\3\r\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\5\16\u0103\n\16\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\7\21\u011a\n\21\f\21\16\21\u011d\13\21\3\22\3\22\3\22\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u012a\n\23\f\23\16\23\u012d\13"+
		"\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u0153\n\26\3\27\6\27"+
		"\u0156\n\27\r\27\16\27\u0157\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u016c\n\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\5\33\u0182\n\33\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\7\34\u018d\n\34\f\34\16\34\u0190\13\34\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\5\37"+
		"\u01a2\n\37\3 \3 \3 \3 \3 \3 \3 \3 \7 \u01ac\n \f \16 \u01af\13 \3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\5!\u01bb\n!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\5\"\u01cf\n\"\3#\3#\3#\3#\3#"+
		"\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3("+
		"\3(\3(\3(\3(\3(\7(\u01f1\n(\f(\16(\u01f4\13(\3)\3)\3)\3)\3)\5)\u01fb\n"+
		")\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\3*\7*\u020b\n*\f*\16*\u020e\13"+
		"*\3+\3+\3+\3+\3+\3+\3+\3+\3+\7+\u0219\n+\f+\16+\u021c\13+\3,\3,\3,\3,"+
		"\3,\3,\3,\3,\3,\3,\3,\3,\3,\5,\u022b\n,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,"+
		"\7,\u0237\n,\f,\16,\u023a\13,\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\5-\u0248"+
		"\n-\3.\3.\3.\3.\3.\3.\3.\3.\3.\3.\3.\3.\5.\u0256\n.\3.\2\r\4\n\22 $\66"+
		">NRTV/\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<"+
		">@BDFHJLNPRTVXZ\2\5\3\2$\'\3\2()\3\2*+\2\u0262\2\\\3\2\2\2\4_\3\2\2\2"+
		"\6l\3\2\2\2\by\3\2\2\2\n{\3\2\2\2\f\u0093\3\2\2\2\16\u00a9\3\2\2\2\20"+
		"\u00b0\3\2\2\2\22\u00b2\3\2\2\2\24\u00c9\3\2\2\2\26\u00cb\3\2\2\2\30\u00d9"+
		"\3\2\2\2\32\u0102\3\2\2\2\34\u0104\3\2\2\2\36\u010a\3\2\2\2 \u0111\3\2"+
		"\2\2\"\u011e\3\2\2\2$\u0121\3\2\2\2&\u012e\3\2\2\2(\u0133\3\2\2\2*\u0152"+
		"\3\2\2\2,\u0155\3\2\2\2.\u015b\3\2\2\2\60\u016b\3\2\2\2\62\u016d\3\2\2"+
		"\2\64\u0181\3\2\2\2\66\u0183\3\2\2\28\u0191\3\2\2\2:\u0197\3\2\2\2<\u01a1"+
		"\3\2\2\2>\u01a3\3\2\2\2@\u01ba\3\2\2\2B\u01ce\3\2\2\2D\u01d0\3\2\2\2F"+
		"\u01d5\3\2\2\2H\u01da\3\2\2\2J\u01e0\3\2\2\2L\u01e4\3\2\2\2N\u01e7\3\2"+
		"\2\2P\u01fa\3\2\2\2R\u01fc\3\2\2\2T\u020f\3\2\2\2V\u022a\3\2\2\2X\u0247"+
		"\3\2\2\2Z\u0255\3\2\2\2\\]\5\4\3\2]^\b\2\1\2^\3\3\2\2\2_`\b\3\1\2`a\5"+
		"\6\4\2ab\b\3\1\2bi\3\2\2\2cd\f\4\2\2de\5\6\4\2ef\b\3\1\2fh\3\2\2\2gc\3"+
		"\2\2\2hk\3\2\2\2ig\3\2\2\2ij\3\2\2\2j\5\3\2\2\2ki\3\2\2\2lm\7\16\2\2m"+
		"n\7\61\2\2no\5\b\5\2op\b\4\1\2p\7\3\2\2\2qr\7\5\2\2rs\5\n\6\2st\7\6\2"+
		"\2tu\b\5\1\2uz\3\2\2\2vw\7\5\2\2wx\7\6\2\2xz\b\5\1\2yq\3\2\2\2yv\3\2\2"+
		"\2z\t\3\2\2\2{|\b\6\1\2|}\5\f\7\2}~\b\6\1\2~\u0085\3\2\2\2\177\u0080\f"+
		"\4\2\2\u0080\u0081\5\f\7\2\u0081\u0082\b\6\1\2\u0082\u0084\3\2\2\2\u0083"+
		"\177\3\2\2\2\u0084\u0087\3\2\2\2\u0085\u0083\3\2\2\2\u0085\u0086\3\2\2"+
		"\2\u0086\13\3\2\2\2\u0087\u0085\3\2\2\2\u0088\u0089\5\16\b\2\u0089\u008a"+
		"\b\7\1\2\u008a\u0094\3\2\2\2\u008b\u008c\58\35\2\u008c\u008d\7\36\2\2"+
		"\u008d\u008e\b\7\1\2\u008e\u0094\3\2\2\2\u008f\u0090\5:\36\2\u0090\u0091"+
		"\7\36\2\2\u0091\u0092\b\7\1\2\u0092\u0094\3\2\2\2\u0093\u0088\3\2\2\2"+
		"\u0093\u008b\3\2\2\2\u0093\u008f\3\2\2\2\u0094\r\3\2\2\2\u0095\u0096\5"+
		"\26\f\2\u0096\u0097\b\b\1\2\u0097\u00aa\3\2\2\2\u0098\u0099\5\20\t\2\u0099"+
		"\u009a\5@!\2\u009a\u009b\7\61\2\2\u009b\u009c\7\3\2\2\u009c\u009d\7\4"+
		"\2\2\u009d\u009e\5\60\31\2\u009e\u009f\b\b\1\2\u009f\u00aa\3\2\2\2\u00a0"+
		"\u00a1\5\20\t\2\u00a1\u00a2\5@!\2\u00a2\u00a3\7\61\2\2\u00a3\u00a4\7\3"+
		"\2\2\u00a4\u00a5\5\22\n\2\u00a5\u00a6\7\4\2\2\u00a6\u00a7\5\60\31\2\u00a7"+
		"\u00a8\b\b\1\2\u00a8\u00aa\3\2\2\2\u00a9\u0095\3\2\2\2\u00a9\u0098\3\2"+
		"\2\2\u00a9\u00a0\3\2\2\2\u00aa\17\3\2\2\2\u00ab\u00ac\7\22\2\2\u00ac\u00b1"+
		"\b\t\1\2\u00ad\u00ae\7\21\2\2\u00ae\u00b1\b\t\1\2\u00af\u00b1\b\t\1\2"+
		"\u00b0\u00ab\3\2\2\2\u00b0\u00ad\3\2\2\2\u00b0\u00af\3\2\2\2\u00b1\21"+
		"\3\2\2\2\u00b2\u00b3\b\n\1\2\u00b3\u00b4\5\24\13\2\u00b4\u00b5\b\n\1\2"+
		"\u00b5\u00bd\3\2\2\2\u00b6\u00b7\f\4\2\2\u00b7\u00b8\7\35\2\2\u00b8\u00b9"+
		"\5\24\13\2\u00b9\u00ba\b\n\1\2\u00ba\u00bc\3\2\2\2\u00bb\u00b6\3\2\2\2"+
		"\u00bc\u00bf\3\2\2\2\u00bd\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\23"+
		"\3\2\2\2\u00bf\u00bd\3\2\2\2\u00c0\u00c1\5@!\2\u00c1\u00c2\7\61\2\2\u00c2"+
		"\u00c3\b\13\1\2\u00c3\u00ca\3\2\2\2\u00c4\u00c5\7(\2\2\u00c5\u00c6\5@"+
		"!\2\u00c6\u00c7\7\61\2\2\u00c7\u00c8\b\13\1\2\u00c8\u00ca\3\2\2\2\u00c9"+
		"\u00c0\3\2\2\2\u00c9\u00c4\3\2\2\2\u00ca\25\3\2\2\2\u00cb\u00cc\7\22\2"+
		"\2\u00cc\u00cd\7\23\2\2\u00cd\u00ce\7\31\2\2\u00ce\u00cf\7\20\2\2\u00cf"+
		"\u00d0\7\3\2\2\u00d0\u00d1\7\24\2\2\u00d1\u00d2\7\r\2\2\u00d2\u00d3\7"+
		"\7\2\2\u00d3\u00d4\7\b\2\2\u00d4\u00d5\7\4\2\2\u00d5\u00d6\5\60\31\2\u00d6"+
		"\u00d7\b\f\1\2\u00d7\27\3\2\2\2\u00d8\u00da\5\32\16\2\u00d9\u00d8\3\2"+
		"\2\2\u00da\u00db\3\2\2\2\u00db\u00d9\3\2\2\2\u00db\u00dc\3\2\2\2\u00dc"+
		"\u00dd\3\2\2\2\u00dd\u00de\b\r\1\2\u00de\31\3\2\2\2\u00df\u00e0\5*\26"+
		"\2\u00e0\u00e1\b\16\1\2\u00e1\u0103\3\2\2\2\u00e2\u00e3\5\62\32\2\u00e3"+
		"\u00e4\7\36\2\2\u00e4\u00e5\b\16\1\2\u00e5\u0103\3\2\2\2\u00e6\u00e7\5"+
		"8\35\2\u00e7\u00e8\7\36\2\2\u00e8\u00e9\b\16\1\2\u00e9\u0103\3\2\2\2\u00ea"+
		"\u00eb\5:\36\2\u00eb\u00ec\7\36\2\2\u00ec\u00ed\b\16\1\2\u00ed\u0103\3"+
		"\2\2\2\u00ee\u00ef\5\64\33\2\u00ef\u00f0\7\36\2\2\u00f0\u00f1\b\16\1\2"+
		"\u00f1\u0103\3\2\2\2\u00f2\u00f3\5<\37\2\u00f3\u00f4\7\36\2\2\u00f4\u00f5"+
		"\b\16\1\2\u00f5\u0103\3\2\2\2\u00f6\u00f7\5\36\20\2\u00f7\u00f8\7\36\2"+
		"\2\u00f8\u00f9\b\16\1\2\u00f9\u0103\3\2\2\2\u00fa\u00fb\5\34\17\2\u00fb"+
		"\u00fc\7\36\2\2\u00fc\u00fd\b\16\1\2\u00fd\u0103\3\2\2\2\u00fe\u00ff\5"+
		"(\25\2\u00ff\u0100\7\36\2\2\u0100\u0101\b\16\1\2\u0101\u0103\3\2\2\2\u0102"+
		"\u00df\3\2\2\2\u0102\u00e2\3\2\2\2\u0102\u00e6\3\2\2\2\u0102\u00ea\3\2"+
		"\2\2\u0102\u00ee\3\2\2\2\u0102\u00f2\3\2\2\2\u0102\u00f6\3\2\2\2\u0102"+
		"\u00fa\3\2\2\2\u0102\u00fe\3\2\2\2\u0103\33\3\2\2\2\u0104\u0105\7\61\2"+
		"\2\u0105\u0106\5> \2\u0106\u0107\7\"\2\2\u0107\u0108\5B\"\2\u0108\u0109"+
		"\b\17\1\2\u0109\35\3\2\2\2\u010a\u010b\5@!\2\u010b\u010c\5 \21\2\u010c"+
		"\u010d\7\61\2\2\u010d\u010e\7\"\2\2\u010e\u010f\5B\"\2\u010f\u0110\b\20"+
		"\1\2\u0110\37\3\2\2\2\u0111\u0112\b\21\1\2\u0112\u0113\5\"\22\2\u0113"+
		"\u0114\b\21\1\2\u0114\u011b\3\2\2\2\u0115\u0116\f\4\2\2\u0116\u0117\5"+
		"\"\22\2\u0117\u0118\b\21\1\2\u0118\u011a\3\2\2\2\u0119\u0115\3\2\2\2\u011a"+
		"\u011d\3\2\2\2\u011b\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c!\3\2\2\2"+
		"\u011d\u011b\3\2\2\2\u011e\u011f\7\7\2\2\u011f\u0120\7\b\2\2\u0120#\3"+
		"\2\2\2\u0121\u0122\b\23\1\2\u0122\u0123\5&\24\2\u0123\u0124\b\23\1\2\u0124"+
		"\u012b\3\2\2\2\u0125\u0126\f\4\2\2\u0126\u0127\5&\24\2\u0127\u0128\b\23"+
		"\1\2\u0128\u012a\3\2\2\2\u0129\u0125\3\2\2\2\u012a\u012d\3\2\2\2\u012b"+
		"\u0129\3\2\2\2\u012b\u012c\3\2\2\2\u012c%\3\2\2\2\u012d\u012b\3\2\2\2"+
		"\u012e\u012f\7\7\2\2\u012f\u0130\5B\"\2\u0130\u0131\7\b\2\2\u0131\u0132"+
		"\b\24\1\2\u0132\'\3\2\2\2\u0133\u0134\7\61\2\2\u0134\u0135\7\"\2\2\u0135"+
		"\u0136\5B\"\2\u0136\u0137\b\25\1\2\u0137)\3\2\2\2\u0138\u0139\7\13\2\2"+
		"\u0139\u013a\7\3\2\2\u013a\u013b\5B\"\2\u013b\u013c\7\4\2\2\u013c\u013d"+
		"\5\60\31\2\u013d\u013e\b\26\1\2\u013e\u0153\3\2\2\2\u013f\u0140\7\13\2"+
		"\2\u0140\u0141\7\3\2\2\u0141\u0142\5B\"\2\u0142\u0143\7\4\2\2\u0143\u0144"+
		"\5\60\31\2\u0144\u0145\7\f\2\2\u0145\u0146\5\60\31\2\u0146\u0147\b\26"+
		"\1\2\u0147\u0153\3\2\2\2\u0148\u0149\7\13\2\2\u0149\u014a\7\3\2\2\u014a"+
		"\u014b\5B\"\2\u014b\u014c\7\4\2\2\u014c\u014d\5\60\31\2\u014d\u014e\5"+
		",\27\2\u014e\u014f\7\f\2\2\u014f\u0150\5\60\31\2\u0150\u0151\b\26\1\2"+
		"\u0151\u0153\3\2\2\2\u0152\u0138\3\2\2\2\u0152\u013f\3\2\2\2\u0152\u0148"+
		"\3\2\2\2\u0153+\3\2\2\2\u0154\u0156\5.\30\2\u0155\u0154\3\2\2\2\u0156"+
		"\u0157\3\2\2\2\u0157\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158\u0159\3\2"+
		"\2\2\u0159\u015a\b\27\1\2\u015a-\3\2\2\2\u015b\u015c\7\f\2\2\u015c\u015d"+
		"\7\13\2\2\u015d\u015e\7\3\2\2\u015e\u015f\5B\"\2\u015f\u0160\7\4\2\2\u0160"+
		"\u0161\5\60\31\2\u0161\u0162\b\30\1\2\u0162/\3\2\2\2\u0163\u0164\7\5\2"+
		"\2\u0164\u0165\5\30\r\2\u0165\u0166\7\6\2\2\u0166\u0167\b\31\1\2\u0167"+
		"\u016c\3\2\2\2\u0168\u0169\7\5\2\2\u0169\u016a\7\6\2\2\u016a\u016c\b\31"+
		"\1\2\u016b\u0163\3\2\2\2\u016b\u0168\3\2\2\2\u016c\61\3\2\2\2\u016d\u016e"+
		"\7\32\2\2\u016e\u016f\7\34\2\2\u016f\u0170\7\t\2\2\u0170\u0171\7\34\2"+
		"\2\u0171\u0172\7\n\2\2\u0172\u0173\7\3\2\2\u0173\u0174\5B\"\2\u0174\u0175"+
		"\7\4\2\2\u0175\u0176\b\32\1\2\u0176\63\3\2\2\2\u0177\u0178\7\61\2\2\u0178"+
		"\u0179\7\3\2\2\u0179\u017a\7\4\2\2\u017a\u0182\b\33\1\2\u017b\u017c\7"+
		"\61\2\2\u017c\u017d\7\3\2\2\u017d\u017e\5\66\34\2\u017e\u017f\7\4\2\2"+
		"\u017f\u0180\b\33\1\2\u0180\u0182\3\2\2\2\u0181\u0177\3\2\2\2\u0181\u017b"+
		"\3\2\2\2\u0182\65\3\2\2\2\u0183\u0184\b\34\1\2\u0184\u0185\5B\"\2\u0185"+
		"\u0186\b\34\1\2\u0186\u018e\3\2\2\2\u0187\u0188\f\4\2\2\u0188\u0189\7"+
		"\35\2\2\u0189\u018a\5B\"\2\u018a\u018b\b\34\1\2\u018b\u018d\3\2\2\2\u018c"+
		"\u0187\3\2\2\2\u018d\u0190\3\2\2\2\u018e\u018c\3\2\2\2\u018e\u018f\3\2"+
		"\2\2\u018f\67\3\2\2\2\u0190\u018e\3\2\2\2\u0191\u0192\5@!\2\u0192\u0193"+
		"\5> \2\u0193\u0194\7\"\2\2\u0194\u0195\5B\"\2\u0195\u0196\b\35\1\2\u0196"+
		"9\3\2\2\2\u0197\u0198\5@!\2\u0198\u0199\5> \2\u0199\u019a\b\36\1\2\u019a"+
		";\3\2\2\2\u019b\u019c\7\25\2\2\u019c\u01a2\b\37\1\2\u019d\u019e\7\25\2"+
		"\2\u019e\u019f\5B\"\2\u019f\u01a0\b\37\1\2\u01a0\u01a2\3\2\2\2\u01a1\u019b"+
		"\3\2\2\2\u01a1\u019d\3\2\2\2\u01a2=\3\2\2\2\u01a3\u01a4\b \1\2\u01a4\u01a5"+
		"\7\61\2\2\u01a5\u01a6\b \1\2\u01a6\u01ad\3\2\2\2\u01a7\u01a8\f\4\2\2\u01a8"+
		"\u01a9\7\35\2\2\u01a9\u01aa\7\61\2\2\u01aa\u01ac\b \1\2\u01ab\u01a7\3"+
		"\2\2\2\u01ac\u01af\3\2\2\2\u01ad\u01ab\3\2\2\2\u01ad\u01ae\3\2\2\2\u01ae"+
		"?\3\2\2\2\u01af\u01ad\3\2\2\2\u01b0\u01b1\7\26\2\2\u01b1\u01bb\b!\1\2"+
		"\u01b2\u01b3\7\30\2\2\u01b3\u01bb\b!\1\2\u01b4\u01b5\7\27\2\2\u01b5\u01bb"+
		"\b!\1\2\u01b6\u01b7\7\33\2\2\u01b7\u01bb\b!\1\2\u01b8\u01b9\7\31\2\2\u01b9"+
		"\u01bb\b!\1\2\u01ba\u01b0\3\2\2\2\u01ba\u01b2\3\2\2\2\u01ba\u01b4\3\2"+
		"\2\2\u01ba\u01b6\3\2\2\2\u01ba\u01b8\3\2\2\2\u01bbA\3\2\2\2\u01bc\u01bd"+
		"\5R*\2\u01bd\u01be\b\"\1\2\u01be\u01cf\3\2\2\2\u01bf\u01c0\5T+\2\u01c0"+
		"\u01c1\b\"\1\2\u01c1\u01cf\3\2\2\2\u01c2\u01c3\5V,\2\u01c3\u01c4\b\"\1"+
		"\2\u01c4\u01cf\3\2\2\2\u01c5\u01c6\5F$\2\u01c6\u01c7\b\"\1\2\u01c7\u01cf"+
		"\3\2\2\2\u01c8\u01c9\5D#\2\u01c9\u01ca\b\"\1\2\u01ca\u01cf\3\2\2\2\u01cb"+
		"\u01cc\5H%\2\u01cc\u01cd\b\"\1\2\u01cd\u01cf\3\2\2\2\u01ce\u01bc\3\2\2"+
		"\2\u01ce\u01bf\3\2\2\2\u01ce\u01c2\3\2\2\2\u01ce\u01c5\3\2\2\2\u01ce\u01c8"+
		"\3\2\2\2\u01ce\u01cb\3\2\2\2\u01cfC\3\2\2\2\u01d0\u01d1\7\5\2\2\u01d1"+
		"\u01d2\5\66\34\2\u01d2\u01d3\7\6\2\2\u01d3\u01d4\b#\1\2\u01d4E\3\2\2\2"+
		"\u01d5\u01d6\7\17\2\2\u01d6\u01d7\5@!\2\u01d7\u01d8\5$\23\2\u01d8\u01d9"+
		"\b$\1\2\u01d9G\3\2\2\2\u01da\u01db\7\17\2\2\u01db\u01dc\7\61\2\2\u01dc"+
		"\u01dd\7\3\2\2\u01dd\u01de\7\4\2\2\u01de\u01df\b%\1\2\u01dfI\3\2\2\2\u01e0"+
		"\u01e1\7\61\2\2\u01e1\u01e2\5$\23\2\u01e2\u01e3\b&\1\2\u01e3K\3\2\2\2"+
		"\u01e4\u01e5\5N(\2\u01e5\u01e6\b\'\1\2\u01e6M\3\2\2\2\u01e7\u01e8\b(\1"+
		"\2\u01e8\u01e9\5P)\2\u01e9\u01ea\b(\1\2\u01ea\u01f2\3\2\2\2\u01eb\u01ec"+
		"\f\4\2\2\u01ec\u01ed\7\34\2\2\u01ed\u01ee\5P)\2\u01ee\u01ef\b(\1\2\u01ef"+
		"\u01f1\3\2\2\2\u01f0\u01eb\3\2\2\2\u01f1\u01f4\3\2\2\2\u01f2\u01f0\3\2"+
		"\2\2\u01f2\u01f3\3\2\2\2\u01f3O\3\2\2\2\u01f4\u01f2\3\2\2\2\u01f5\u01f6"+
		"\7\61\2\2\u01f6\u01fb\b)\1\2\u01f7\u01f8\5J&\2\u01f8\u01f9\b)\1\2\u01f9"+
		"\u01fb\3\2\2\2\u01fa\u01f5\3\2\2\2\u01fa\u01f7\3\2\2\2\u01fbQ\3\2\2\2"+
		"\u01fc\u01fd\b*\1\2\u01fd\u01fe\5T+\2\u01fe\u01ff\b*\1\2\u01ff\u020c\3"+
		"\2\2\2\u0200\u0201\f\5\2\2\u0201\u0202\7\37\2\2\u0202\u0203\5R*\6\u0203"+
		"\u0204\b*\1\2\u0204\u020b\3\2\2\2\u0205\u0206\f\4\2\2\u0206\u0207\7 \2"+
		"\2\u0207\u0208\5R*\5\u0208\u0209\b*\1\2\u0209\u020b\3\2\2\2\u020a\u0200"+
		"\3\2\2\2\u020a\u0205\3\2\2\2\u020b\u020e\3\2\2\2\u020c\u020a\3\2\2\2\u020c"+
		"\u020d\3\2\2\2\u020dS\3\2\2\2\u020e\u020c\3\2\2\2\u020f\u0210\b+\1\2\u0210"+
		"\u0211\5V,\2\u0211\u0212\b+\1\2\u0212\u021a\3\2\2\2\u0213\u0214\f\4\2"+
		"\2\u0214\u0215\t\2\2\2\u0215\u0216\5T+\5\u0216\u0217\b+\1\2\u0217\u0219"+
		"\3\2\2\2\u0218\u0213\3\2\2\2\u0219\u021c\3\2\2\2\u021a\u0218\3\2\2\2\u021a"+
		"\u021b\3\2\2\2\u021bU\3\2\2\2\u021c\u021a\3\2\2\2\u021d\u021e\b,\1\2\u021e"+
		"\u021f\7+\2\2\u021f\u0220\5B\"\2\u0220\u0221\b,\1\2\u0221\u022b\3\2\2"+
		"\2\u0222\u0223\5X-\2\u0223\u0224\b,\1\2\u0224\u022b\3\2\2\2\u0225\u0226"+
		"\7\3\2\2\u0226\u0227\5B\"\2\u0227\u0228\7\4\2\2\u0228\u0229\b,\1\2\u0229"+
		"\u022b\3\2\2\2\u022a\u021d\3\2\2\2\u022a\u0222\3\2\2\2\u022a\u0225\3\2"+
		"\2\2\u022b\u0238\3\2\2\2\u022c\u022d\f\6\2\2\u022d\u022e\t\3\2\2\u022e"+
		"\u022f\5V,\7\u022f\u0230\b,\1\2\u0230\u0237\3\2\2\2\u0231\u0232\f\5\2"+
		"\2\u0232\u0233\t\4\2\2\u0233\u0234\5V,\6\u0234\u0235\b,\1\2\u0235\u0237"+
		"\3\2\2\2\u0236\u022c\3\2\2\2\u0236\u0231\3\2\2\2\u0237\u023a\3\2\2\2\u0238"+
		"\u0236\3\2\2\2\u0238\u0239\3\2\2\2\u0239W\3\2\2\2\u023a\u0238\3\2\2\2"+
		"\u023b\u023c\5Z.\2\u023c\u023d\b-\1\2\u023d\u0248\3\2\2\2\u023e\u023f"+
		"\5\64\33\2\u023f\u0240\b-\1\2\u0240\u0248\3\2\2\2\u0241\u0242\5J&\2\u0242"+
		"\u0243\b-\1\2\u0243\u0248\3\2\2\2\u0244\u0245\5L\'\2\u0245\u0246\b-\1"+
		"\2\u0246\u0248\3\2\2\2\u0247\u023b\3\2\2\2\u0247\u023e\3\2\2\2\u0247\u0241"+
		"\3\2\2\2\u0247\u0244\3\2\2\2\u0248Y\3\2\2\2\u0249\u024a\7,\2\2\u024a\u0256"+
		"\b.\1\2\u024b\u024c\7-\2\2\u024c\u0256\b.\1\2\u024d\u024e\7.\2\2\u024e"+
		"\u0256\b.\1\2\u024f\u0250\7\61\2\2\u0250\u0256\b.\1\2\u0251\u0252\7/\2"+
		"\2\u0252\u0256\b.\1\2\u0253\u0254\7\60\2\2\u0254\u0256\b.\1\2\u0255\u0249"+
		"\3\2\2\2\u0255\u024b\3\2\2\2\u0255\u024d\3\2\2\2\u0255\u024f\3\2\2\2\u0255"+
		"\u0251\3\2\2\2\u0255\u0253\3\2\2\2\u0256[\3\2\2\2!iy\u0085\u0093\u00a9"+
		"\u00b0\u00bd\u00c9\u00db\u0102\u011b\u012b\u0152\u0157\u016b\u0181\u018e"+
		"\u01a1\u01ad\u01ba\u01ce\u01f2\u01fa\u020a\u020c\u021a\u022a\u0236\u0238"+
		"\u0247\u0255";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}