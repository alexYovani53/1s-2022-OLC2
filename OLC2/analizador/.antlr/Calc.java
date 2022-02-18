// Generated from e:\Documentos\Cursos Sistemas\PracticasFinales\OLC2-LABORATORIO\OLC2\analizador\Calc.g4 by ANTLR 4.8

    import "OLC2/analizador/ast"
    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/funbasica"
    import "OLC2/analizador/ast/definicion"
    import "OLC2/analizador/ast/control"
    import "OLC2/analizador/entorno"
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
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LP=1, RP=2, L_LLAVE=3, R_LLAVE=4, SYSTEM=5, OUT=6, PRINTLN=7, INTTYPE=8, 
		FLOATTYPE=9, STRINGTYPE=10, BOOLTYPE=11, IF_T=12, ELSE_T=13, PUNTO=14, 
		COMA=15, PTCOMA=16, AND=17, OR=18, NOT=19, IGUAL=20, DIFERENTE=21, MAYORIGUAL=22, 
		MENORIGUAL=23, MAYOR=24, MENOR=25, MUL=26, DIV=27, ADD=28, SUB=29, NUMBER=30, 
		FLOAT=31, STRING=32, TRUE=33, FALSE=34, ID=35, WHITESPACE=36;
	public static final int
		RULE_start = 0, RULE_prueba = 1, RULE_instrucciones = 2, RULE_instruccion = 3, 
		RULE_consola = 4, RULE_declaracion = 5, RULE_if_instr = 6, RULE_listaelseif = 7, 
		RULE_else_if = 8, RULE_bloque = 9, RULE_listides = 10, RULE_tiposvars = 11, 
		RULE_expression = 12, RULE_expr_rel = 13, RULE_expr_arit = 14, RULE_expr_valor = 15, 
		RULE_primitivo = 16;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "prueba", "instrucciones", "instruccion", "consola", "declaracion", 
			"if_instr", "listaelseif", "else_if", "bloque", "listides", "tiposvars", 
			"expression", "expr_rel", "expr_arit", "expr_valor", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'system'", "'out'", "'println'", "'int'", 
			"'float'", "'string'", "'boolean'", "'if'", "'else'", "'.'", "','", "';'", 
			"'&&'", "'||'", "'!'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", 
			"'/'", "'+'", "'-'", null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE", 
			"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "IF_T", "ELSE_T", "PUNTO", "COMA", 
			"PTCOMA", "AND", "OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", 
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
		public InstruccionesContext instrucciones;
		public TerminalNode L_LLAVE() { return getToken(Calc.L_LLAVE, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode R_LLAVE() { return getToken(Calc.R_LLAVE, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			match(L_LLAVE);
			setState(35);
			((StartContext)_localctx).instrucciones = instrucciones();
			setState(36);
			match(R_LLAVE);
			_localctx.ast = ast.NewAst(((StartContext)_localctx).instrucciones.l)
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

	public static class PruebaContext extends ParserRuleContext {
		public TerminalNode L_LLAVE() { return getToken(Calc.L_LLAVE, 0); }
		public TerminalNode R_LLAVE() { return getToken(Calc.R_LLAVE, 0); }
		public TerminalNode PTCOMA() { return getToken(Calc.PTCOMA, 0); }
		public PruebaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prueba; }
	}

	public final PruebaContext prueba() throws RecognitionException {
		PruebaContext _localctx = new PruebaContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_prueba);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			match(L_LLAVE);
			setState(40);
			match(R_LLAVE);
			setState(41);
			match(PTCOMA);
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
		public *arrayList.List l;
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
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instrucciones);

		    _localctx.l =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(43);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(46); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << SYSTEM) | (1L << INTTYPE) | (1L << FLOATTYPE) | (1L << STRINGTYPE) | (1L << BOOLTYPE) | (1L << IF_T))) != 0) );

			                                                                    listInt := localctx.(*InstruccionesContext).GetE()
			                                                                        for _, e := range listInt {
			                                                                          _localctx.l.Add(e.GetInstr())
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
		public ConsolaContext consola;
		public DeclaracionContext declaracion;
		public If_instrContext if_instr;
		public ConsolaContext consola() {
			return getRuleContext(ConsolaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Calc.PTCOMA, 0); }
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public If_instrContext if_instr() {
			return getRuleContext(If_instrContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instruccion);
		try {
			setState(61);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SYSTEM:
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				((InstruccionContext)_localctx).consola = consola();
				setState(51);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).consola.instr
				}
				break;
			case INTTYPE:
			case FLOATTYPE:
			case STRINGTYPE:
			case BOOLTYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				((InstruccionContext)_localctx).declaracion = declaracion();
				setState(55);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).declaracion.instr
				}
				break;
			case IF_T:
				enterOuterAlt(_localctx, 3);
				{
				setState(58);
				((InstruccionContext)_localctx).if_instr = if_instr();
				_localctx.instr = ((InstruccionContext)_localctx).if_instr.instr
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
	}

	public final ConsolaContext consola() throws RecognitionException {
		ConsolaContext _localctx = new ConsolaContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_consola);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(63);
			match(SYSTEM);
			setState(64);
			match(PUNTO);
			setState(65);
			match(OUT);
			setState(66);
			match(PUNTO);
			setState(67);
			match(PRINTLN);
			setState(68);
			match(LP);
			setState(69);
			((ConsolaContext)_localctx).expression = expression();
			setState(70);
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

	public static class DeclaracionContext extends ParserRuleContext {
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
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			((DeclaracionContext)_localctx).tiposvars = tiposvars();
			setState(74);
			((DeclaracionContext)_localctx).listides = listides(0);
			setState(75);
			match(IGUAL);
			setState(76);
			((DeclaracionContext)_localctx).expression = expression();

			                                                                    _localctx.instr = definicion.NewDeclaracionInicializacion(((DeclaracionContext)_localctx).listides.lista,((DeclaracionContext)_localctx).tiposvars.tipo,((DeclaracionContext)_localctx).expression.p)
			                                                                
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
		public TerminalNode IF_T() { return getToken(Calc.IF_T, 0); }
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
		public TerminalNode ELSE_T() { return getToken(Calc.ELSE_T, 0); }
		public ListaelseifContext listaelseif() {
			return getRuleContext(ListaelseifContext.class,0);
		}
		public If_instrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_instr; }
	}

	public final If_instrContext if_instr() throws RecognitionException {
		If_instrContext _localctx = new If_instrContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_if_instr);
		try {
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(79);
				match(IF_T);
				setState(80);
				match(LP);
				setState(81);
				((If_instrContext)_localctx).expression = expression();
				setState(82);
				match(RP);
				setState(83);
				((If_instrContext)_localctx).bloque = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bloque.contenido,nil,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(IF_T);
				setState(87);
				match(LP);
				setState(88);
				((If_instrContext)_localctx).expression = expression();
				setState(89);
				match(RP);
				setState(90);
				((If_instrContext)_localctx).bprincipal = bloque();
				setState(91);
				match(ELSE_T);
				setState(92);
				((If_instrContext)_localctx).belse = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bprincipal.contenido,nil,((If_instrContext)_localctx).belse.contenido)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(95);
				match(IF_T);
				setState(96);
				match(LP);
				setState(97);
				((If_instrContext)_localctx).expression = expression();
				setState(98);
				match(RP);
				setState(99);
				((If_instrContext)_localctx).bprincipal = ((If_instrContext)_localctx).bloque = bloque();
				setState(100);
				((If_instrContext)_localctx).listaelseif = listaelseif();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bloque.contenido,((If_instrContext)_localctx).listaelseif.lista,nil)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(103);
				match(IF_T);
				setState(104);
				match(LP);
				setState(105);
				((If_instrContext)_localctx).expression = expression();
				setState(106);
				match(RP);
				setState(107);
				((If_instrContext)_localctx).bprincipal = bloque();
				setState(108);
				((If_instrContext)_localctx).listaelseif = listaelseif();
				setState(109);
				match(ELSE_T);
				setState(110);
				((If_instrContext)_localctx).belse = bloque();
				_localctx.instr = control.NewIfInstruccion(((If_instrContext)_localctx).expression.p,((If_instrContext)_localctx).bprincipal.contenido,((If_instrContext)_localctx).listaelseif.lista,((If_instrContext)_localctx).belse.contenido)
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
	}

	public final ListaelseifContext listaelseif() throws RecognitionException {
		ListaelseifContext _localctx = new ListaelseifContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_listaelseif);

		  _localctx.lista =  arrayList.New()

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(116); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(115);
					((ListaelseifContext)_localctx).else_if = else_if();
					((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).else_if);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(118); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );


			                          pivoteLista := localctx.(*ListaelseifContext).GetList()

			                          for _ , e := range pivoteLista {
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
		public TerminalNode ELSE_T() { return getToken(Calc.ELSE_T, 0); }
		public TerminalNode IF_T() { return getToken(Calc.IF_T, 0); }
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
	}

	public final Else_ifContext else_if() throws RecognitionException {
		Else_ifContext _localctx = new Else_ifContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_else_if);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(ELSE_T);
			setState(123);
			match(IF_T);
			setState(124);
			match(LP);
			setState(125);
			((Else_ifContext)_localctx).expression = expression();
			setState(126);
			match(RP);
			setState(127);
			((Else_ifContext)_localctx).bloque = bloque();
			_localctx.instr = control.NewIfInstruccion(((Else_ifContext)_localctx).expression.p,((Else_ifContext)_localctx).bloque.contenido,nil,nil)
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
		public *arrayList.List contenido;
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
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_bloque);
		try {
			setState(138);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(130);
				match(L_LLAVE);
				setState(131);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(132);
				match(R_LLAVE);
				_localctx.contenido = ((BloqueContext)_localctx).instrucciones.l
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(135);
				match(L_LLAVE);
				setState(136);
				match(R_LLAVE);
				_localctx.contenido = arrayList.New() 
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
	}

	public final ListidesContext listides() throws RecognitionException {
		return listides(0);
	}

	private ListidesContext listides(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListidesContext _localctx = new ListidesContext(_ctx, _parentState);
		ListidesContext _prevctx = _localctx;
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_listides, _p);

		    _localctx.lista =  arrayList.New()
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(141);
			((ListidesContext)_localctx).ID = match(ID);
			_localctx.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
			}
			_ctx.stop = _input.LT(-1);
			setState(150);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
					setState(144);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(145);
					match(COMA);
					setState(146);
					((ListidesContext)_localctx).ID = match(ID);

					                                                                              ((ListidesContext)_localctx).sub.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
					                                                                              _localctx.lista = ((ListidesContext)_localctx).sub.lista
					                                                                          
					}
					} 
				}
				setState(152);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
	}

	public final TiposvarsContext tiposvars() throws RecognitionException {
		TiposvarsContext _localctx = new TiposvarsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_tiposvars);
		try {
			setState(161);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTTYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(153);
				match(INTTYPE);
				_localctx.tipo = entorno.INTEGER
				}
				break;
			case STRINGTYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(155);
				match(STRINGTYPE);
				_localctx.tipo = entorno.STRING
				}
				break;
			case FLOATTYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(157);
				match(FLOATTYPE);
				_localctx.tipo = entorno.FLOAT
				}
				break;
			case BOOLTYPE:
				enterOuterAlt(_localctx, 4);
				{
				setState(159);
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
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_expression);
		try {
			setState(169);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				((ExpressionContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((ExpressionContext)_localctx).expr_rel.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
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
	}

	public final Expr_relContext expr_rel() throws RecognitionException {
		return expr_rel(0);
	}

	private Expr_relContext expr_rel(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_relContext _localctx = new Expr_relContext(_ctx, _parentState);
		Expr_relContext _prevctx = _localctx;
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_expr_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(172);
			((Expr_relContext)_localctx).expr_arit = expr_arit(0);
			_localctx.p = ((Expr_relContext)_localctx).expr_arit.p
			}
			_ctx.stop = _input.LT(-1);
			setState(182);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
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
					setState(175);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(176);
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
					setState(177);
					((Expr_relContext)_localctx).opDe = expr_rel(3);
					_localctx.p = expresion.NewOperacion(((Expr_relContext)_localctx).opIz.p,(((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null),((Expr_relContext)_localctx).opDe.p,false)
					}
					} 
				}
				setState(184);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
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
	}

	public final Expr_aritContext expr_arit() throws RecognitionException {
		return expr_arit(0);
	}

	private Expr_aritContext expr_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_aritContext _localctx = new Expr_aritContext(_ctx, _parentState);
		Expr_aritContext _prevctx = _localctx;
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				setState(186);
				match(SUB);
				setState(187);
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
				setState(190);
				((Expr_aritContext)_localctx).expr_valor = expr_valor();
				_localctx.p = ((Expr_aritContext)_localctx).expr_valor.p
				}
				break;
			case LP:
				{
				setState(193);
				match(LP);
				setState(194);
				((Expr_aritContext)_localctx).expression = expression();
				setState(195);
				match(RP);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(212);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(210);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(200);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(201);
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
						setState(202);
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
						setState(205);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(206);
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
						setState(207);
						((Expr_aritContext)_localctx).opDe = expr_arit(4);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false)
						}
						break;
					}
					} 
				}
				setState(214);
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
	}

	public final Expr_valorContext expr_valor() throws RecognitionException {
		Expr_valorContext _localctx = new Expr_valorContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_expr_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
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
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_primitivo);
		try {
			setState(230);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(218);
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
				setState(220);
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
				setState(222);
				((PrimitivoContext)_localctx).STRING = match(STRING);

				                                                                    str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				                                                                    _localctx.p = expresion.NewPrimitivo(str,entorno.STRING)
				                                                                
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(224);
				((PrimitivoContext)_localctx).ID = match(ID);
				 _localctx.p = expresion.NewIdentificador((((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getText():null))
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(226);
				match(TRUE);
				 _localctx.p = expresion.NewPrimitivo(true,entorno.BOOLEAN)
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 6);
				{
				setState(228);
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
		case 10:
			return listides_sempred((ListidesContext)_localctx, predIndex);
		case 13:
			return expr_rel_sempred((Expr_relContext)_localctx, predIndex);
		case 14:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listides_sempred(ListidesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_rel_sempred(Expr_relContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3&\u00eb\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4\6\4/\n\4\r\4\16\4\60\3\4\3\4"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5@\n\5\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bt\n\b\3\t\6\tw\n\t\r\t"+
		"\16\tx\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\5\13\u008d\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f\u0097"+
		"\n\f\f\f\16\f\u009a\13\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00a4\n\r"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00ac\n\16\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\7\17\u00b7\n\17\f\17\16\17\u00ba\13\17\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00c9\n\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\7\20\u00d5\n\20\f\20"+
		"\16\20\u00d8\13\20\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\3\22\5\22\u00e9\n\22\3\22\2\5\26\34\36\23\2\4\6\b\n"+
		"\f\16\20\22\24\26\30\32\34\36 \"\2\5\3\2\30\33\3\2\34\35\3\2\36\37\2\u00f0"+
		"\2$\3\2\2\2\4)\3\2\2\2\6.\3\2\2\2\b?\3\2\2\2\nA\3\2\2\2\fK\3\2\2\2\16"+
		"s\3\2\2\2\20v\3\2\2\2\22|\3\2\2\2\24\u008c\3\2\2\2\26\u008e\3\2\2\2\30"+
		"\u00a3\3\2\2\2\32\u00ab\3\2\2\2\34\u00ad\3\2\2\2\36\u00c8\3\2\2\2 \u00d9"+
		"\3\2\2\2\"\u00e8\3\2\2\2$%\7\5\2\2%&\5\6\4\2&\'\7\6\2\2\'(\b\2\1\2(\3"+
		"\3\2\2\2)*\7\5\2\2*+\7\6\2\2+,\7\22\2\2,\5\3\2\2\2-/\5\b\5\2.-\3\2\2\2"+
		"/\60\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61\62\3\2\2\2\62\63\b\4\1\2\63"+
		"\7\3\2\2\2\64\65\5\n\6\2\65\66\7\22\2\2\66\67\b\5\1\2\67@\3\2\2\289\5"+
		"\f\7\29:\7\22\2\2:;\b\5\1\2;@\3\2\2\2<=\5\16\b\2=>\b\5\1\2>@\3\2\2\2?"+
		"\64\3\2\2\2?8\3\2\2\2?<\3\2\2\2@\t\3\2\2\2AB\7\7\2\2BC\7\20\2\2CD\7\b"+
		"\2\2DE\7\20\2\2EF\7\t\2\2FG\7\3\2\2GH\5\32\16\2HI\7\4\2\2IJ\b\6\1\2J\13"+
		"\3\2\2\2KL\5\30\r\2LM\5\26\f\2MN\7\26\2\2NO\5\32\16\2OP\b\7\1\2P\r\3\2"+
		"\2\2QR\7\16\2\2RS\7\3\2\2ST\5\32\16\2TU\7\4\2\2UV\5\24\13\2VW\b\b\1\2"+
		"Wt\3\2\2\2XY\7\16\2\2YZ\7\3\2\2Z[\5\32\16\2[\\\7\4\2\2\\]\5\24\13\2]^"+
		"\7\17\2\2^_\5\24\13\2_`\b\b\1\2`t\3\2\2\2ab\7\16\2\2bc\7\3\2\2cd\5\32"+
		"\16\2de\7\4\2\2ef\5\24\13\2fg\5\20\t\2gh\b\b\1\2ht\3\2\2\2ij\7\16\2\2"+
		"jk\7\3\2\2kl\5\32\16\2lm\7\4\2\2mn\5\24\13\2no\5\20\t\2op\7\17\2\2pq\5"+
		"\24\13\2qr\b\b\1\2rt\3\2\2\2sQ\3\2\2\2sX\3\2\2\2sa\3\2\2\2si\3\2\2\2t"+
		"\17\3\2\2\2uw\5\22\n\2vu\3\2\2\2wx\3\2\2\2xv\3\2\2\2xy\3\2\2\2yz\3\2\2"+
		"\2z{\b\t\1\2{\21\3\2\2\2|}\7\17\2\2}~\7\16\2\2~\177\7\3\2\2\177\u0080"+
		"\5\32\16\2\u0080\u0081\7\4\2\2\u0081\u0082\5\24\13\2\u0082\u0083\b\n\1"+
		"\2\u0083\23\3\2\2\2\u0084\u0085\7\5\2\2\u0085\u0086\5\6\4\2\u0086\u0087"+
		"\7\6\2\2\u0087\u0088\b\13\1\2\u0088\u008d\3\2\2\2\u0089\u008a\7\5\2\2"+
		"\u008a\u008b\7\6\2\2\u008b\u008d\b\13\1\2\u008c\u0084\3\2\2\2\u008c\u0089"+
		"\3\2\2\2\u008d\25\3\2\2\2\u008e\u008f\b\f\1\2\u008f\u0090\7%\2\2\u0090"+
		"\u0091\b\f\1\2\u0091\u0098\3\2\2\2\u0092\u0093\f\4\2\2\u0093\u0094\7\21"+
		"\2\2\u0094\u0095\7%\2\2\u0095\u0097\b\f\1\2\u0096\u0092\3\2\2\2\u0097"+
		"\u009a\3\2\2\2\u0098\u0096\3\2\2\2\u0098\u0099\3\2\2\2\u0099\27\3\2\2"+
		"\2\u009a\u0098\3\2\2\2\u009b\u009c\7\n\2\2\u009c\u00a4\b\r\1\2\u009d\u009e"+
		"\7\f\2\2\u009e\u00a4\b\r\1\2\u009f\u00a0\7\13\2\2\u00a0\u00a4\b\r\1\2"+
		"\u00a1\u00a2\7\r\2\2\u00a2\u00a4\b\r\1\2\u00a3\u009b\3\2\2\2\u00a3\u009d"+
		"\3\2\2\2\u00a3\u009f\3\2\2\2\u00a3\u00a1\3\2\2\2\u00a4\31\3\2\2\2\u00a5"+
		"\u00a6\5\34\17\2\u00a6\u00a7\b\16\1\2\u00a7\u00ac\3\2\2\2\u00a8\u00a9"+
		"\5\36\20\2\u00a9\u00aa\b\16\1\2\u00aa\u00ac\3\2\2\2\u00ab\u00a5\3\2\2"+
		"\2\u00ab\u00a8\3\2\2\2\u00ac\33\3\2\2\2\u00ad\u00ae\b\17\1\2\u00ae\u00af"+
		"\5\36\20\2\u00af\u00b0\b\17\1\2\u00b0\u00b8\3\2\2\2\u00b1\u00b2\f\4\2"+
		"\2\u00b2\u00b3\t\2\2\2\u00b3\u00b4\5\34\17\5\u00b4\u00b5\b\17\1\2\u00b5"+
		"\u00b7\3\2\2\2\u00b6\u00b1\3\2\2\2\u00b7\u00ba\3\2\2\2\u00b8\u00b6\3\2"+
		"\2\2\u00b8\u00b9\3\2\2\2\u00b9\35\3\2\2\2\u00ba\u00b8\3\2\2\2\u00bb\u00bc"+
		"\b\20\1\2\u00bc\u00bd\7\37\2\2\u00bd\u00be\5\32\16\2\u00be\u00bf\b\20"+
		"\1\2\u00bf\u00c9\3\2\2\2\u00c0\u00c1\5 \21\2\u00c1\u00c2\b\20\1\2\u00c2"+
		"\u00c9\3\2\2\2\u00c3\u00c4\7\3\2\2\u00c4\u00c5\5\32\16\2\u00c5\u00c6\7"+
		"\4\2\2\u00c6\u00c7\b\20\1\2\u00c7\u00c9\3\2\2\2\u00c8\u00bb\3\2\2\2\u00c8"+
		"\u00c0\3\2\2\2\u00c8\u00c3\3\2\2\2\u00c9\u00d6\3\2\2\2\u00ca\u00cb\f\6"+
		"\2\2\u00cb\u00cc\t\3\2\2\u00cc\u00cd\5\36\20\7\u00cd\u00ce\b\20\1\2\u00ce"+
		"\u00d5\3\2\2\2\u00cf\u00d0\f\5\2\2\u00d0\u00d1\t\4\2\2\u00d1\u00d2\5\36"+
		"\20\6\u00d2\u00d3\b\20\1\2\u00d3\u00d5\3\2\2\2\u00d4\u00ca\3\2\2\2\u00d4"+
		"\u00cf\3\2\2\2\u00d5\u00d8\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d6\u00d7\3\2"+
		"\2\2\u00d7\37\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d9\u00da\5\"\22\2\u00da\u00db"+
		"\b\21\1\2\u00db!\3\2\2\2\u00dc\u00dd\7 \2\2\u00dd\u00e9\b\22\1\2\u00de"+
		"\u00df\7!\2\2\u00df\u00e9\b\22\1\2\u00e0\u00e1\7\"\2\2\u00e1\u00e9\b\22"+
		"\1\2\u00e2\u00e3\7%\2\2\u00e3\u00e9\b\22\1\2\u00e4\u00e5\7#\2\2\u00e5"+
		"\u00e9\b\22\1\2\u00e6\u00e7\7$\2\2\u00e7\u00e9\b\22\1\2\u00e8\u00dc\3"+
		"\2\2\2\u00e8\u00de\3\2\2\2\u00e8\u00e0\3\2\2\2\u00e8\u00e2\3\2\2\2\u00e8"+
		"\u00e4\3\2\2\2\u00e8\u00e6\3\2\2\2\u00e9#\3\2\2\2\17\60?sx\u008c\u0098"+
		"\u00a3\u00ab\u00b8\u00c8\u00d4\u00d6\u00e8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}