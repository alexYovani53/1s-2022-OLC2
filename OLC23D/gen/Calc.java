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
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LP=1, RP=2, L_LLAVE=3, R_LLAVE=4, L_CORCH=5, R_CORCH=6, OUT=7, PRINTLN=8, 
		IF_TOK=9, ELSE=10, PUBLIC=11, PRIVATE=12, STATIC=13, CLASS=14, MAIN=15, 
		STRINGARGS=16, ARGS=17, INTTYPE=18, FLOATTYPE=19, STRINGTYPE=20, VOIDTYPE=21, 
		SYSTEM=22, BOOLTYPE=23, PUNTO=24, COMA=25, PTCOMA=26, AND=27, OR=28, NOT=29, 
		IGUAL=30, DIFERENTE=31, MAYORIGUAL=32, MENORIGUAL=33, MAYOR=34, MENOR=35, 
		MUL=36, DIV=37, ADD=38, SUB=39, NUMBER=40, FLOAT=41, STRING=42, TRUE=43, 
		FALSE=44, ID=45, WHITESPACE=46;
	public static final int
		RULE_start = 0, RULE_listaFunciones = 1, RULE_funcionItem = 2, RULE_modaccess = 3, 
		RULE_parametros = 4, RULE_funcmain = 5, RULE_instrucciones = 6, RULE_instruccion = 7, 
		RULE_if_instr = 8, RULE_listaelseif = 9, RULE_else_if = 10, RULE_bloque = 11, 
		RULE_consola = 12, RULE_llamada = 13, RULE_listaExpresiones = 14, RULE_declaracionIni = 15, 
		RULE_declaracion = 16, RULE_listides = 17, RULE_tiposvars = 18, RULE_expression = 19, 
		RULE_expr_rel = 20, RULE_expr_arit = 21, RULE_expr_valor = 22, RULE_primitivo = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "listaFunciones", "funcionItem", "modaccess", "parametros", 
			"funcmain", "instrucciones", "instruccion", "if_instr", "listaelseif", 
			"else_if", "bloque", "consola", "llamada", "listaExpresiones", "declaracionIni", 
			"declaracion", "listides", "tiposvars", "expression", "expr_rel", "expr_arit", 
			"expr_valor", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'['", "']'", "'out'", "'println'", 
			"'if'", "'else'", "'public'", "'private'", "'static'", "'class'", "'main'", 
			"'String'", "'args'", "'int'", "'float'", "'string'", "'void'", "'system'", 
			"'boolean'", "'.'", "','", "';'", "'&&'", "'||'", "'!'", "'='", "'!='", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", null, null, 
			null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LP", "RP", "L_LLAVE", "R_LLAVE", "L_CORCH", "R_CORCH", "OUT", 
			"PRINTLN", "IF_TOK", "ELSE", "PUBLIC", "PRIVATE", "STATIC", "CLASS", 
			"MAIN", "STRINGARGS", "ARGS", "INTTYPE", "FLOATTYPE", "STRINGTYPE", "VOIDTYPE", 
			"SYSTEM", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", "OR", "NOT", 
			"IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "TRUE", "FALSE", "ID", 
			"WHITESPACE"
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


	    type  Prueba2 struct {
	        Op1 int
	        Operador string
	        Op2 int
	    }

	public Calc(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public ast.Ast ast;
		public ListaFuncionesContext listaFunciones;
		public ListaFuncionesContext listaFunciones() {
			return getRuleContext(ListaFuncionesContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitStart(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			((StartContext)_localctx).listaFunciones = listaFunciones(0);
			 _localctx.ast = ast.NewAst( ((StartContext)_localctx).listaFunciones.lista)
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

	public static class ListaFuncionesContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ListaFuncionesContext SUBLISTA;
		public FuncionItemContext funcionItem;
		public FuncionItemContext funcionItem() {
			return getRuleContext(FuncionItemContext.class,0);
		}
		public ListaFuncionesContext listaFunciones() {
			return getRuleContext(ListaFuncionesContext.class,0);
		}
		public ListaFuncionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaFunciones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterListaFunciones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitListaFunciones(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitListaFunciones(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ListaFuncionesContext listaFunciones() throws RecognitionException {
		return listaFunciones(0);
	}

	private ListaFuncionesContext listaFunciones(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListaFuncionesContext _localctx = new ListaFuncionesContext(_ctx, _parentState);
		ListaFuncionesContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_listaFunciones, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(52);
			((ListaFuncionesContext)_localctx).funcionItem = funcionItem();
			 _localctx.lista.Add( ((ListaFuncionesContext)_localctx).funcionItem.instr ) 
			}
			_ctx.stop = _input.LT(-1);
			setState(61);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListaFuncionesContext(_parentctx, _parentState);
					_localctx.SUBLISTA = _prevctx;
					_localctx.SUBLISTA = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listaFunciones);
					setState(55);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(56);
					((ListaFuncionesContext)_localctx).funcionItem = funcionItem();

					                                                          ((ListaFuncionesContext)_localctx).SUBLISTA.lista.Add( ((ListaFuncionesContext)_localctx).funcionItem.instr)
					                                                          _localctx.lista =  ((ListaFuncionesContext)_localctx).SUBLISTA.lista
					              
					}
					} 
				}
				setState(63);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitFuncionItem(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FuncionItemContext funcionItem() throws RecognitionException {
		FuncionItemContext _localctx = new FuncionItemContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_funcionItem);
		 listaParams :=  arrayList.New() 
		try {
			setState(84);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(64);
				((FuncionItemContext)_localctx).funcmain = funcmain();
				_localctx.instr =  ((FuncionItemContext)_localctx).funcmain.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(67);
				modaccess();
				setState(68);
				tiposvars();
				setState(69);
				((FuncionItemContext)_localctx).ID = match(ID);
				setState(70);
				match(LP);
				setState(71);
				match(RP);
				setState(72);
				((FuncionItemContext)_localctx).bloque = bloque();
				 _localctx.instr = Simbolos.NewFuncion((((FuncionItemContext)_localctx).ID!=null?((FuncionItemContext)_localctx).ID.getText():null),listaParams,((FuncionItemContext)_localctx).bloque.lista,entorno.VOID)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(75);
				modaccess();
				setState(76);
				((FuncionItemContext)_localctx).tiposvars = tiposvars();
				setState(77);
				((FuncionItemContext)_localctx).ID = match(ID);
				setState(78);
				match(LP);
				setState(79);
				((FuncionItemContext)_localctx).parametros = parametros(0);
				setState(80);
				match(RP);
				setState(81);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitModaccess(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ModaccessContext modaccess() throws RecognitionException {
		ModaccessContext _localctx = new ModaccessContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_modaccess);
		try {
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PUBLIC:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				match(PUBLIC);
				 _localctx.modAccess = entorno.PUBLIC
				}
				break;
			case PRIVATE:
				enterOuterAlt(_localctx, 2);
				{
				setState(88);
				match(PRIVATE);
				 _localctx.modAccess = entorno.PRIVATE
				}
				break;
			case INTTYPE:
			case FLOATTYPE:
			case STRINGTYPE:
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
		public ParametrosContext SUBLIST;
		public TiposvarsContext tiposvars;
		public Token ID;
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitParametros(this);
			else return visitor.visitChildren(this);
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
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_parametros, _p);

		_localctx.lista =  arrayList.New()
		listaIdes := arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(94);
			((ParametrosContext)_localctx).tiposvars = tiposvars();
			setState(95);
			((ParametrosContext)_localctx).ID = match(ID);

			                                                                    listaIdes.Add(expresion.NewIdentificador((((ParametrosContext)_localctx).ID!=null?((ParametrosContext)_localctx).ID.getText():null)))
			                                                                    decl := definicion.NewDeclaracion(listaIdes, ((ParametrosContext)_localctx).tiposvars.tipo)
			                                                                    _localctx.lista.Add( decl)
			                                                                 
			}
			_ctx.stop = _input.LT(-1);
			setState(106);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParametrosContext(_parentctx, _parentState);
					_localctx.SUBLIST = _prevctx;
					_localctx.SUBLIST = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_parametros);
					setState(98);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(99);
					match(COMA);
					setState(100);
					((ParametrosContext)_localctx).tiposvars = tiposvars();
					setState(101);
					((ParametrosContext)_localctx).ID = match(ID);

					                                                                              listaIdes.Add(expresion.NewIdentificador((((ParametrosContext)_localctx).ID!=null?((ParametrosContext)_localctx).ID.getText():null)))
					                                                                              decl := definicion.NewDeclaracion(listaIdes, ((ParametrosContext)_localctx).tiposvars.tipo)
					                                                                              ((ParametrosContext)_localctx).SUBLIST.lista.Add( decl )
					                                                                              _localctx.lista =  ((ParametrosContext)_localctx).SUBLIST.lista
					                                                                           
					}
					} 
				}
				setState(108);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitFuncmain(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FuncmainContext funcmain() throws RecognitionException {
		FuncmainContext _localctx = new FuncmainContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_funcmain);
		 listaParams:= arrayList.New() 
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(PUBLIC);
			setState(110);
			match(STATIC);
			setState(111);
			match(VOIDTYPE);
			setState(112);
			match(MAIN);
			setState(113);
			match(LP);
			setState(114);
			match(STRINGARGS);
			setState(115);
			match(ARGS);
			setState(116);
			match(L_CORCH);
			setState(117);
			match(R_CORCH);
			setState(118);
			match(RP);
			setState(119);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitInstrucciones(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_instrucciones);
		 _localctx.lista =  arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(122);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(125); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF_TOK) | (1L << INTTYPE) | (1L << FLOATTYPE) | (1L << STRINGTYPE) | (1L << SYSTEM) | (1L << BOOLTYPE) | (1L << ID))) != 0) );

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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitInstruccion(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_instruccion);
		try {
			setState(148);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(129);
				((InstruccionContext)_localctx).if_instr = if_instr();
				_localctx.instr = ((InstruccionContext)_localctx).if_instr.instr
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(132);
				((InstruccionContext)_localctx).consola = consola();
				setState(133);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).consola.instr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(136);
				((InstruccionContext)_localctx).declaracionIni = declaracionIni();
				setState(137);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).declaracionIni.instr
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(140);
				((InstruccionContext)_localctx).declaracion = declaracion();
				setState(141);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).declaracion.instr
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(144);
				((InstruccionContext)_localctx).llamada = llamada();
				setState(145);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).llamada.instr
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

	public static class If_instrContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public ExpressionContext expression;
		public BloqueContext bloque;
		public BloqueContext bprincipal;
		public BloqueContext belse;
		public ListaelseifContext listaelseif;
		public TerminalNode IF_TOK() { return getToken(Calc.IF_TOK, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitIf_instr(this);
			else return visitor.visitChildren(this);
		}
	}

	public final If_instrContext if_instr() throws RecognitionException {
		If_instrContext _localctx = new If_instrContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_if_instr);
		try {
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(150);
				match(IF_TOK);
				setState(151);
				match(LP);
				setState(152);
				((If_instrContext)_localctx).expression = expression();
				setState(153);
				match(RP);
				setState(154);
				((If_instrContext)_localctx).bloque = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bloque.lista,nil,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(157);
				match(IF_TOK);
				setState(158);
				match(LP);
				setState(159);
				((If_instrContext)_localctx).expression = expression();
				setState(160);
				match(RP);
				setState(161);
				((If_instrContext)_localctx).bprincipal = bloque();
				setState(162);
				match(ELSE);
				setState(163);
				((If_instrContext)_localctx).belse = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bprincipal.lista,nil,((If_instrContext)_localctx).belse.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(166);
				match(IF_TOK);
				setState(167);
				match(LP);
				setState(168);
				((If_instrContext)_localctx).expression = expression();
				setState(169);
				match(RP);
				setState(170);
				((If_instrContext)_localctx).bprincipal = bloque();
				setState(171);
				((If_instrContext)_localctx).listaelseif = listaelseif();
				setState(172);
				match(ELSE);
				setState(173);
				((If_instrContext)_localctx).belse = bloque();

				        _localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bprincipal.lista,((If_instrContext)_localctx).listaelseif.lista,((If_instrContext)_localctx).belse.lista)
				    
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitListaelseif(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ListaelseifContext listaelseif() throws RecognitionException {
		ListaelseifContext _localctx = new ListaelseifContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(179); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(178);
					((ListaelseifContext)_localctx).else_if = else_if();
					((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).else_if);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(181); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
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
		public ExpressionContext expression;
		public BloqueContext bloque;
		public TerminalNode ELSE() { return getToken(Calc.ELSE, 0); }
		public TerminalNode IF_TOK() { return getToken(Calc.IF_TOK, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitElse_if(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Else_ifContext else_if() throws RecognitionException {
		Else_ifContext _localctx = new Else_ifContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_else_if);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(185);
			match(ELSE);
			setState(186);
			match(IF_TOK);
			setState(187);
			match(LP);
			setState(188);
			((Else_ifContext)_localctx).expression = expression();
			setState(189);
			match(RP);
			setState(190);
			((Else_ifContext)_localctx).bloque = bloque();
			_localctx.instr = control.NewIfInstruccion(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).bloque.lista,nil,nil)
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitBloque(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_bloque);
		try {
			setState(201);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				match(L_LLAVE);
				setState(194);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(195);
				match(R_LLAVE);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				match(L_LLAVE);
				setState(199);
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
		public ExpressionContext expression;
		public TerminalNode SYSTEM() { return getToken(Calc.SYSTEM, 0); }
		public List<TerminalNode> PUNTO() { return getTokens(Calc.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(Calc.PUNTO, i);
		}
		public TerminalNode OUT() { return getToken(Calc.OUT, 0); }
		public TerminalNode PRINTLN() { return getToken(Calc.PRINTLN, 0); }
		public TerminalNode LP() { return getToken(Calc.LP, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitConsola(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConsolaContext consola() throws RecognitionException {
		ConsolaContext _localctx = new ConsolaContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_consola);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			match(SYSTEM);
			setState(204);
			match(PUNTO);
			setState(205);
			match(OUT);
			setState(206);
			match(PUNTO);
			setState(207);
			match(PRINTLN);
			setState(208);
			match(LP);
			setState(209);
			((ConsolaContext)_localctx).expression = expression();
			setState(210);
			match(RP);
			_localctx.instr = funbasica.NewImprimir(((ConsolaContext)_localctx).expression.p)
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitLlamada(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LlamadaContext llamada() throws RecognitionException {
		LlamadaContext _localctx = new LlamadaContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_llamada);
		try {
			setState(223);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				((LlamadaContext)_localctx).ID = match(ID);
				setState(214);
				match(LP);
				setState(215);
				match(RP);
				 _localctx.instr = expresion.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), arrayList.New())
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(217);
				((LlamadaContext)_localctx).ID = match(ID);
				setState(218);
				match(LP);
				setState(219);
				((LlamadaContext)_localctx).listaExpresiones = listaExpresiones(0);
				setState(220);
				match(RP);
				 _localctx.instr = expresion.NewLlamada((((LlamadaContext)_localctx).ID!=null?((LlamadaContext)_localctx).ID.getText():null), ((LlamadaContext)_localctx).listaExpresiones.lista)
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
		public ExpressionContext expression;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitListaExpresiones(this);
			else return visitor.visitChildren(this);
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
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_listaExpresiones, _p);

		    _localctx.lista = arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(226);
			((ListaExpresionesContext)_localctx).expression = expression();

			                                                                        _localctx.lista.Add( ((ListaExpresionesContext)_localctx).expression.p )
			                                                                     
			}
			_ctx.stop = _input.LT(-1);
			setState(236);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
					setState(229);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(230);
					match(COMA);
					setState(231);
					((ListaExpresionesContext)_localctx).expression = expression();

					                                                                                  ((ListaExpresionesContext)_localctx).LISTA.lista.Add( ((ListaExpresionesContext)_localctx).expression.p )
					                                                                                  _localctx.lista =  ((ListaExpresionesContext)_localctx).LISTA.lista
					                                                                               
					}
					} 
				}
				setState(238);
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

	public static class DeclaracionIniContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public TiposvarsContext tiposvars;
		public ListidesContext listides;
		public ExpressionContext expression;
		public TiposvarsContext tiposvars() {
			return getRuleContext(TiposvarsContext.class,0);
		}
		public ListidesContext listides() {
			return getRuleContext(ListidesContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Calc.IGUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitDeclaracionIni(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DeclaracionIniContext declaracionIni() throws RecognitionException {
		DeclaracionIniContext _localctx = new DeclaracionIniContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_declaracionIni);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			((DeclaracionIniContext)_localctx).tiposvars = tiposvars();
			setState(240);
			((DeclaracionIniContext)_localctx).listides = listides(0);
			setState(241);
			match(IGUAL);
			setState(242);
			((DeclaracionIniContext)_localctx).expression = expression();

			                                                                        _localctx.instr = definicion.NewDeclaracionInicializacion(((DeclaracionIniContext)_localctx).listides.lista,((DeclaracionIniContext)_localctx).tiposvars.tipo,((DeclaracionIniContext)_localctx).expression.p)
			                                                                     
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitDeclaracion(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			((DeclaracionContext)_localctx).tiposvars = tiposvars();
			setState(246);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitListides(this);
			else return visitor.visitChildren(this);
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
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_listides, _p);
		  _localctx.lista =  arrayList.New() 
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(250);
			((ListidesContext)_localctx).ID = match(ID);
			_localctx.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
			}
			_ctx.stop = _input.LT(-1);
			setState(259);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
					setState(253);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(254);
					match(COMA);
					setState(255);
					((ListidesContext)_localctx).ID = match(ID);

					                                                                              ((ListidesContext)_localctx).sub.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
					                                                                              _localctx.lista = ((ListidesContext)_localctx).sub.lista
					                                                                          
					}
					} 
				}
				setState(261);
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

	public static class TiposvarsContext extends ParserRuleContext {
		public entorno.TipoDato tipo;
		public TerminalNode INTTYPE() { return getToken(Calc.INTTYPE, 0); }
		public TerminalNode STRINGTYPE() { return getToken(Calc.STRINGTYPE, 0); }
		public TerminalNode FLOATTYPE() { return getToken(Calc.FLOATTYPE, 0); }
		public TerminalNode BOOLTYPE() { return getToken(Calc.BOOLTYPE, 0); }
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitTiposvars(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TiposvarsContext tiposvars() throws RecognitionException {
		TiposvarsContext _localctx = new TiposvarsContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_tiposvars);
		try {
			setState(270);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTTYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(262);
				match(INTTYPE);
				_localctx.tipo = entorno.INTEGER
				}
				break;
			case STRINGTYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(264);
				match(STRINGTYPE);
				_localctx.tipo = entorno.STRING
				}
				break;
			case FLOATTYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(266);
				match(FLOATTYPE);
				_localctx.tipo = entorno.FLOAT
				}
				break;
			case BOOLTYPE:
				enterOuterAlt(_localctx, 4);
				{
				setState(268);
				match(BOOLTYPE);
				_localctx.tipo = entorno.BOOLEAN
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

	public static class ExpressionContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_relContext expr_rel;
		public Expr_aritContext expr_arit;
		public Expr_relContext expr_rel() {
			return getRuleContext(Expr_relContext.class,0);
		}
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof CalcListener ) ((CalcListener)listener).exitExpression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expression);
		try {
			setState(278);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(272);
				((ExpressionContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((ExpressionContext)_localctx).expr_rel.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(275);
				((ExpressionContext)_localctx).expr_arit = expr_arit(0);
				_localctx.p = ((ExpressionContext)_localctx).expr_arit.p
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

	public static class Expr_relContext extends ParserRuleContext {
		public interfaces.Expresion p;
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitExpr_rel(this);
			else return visitor.visitChildren(this);
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
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_expr_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(281);
			((Expr_relContext)_localctx).expr_arit = expr_arit(0);
			_localctx.p = ((Expr_relContext)_localctx).expr_arit.p
			}
			_ctx.stop = _input.LT(-1);
			setState(291);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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
					setState(284);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(285);
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
					setState(286);
					((Expr_relContext)_localctx).opDe = expr_rel(3);
					_localctx.p = expresion.NewOperacion(((Expr_relContext)_localctx).opIz.p,(((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null),((Expr_relContext)_localctx).opDe.p,false)
					}
					} 
				}
				setState(293);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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
		public interfaces.Expresion p;
		public Expr_aritContext opIz;
		public ExpressionContext opU;
		public ExpressionContext expression;
		public Expr_valorContext expr_valor;
		public Token op;
		public Expr_aritContext opDe;
		public TerminalNode SUB() { return getToken(Calc.SUB, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitExpr_arit(this);
			else return visitor.visitChildren(this);
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
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				setState(295);
				match(SUB);
				setState(296);
				((Expr_aritContext)_localctx).opU = ((Expr_aritContext)_localctx).expression = expression();
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opU.p,"-",nil,true)
				}
				break;
			case NUMBER:
			case FLOAT:
			case STRING:
			case TRUE:
			case FALSE:
			case ID:
				{
				setState(299);
				((Expr_aritContext)_localctx).expr_valor = expr_valor();
				_localctx.p = ((Expr_aritContext)_localctx).expr_valor.p
				}
				break;
			case LP:
				{
				setState(302);
				match(LP);
				setState(303);
				((Expr_aritContext)_localctx).expression = expression();
				setState(304);
				match(RP);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(321);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(319);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(309);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(310);
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
						setState(311);
						((Expr_aritContext)_localctx).opDe = expr_arit(5);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false)
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(314);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(315);
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
						setState(316);
						((Expr_aritContext)_localctx).opDe = expr_arit(4);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false)
						}
						break;
					}
					} 
				}
				setState(323);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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
		public interfaces.Expresion p;
		public PrimitivoContext primitivo;
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitExpr_valor(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Expr_valorContext expr_valor() throws RecognitionException {
		Expr_valorContext _localctx = new Expr_valorContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_expr_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			((Expr_valorContext)_localctx).primitivo = primitivo();
			_localctx.p = ((Expr_valorContext)_localctx).primitivo.p
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
		public interfaces.Expresion p;
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
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof CalcVisitor ) return ((CalcVisitor<? extends T>)visitor).visitPrimitivo(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_primitivo);
		try {
			setState(339);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(327);
				((PrimitivoContext)_localctx).NUMBER = match(NUMBER);

				                                                                    num,err := strconv.Atoi((((PrimitivoContext)_localctx).NUMBER!=null?((PrimitivoContext)_localctx).NUMBER.getText():null))
				                                                                    if err!= nil{
				                                                                        fmt.Println(err)
				                                                                    }
				                                                                    _localctx.p = expresion.NewPrimitivo (num,entorno.INTEGER)
				                                                                
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(329);
				((PrimitivoContext)_localctx).FLOAT = match(FLOAT);

				                                                                     num,err := strconv.ParseFloat((((PrimitivoContext)_localctx).FLOAT!=null?((PrimitivoContext)_localctx).FLOAT.getText():null),64)
				                                                                     if err!= nil{
				                                                                         fmt.Println(err)
				                                                                     }
				                                                                     _localctx.p = expresion.NewPrimitivo (num,entorno.FLOAT)
				                                                                
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(331);
				((PrimitivoContext)_localctx).STRING = match(STRING);

				                                                                    str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				                                                                    _localctx.p = expresion.NewPrimitivo(str,entorno.STRING)
				                                                                
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(333);
				((PrimitivoContext)_localctx).ID = match(ID);
				 _localctx.p = expresion.NewIdentificador((((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getText():null))
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(335);
				match(TRUE);
				 _localctx.p = expresion.NewPrimitivo(true,entorno.BOOLEAN)
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 6);
				{
				setState(337);
				match(FALSE);
				 _localctx.p = expresion.NewPrimitivo(false,entorno.BOOLEAN)
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
			return listaFunciones_sempred((ListaFuncionesContext)_localctx, predIndex);
		case 4:
			return parametros_sempred((ParametrosContext)_localctx, predIndex);
		case 14:
			return listaExpresiones_sempred((ListaExpresionesContext)_localctx, predIndex);
		case 17:
			return listides_sempred((ListidesContext)_localctx, predIndex);
		case 20:
			return expr_rel_sempred((Expr_relContext)_localctx, predIndex);
		case 21:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listaFunciones_sempred(ListaFuncionesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean parametros_sempred(ParametrosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listaExpresiones_sempred(ListaExpresionesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listides_sempred(ListidesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_rel_sempred(Expr_relContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 4);
		case 6:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\60\u0158\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\7\3>\n\3\f\3\16\3A\13\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\5\4W\n\4\3\5\3\5\3\5\3\5\3\5\5\5^\n\5\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\7\6k\n\6\f\6\16\6n\13\6\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\6\b~\n\b\r\b\16\b\177\3\b\3\b\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t"+
		"\u0097\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00b3\n\n\3\13\6\13\u00b6"+
		"\n\13\r\13\16\13\u00b7\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00cc\n\r\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\5\17\u00e2\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\7\20\u00ed"+
		"\n\20\f\20\16\20\u00f0\13\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3"+
		"\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u0104\n\23\f\23"+
		"\16\23\u0107\13\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u0111"+
		"\n\24\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u0119\n\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\7\26\u0124\n\26\f\26\16\26\u0127\13\26\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0136"+
		"\n\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u0142\n\27"+
		"\f\27\16\27\u0145\13\27\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\31\3\31\5\31\u0156\n\31\3\31\2\b\4\n\36$*,\32\2\4"+
		"\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\2\5\3\2\"%\3\2&\'\3\2"+
		"()\2\u015f\2\62\3\2\2\2\4\65\3\2\2\2\6V\3\2\2\2\b]\3\2\2\2\n_\3\2\2\2"+
		"\fo\3\2\2\2\16}\3\2\2\2\20\u0096\3\2\2\2\22\u00b2\3\2\2\2\24\u00b5\3\2"+
		"\2\2\26\u00bb\3\2\2\2\30\u00cb\3\2\2\2\32\u00cd\3\2\2\2\34\u00e1\3\2\2"+
		"\2\36\u00e3\3\2\2\2 \u00f1\3\2\2\2\"\u00f7\3\2\2\2$\u00fb\3\2\2\2&\u0110"+
		"\3\2\2\2(\u0118\3\2\2\2*\u011a\3\2\2\2,\u0135\3\2\2\2.\u0146\3\2\2\2\60"+
		"\u0155\3\2\2\2\62\63\5\4\3\2\63\64\b\2\1\2\64\3\3\2\2\2\65\66\b\3\1\2"+
		"\66\67\5\6\4\2\678\b\3\1\28?\3\2\2\29:\f\4\2\2:;\5\6\4\2;<\b\3\1\2<>\3"+
		"\2\2\2=9\3\2\2\2>A\3\2\2\2?=\3\2\2\2?@\3\2\2\2@\5\3\2\2\2A?\3\2\2\2BC"+
		"\5\f\7\2CD\b\4\1\2DW\3\2\2\2EF\5\b\5\2FG\5&\24\2GH\7/\2\2HI\7\3\2\2IJ"+
		"\7\4\2\2JK\5\30\r\2KL\b\4\1\2LW\3\2\2\2MN\5\b\5\2NO\5&\24\2OP\7/\2\2P"+
		"Q\7\3\2\2QR\5\n\6\2RS\7\4\2\2ST\5\30\r\2TU\b\4\1\2UW\3\2\2\2VB\3\2\2\2"+
		"VE\3\2\2\2VM\3\2\2\2W\7\3\2\2\2XY\7\r\2\2Y^\b\5\1\2Z[\7\16\2\2[^\b\5\1"+
		"\2\\^\b\5\1\2]X\3\2\2\2]Z\3\2\2\2]\\\3\2\2\2^\t\3\2\2\2_`\b\6\1\2`a\5"+
		"&\24\2ab\7/\2\2bc\b\6\1\2cl\3\2\2\2de\f\4\2\2ef\7\33\2\2fg\5&\24\2gh\7"+
		"/\2\2hi\b\6\1\2ik\3\2\2\2jd\3\2\2\2kn\3\2\2\2lj\3\2\2\2lm\3\2\2\2m\13"+
		"\3\2\2\2nl\3\2\2\2op\7\r\2\2pq\7\17\2\2qr\7\27\2\2rs\7\21\2\2st\7\3\2"+
		"\2tu\7\22\2\2uv\7\23\2\2vw\7\7\2\2wx\7\b\2\2xy\7\4\2\2yz\5\30\r\2z{\b"+
		"\7\1\2{\r\3\2\2\2|~\5\20\t\2}|\3\2\2\2~\177\3\2\2\2\177}\3\2\2\2\177\u0080"+
		"\3\2\2\2\u0080\u0081\3\2\2\2\u0081\u0082\b\b\1\2\u0082\17\3\2\2\2\u0083"+
		"\u0084\5\22\n\2\u0084\u0085\b\t\1\2\u0085\u0097\3\2\2\2\u0086\u0087\5"+
		"\32\16\2\u0087\u0088\7\34\2\2\u0088\u0089\b\t\1\2\u0089\u0097\3\2\2\2"+
		"\u008a\u008b\5 \21\2\u008b\u008c\7\34\2\2\u008c\u008d\b\t\1\2\u008d\u0097"+
		"\3\2\2\2\u008e\u008f\5\"\22\2\u008f\u0090\7\34\2\2\u0090\u0091\b\t\1\2"+
		"\u0091\u0097\3\2\2\2\u0092\u0093\5\34\17\2\u0093\u0094\7\34\2\2\u0094"+
		"\u0095\b\t\1\2\u0095\u0097\3\2\2\2\u0096\u0083\3\2\2\2\u0096\u0086\3\2"+
		"\2\2\u0096\u008a\3\2\2\2\u0096\u008e\3\2\2\2\u0096\u0092\3\2\2\2\u0097"+
		"\21\3\2\2\2\u0098\u0099\7\13\2\2\u0099\u009a\7\3\2\2\u009a\u009b\5(\25"+
		"\2\u009b\u009c\7\4\2\2\u009c\u009d\5\30\r\2\u009d\u009e\b\n\1\2\u009e"+
		"\u00b3\3\2\2\2\u009f\u00a0\7\13\2\2\u00a0\u00a1\7\3\2\2\u00a1\u00a2\5"+
		"(\25\2\u00a2\u00a3\7\4\2\2\u00a3\u00a4\5\30\r\2\u00a4\u00a5\7\f\2\2\u00a5"+
		"\u00a6\5\30\r\2\u00a6\u00a7\b\n\1\2\u00a7\u00b3\3\2\2\2\u00a8\u00a9\7"+
		"\13\2\2\u00a9\u00aa\7\3\2\2\u00aa\u00ab\5(\25\2\u00ab\u00ac\7\4\2\2\u00ac"+
		"\u00ad\5\30\r\2\u00ad\u00ae\5\24\13\2\u00ae\u00af\7\f\2\2\u00af\u00b0"+
		"\5\30\r\2\u00b0\u00b1\b\n\1\2\u00b1\u00b3\3\2\2\2\u00b2\u0098\3\2\2\2"+
		"\u00b2\u009f\3\2\2\2\u00b2\u00a8\3\2\2\2\u00b3\23\3\2\2\2\u00b4\u00b6"+
		"\5\26\f\2\u00b5\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b5\3\2\2\2"+
		"\u00b7\u00b8\3\2\2\2\u00b8\u00b9\3\2\2\2\u00b9\u00ba\b\13\1\2\u00ba\25"+
		"\3\2\2\2\u00bb\u00bc\7\f\2\2\u00bc\u00bd\7\13\2\2\u00bd\u00be\7\3\2\2"+
		"\u00be\u00bf\5(\25\2\u00bf\u00c0\7\4\2\2\u00c0\u00c1\5\30\r\2\u00c1\u00c2"+
		"\b\f\1\2\u00c2\27\3\2\2\2\u00c3\u00c4\7\5\2\2\u00c4\u00c5\5\16\b\2\u00c5"+
		"\u00c6\7\6\2\2\u00c6\u00c7\b\r\1\2\u00c7\u00cc\3\2\2\2\u00c8\u00c9\7\5"+
		"\2\2\u00c9\u00ca\7\6\2\2\u00ca\u00cc\b\r\1\2\u00cb\u00c3\3\2\2\2\u00cb"+
		"\u00c8\3\2\2\2\u00cc\31\3\2\2\2\u00cd\u00ce\7\30\2\2\u00ce\u00cf\7\32"+
		"\2\2\u00cf\u00d0\7\t\2\2\u00d0\u00d1\7\32\2\2\u00d1\u00d2\7\n\2\2\u00d2"+
		"\u00d3\7\3\2\2\u00d3\u00d4\5(\25\2\u00d4\u00d5\7\4\2\2\u00d5\u00d6\b\16"+
		"\1\2\u00d6\33\3\2\2\2\u00d7\u00d8\7/\2\2\u00d8\u00d9\7\3\2\2\u00d9\u00da"+
		"\7\4\2\2\u00da\u00e2\b\17\1\2\u00db\u00dc\7/\2\2\u00dc\u00dd\7\3\2\2\u00dd"+
		"\u00de\5\36\20\2\u00de\u00df\7\4\2\2\u00df\u00e0\b\17\1\2\u00e0\u00e2"+
		"\3\2\2\2\u00e1\u00d7\3\2\2\2\u00e1\u00db\3\2\2\2\u00e2\35\3\2\2\2\u00e3"+
		"\u00e4\b\20\1\2\u00e4\u00e5\5(\25\2\u00e5\u00e6\b\20\1\2\u00e6\u00ee\3"+
		"\2\2\2\u00e7\u00e8\f\4\2\2\u00e8\u00e9\7\33\2\2\u00e9\u00ea\5(\25\2\u00ea"+
		"\u00eb\b\20\1\2\u00eb\u00ed\3\2\2\2\u00ec\u00e7\3\2\2\2\u00ed\u00f0\3"+
		"\2\2\2\u00ee\u00ec\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\37\3\2\2\2\u00f0"+
		"\u00ee\3\2\2\2\u00f1\u00f2\5&\24\2\u00f2\u00f3\5$\23\2\u00f3\u00f4\7 "+
		"\2\2\u00f4\u00f5\5(\25\2\u00f5\u00f6\b\21\1\2\u00f6!\3\2\2\2\u00f7\u00f8"+
		"\5&\24\2\u00f8\u00f9\5$\23\2\u00f9\u00fa\b\22\1\2\u00fa#\3\2\2\2\u00fb"+
		"\u00fc\b\23\1\2\u00fc\u00fd\7/\2\2\u00fd\u00fe\b\23\1\2\u00fe\u0105\3"+
		"\2\2\2\u00ff\u0100\f\4\2\2\u0100\u0101\7\33\2\2\u0101\u0102\7/\2\2\u0102"+
		"\u0104\b\23\1\2\u0103\u00ff\3\2\2\2\u0104\u0107\3\2\2\2\u0105\u0103\3"+
		"\2\2\2\u0105\u0106\3\2\2\2\u0106%\3\2\2\2\u0107\u0105\3\2\2\2\u0108\u0109"+
		"\7\24\2\2\u0109\u0111\b\24\1\2\u010a\u010b\7\26\2\2\u010b\u0111\b\24\1"+
		"\2\u010c\u010d\7\25\2\2\u010d\u0111\b\24\1\2\u010e\u010f\7\31\2\2\u010f"+
		"\u0111\b\24\1\2\u0110\u0108\3\2\2\2\u0110\u010a\3\2\2\2\u0110\u010c\3"+
		"\2\2\2\u0110\u010e\3\2\2\2\u0111\'\3\2\2\2\u0112\u0113\5*\26\2\u0113\u0114"+
		"\b\25\1\2\u0114\u0119\3\2\2\2\u0115\u0116\5,\27\2\u0116\u0117\b\25\1\2"+
		"\u0117\u0119\3\2\2\2\u0118\u0112\3\2\2\2\u0118\u0115\3\2\2\2\u0119)\3"+
		"\2\2\2\u011a\u011b\b\26\1\2\u011b\u011c\5,\27\2\u011c\u011d\b\26\1\2\u011d"+
		"\u0125\3\2\2\2\u011e\u011f\f\4\2\2\u011f\u0120\t\2\2\2\u0120\u0121\5*"+
		"\26\5\u0121\u0122\b\26\1\2\u0122\u0124\3\2\2\2\u0123\u011e\3\2\2\2\u0124"+
		"\u0127\3\2\2\2\u0125\u0123\3\2\2\2\u0125\u0126\3\2\2\2\u0126+\3\2\2\2"+
		"\u0127\u0125\3\2\2\2\u0128\u0129\b\27\1\2\u0129\u012a\7)\2\2\u012a\u012b"+
		"\5(\25\2\u012b\u012c\b\27\1\2\u012c\u0136\3\2\2\2\u012d\u012e\5.\30\2"+
		"\u012e\u012f\b\27\1\2\u012f\u0136\3\2\2\2\u0130\u0131\7\3\2\2\u0131\u0132"+
		"\5(\25\2\u0132\u0133\7\4\2\2\u0133\u0134\b\27\1\2\u0134\u0136\3\2\2\2"+
		"\u0135\u0128\3\2\2\2\u0135\u012d\3\2\2\2\u0135\u0130\3\2\2\2\u0136\u0143"+
		"\3\2\2\2\u0137\u0138\f\6\2\2\u0138\u0139\t\3\2\2\u0139\u013a\5,\27\7\u013a"+
		"\u013b\b\27\1\2\u013b\u0142\3\2\2\2\u013c\u013d\f\5\2\2\u013d\u013e\t"+
		"\4\2\2\u013e\u013f\5,\27\6\u013f\u0140\b\27\1\2\u0140\u0142\3\2\2\2\u0141"+
		"\u0137\3\2\2\2\u0141\u013c\3\2\2\2\u0142\u0145\3\2\2\2\u0143\u0141\3\2"+
		"\2\2\u0143\u0144\3\2\2\2\u0144-\3\2\2\2\u0145\u0143\3\2\2\2\u0146\u0147"+
		"\5\60\31\2\u0147\u0148\b\30\1\2\u0148/\3\2\2\2\u0149\u014a\7*\2\2\u014a"+
		"\u0156\b\31\1\2\u014b\u014c\7+\2\2\u014c\u0156\b\31\1\2\u014d\u014e\7"+
		",\2\2\u014e\u0156\b\31\1\2\u014f\u0150\7/\2\2\u0150\u0156\b\31\1\2\u0151"+
		"\u0152\7-\2\2\u0152\u0156\b\31\1\2\u0153\u0154\7.\2\2\u0154\u0156\b\31"+
		"\1\2\u0155\u0149\3\2\2\2\u0155\u014b\3\2\2\2\u0155\u014d\3\2\2\2\u0155"+
		"\u014f\3\2\2\2\u0155\u0151\3\2\2\2\u0155\u0153\3\2\2\2\u0156\61\3\2\2"+
		"\2\25?V]l\177\u0096\u00b2\u00b7\u00cb\u00e1\u00ee\u0105\u0110\u0118\u0125"+
		"\u0135\u0141\u0143\u0155";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}