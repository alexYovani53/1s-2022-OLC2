// Generated from c:\Users\Yovani\Desktop\OLC2\analizador\Calc.g4 by ANTLR 4.8

    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/funbasica"
    import "OLC2/analizador/ast/definicion"
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
		FLOATTYPE=9, STRINGTYPE=10, BOOLTYPE=11, PUNTO=12, COMA=13, PTCOMA=14, 
		AND=15, OR=16, NOT=17, IGUAL=18, DIFERENTE=19, MAYORIGUAL=20, MENORIGUAL=21, 
		MAYOR=22, MENOR=23, MUL=24, DIV=25, ADD=26, SUB=27, NUMBER=28, FLOAT=29, 
		STRING=30, ID=31, TRUE=32, FALSE=33, WHITESPACE=34;
	public static final int
		RULE_start = 0, RULE_prueba = 1, RULE_instrucciones = 2, RULE_instruccion = 3, 
		RULE_declaracion = 4, RULE_listides = 5, RULE_tiposvars = 6, RULE_expression = 7, 
		RULE_expr_rel = 8, RULE_expr_arit = 9, RULE_expr_valor = 10, RULE_primitivo = 11;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "prueba", "instrucciones", "instruccion", "declaracion", "listides", 
			"tiposvars", "expression", "expr_rel", "expr_arit", "expr_valor", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'system'", "'out'", "'println'", "'int'", 
			"'float'", "'string'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'", 
			"'!'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", 
			"'-'", null, null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE", 
			"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", 
			"OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "ID", 
			"TRUE", "FALSE", "WHITESPACE"
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
		public *arrayList.List lista;
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
			setState(24);
			match(L_LLAVE);
			setState(25);
			((StartContext)_localctx).instrucciones = instrucciones();
			setState(26);
			match(R_LLAVE);
			_localctx.lista = ((StartContext)_localctx).instrucciones.l
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
			setState(29);
			match(L_LLAVE);
			setState(30);
			match(R_LLAVE);
			setState(31);
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
			setState(34); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(33);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(36); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << SYSTEM) | (1L << INTTYPE) | (1L << FLOATTYPE) | (1L << STRINGTYPE) | (1L << BOOLTYPE))) != 0) );

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
		public ExpressionContext expression;
		public DeclaracionContext declaracion;
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
		public TerminalNode PTCOMA() { return getToken(Calc.PTCOMA, 0); }
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
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
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SYSTEM:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				match(SYSTEM);
				setState(41);
				match(PUNTO);
				setState(42);
				match(OUT);
				setState(43);
				match(PUNTO);
				setState(44);
				match(PRINTLN);
				setState(45);
				match(LP);
				setState(46);
				((InstruccionContext)_localctx).expression = expression();
				setState(47);
				match(RP);
				setState(48);
				match(PTCOMA);
				_localctx.instr = funbasica.NewImprimir(((InstruccionContext)_localctx).expression.p)
				}
				break;
			case INTTYPE:
			case FLOATTYPE:
			case STRINGTYPE:
			case BOOLTYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				((InstruccionContext)_localctx).declaracion = declaracion();
				setState(52);
				match(PTCOMA);
				_localctx.instr = ((InstruccionContext)_localctx).declaracion.instr
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
		enterRule(_localctx, 8, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			((DeclaracionContext)_localctx).tiposvars = tiposvars();
			setState(58);
			((DeclaracionContext)_localctx).listides = listides(0);
			setState(59);
			match(IGUAL);
			setState(60);
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
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_listides, _p);

		    _localctx.lista =  arrayList.New()
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(64);
			((ListidesContext)_localctx).ID = match(ID);
			_localctx.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
			}
			_ctx.stop = _input.LT(-1);
			setState(73);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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
					setState(67);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(68);
					match(COMA);
					setState(69);
					((ListidesContext)_localctx).ID = match(ID);

					                                                                              ((ListidesContext)_localctx).sub.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
					                                                                              _localctx.lista = ((ListidesContext)_localctx).sub.lista
					                                                                          
					}
					} 
				}
				setState(75);
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
		enterRule(_localctx, 12, RULE_tiposvars);
		try {
			setState(84);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTTYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(76);
				match(INTTYPE);
				_localctx.tipo = entorno.INTEGER
				}
				break;
			case STRINGTYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(78);
				match(STRINGTYPE);
				_localctx.tipo = entorno.STRING
				}
				break;
			case FLOATTYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(80);
				match(FLOATTYPE);
				_localctx.tipo = entorno.FLOAT
				}
				break;
			case BOOLTYPE:
				enterOuterAlt(_localctx, 4);
				{
				setState(82);
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
		enterRule(_localctx, 14, RULE_expression);
		try {
			setState(92);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				((ExpressionContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((ExpressionContext)_localctx).expr_rel.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(89);
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
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_expr_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(95);
			((Expr_relContext)_localctx).expr_arit = expr_arit(0);
			_localctx.p = ((Expr_relContext)_localctx).expr_arit.p
			}
			_ctx.stop = _input.LT(-1);
			setState(105);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
					setState(98);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(99);
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
					setState(100);
					((Expr_relContext)_localctx).opDe = expr_rel(3);
					_localctx.p = expresion.NewOperacion(((Expr_relContext)_localctx).opIz.p,(((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null),((Expr_relContext)_localctx).opDe.p,false)
					}
					} 
				}
				setState(107);
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
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				setState(109);
				match(SUB);
				setState(110);
				((Expr_aritContext)_localctx).opU = ((Expr_aritContext)_localctx).expression = expression();
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opU.p,"-",nil,true)
				}
				break;
			case NUMBER:
			case FLOAT:
			case STRING:
			case ID:
			case TRUE:
			case FALSE:
				{
				setState(113);
				((Expr_aritContext)_localctx).expr_valor = expr_valor();
				_localctx.p = ((Expr_aritContext)_localctx).expr_valor.p
				}
				break;
			case LP:
				{
				setState(116);
				match(LP);
				setState(117);
				((Expr_aritContext)_localctx).expression = expression();
				setState(118);
				match(RP);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(135);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(133);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(123);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(124);
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
						setState(125);
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
						setState(128);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(129);
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
						setState(130);
						((Expr_aritContext)_localctx).opDe = expr_arit(4);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false)
						}
						break;
					}
					} 
				}
				setState(137);
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
		enterRule(_localctx, 20, RULE_expr_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
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
		public TerminalNode TRUE() { return getToken(Calc.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Calc.FALSE, 0); }
		public TerminalNode ID() { return getToken(Calc.ID, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_primitivo);
		try {
			setState(153);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
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
				setState(143);
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
				setState(145);
				((PrimitivoContext)_localctx).STRING = match(STRING);

				                                                                    str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				                                                                    _localctx.p = expresion.NewPrimitivo(str,entorno.STRING)
				                                                                
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 4);
				{
				setState(147);
				match(TRUE);
				 _localctx.p = expresion.NewPrimitivo(true,entorno.BOOLEAN)
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 5);
				{
				setState(149);
				match(FALSE);
				 _localctx.p = expresion.NewPrimitivo(false,entorno.BOOLEAN)
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(151);
				((PrimitivoContext)_localctx).ID = match(ID);
				 _localctx.p = expresion.NewIdentificador((((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getText():null))
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
		case 5:
			return listides_sempred((ListidesContext)_localctx, predIndex);
		case 8:
			return expr_rel_sempred((Expr_relContext)_localctx, predIndex);
		case 9:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3$\u009e\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4\6\4%\n\4\r"+
		"\4\16\4&\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\5\5:\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\7\7J\n\7\f\7\16\7M\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bW\n\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\5\t_\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7"+
		"\nj\n\n\f\n\16\nm\13\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\5\13|\n\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\7\13\u0088\n\13\f\13\16\13\u008b\13\13\3\f\3\f\3\f\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u009c\n\r\3\r\2\5\f\22\24"+
		"\16\2\4\6\b\n\f\16\20\22\24\26\30\2\5\3\2\26\31\3\2\32\33\3\2\34\35\2"+
		"\u00a2\2\32\3\2\2\2\4\37\3\2\2\2\6$\3\2\2\2\b9\3\2\2\2\n;\3\2\2\2\fA\3"+
		"\2\2\2\16V\3\2\2\2\20^\3\2\2\2\22`\3\2\2\2\24{\3\2\2\2\26\u008c\3\2\2"+
		"\2\30\u009b\3\2\2\2\32\33\7\5\2\2\33\34\5\6\4\2\34\35\7\6\2\2\35\36\b"+
		"\2\1\2\36\3\3\2\2\2\37 \7\5\2\2 !\7\6\2\2!\"\7\20\2\2\"\5\3\2\2\2#%\5"+
		"\b\5\2$#\3\2\2\2%&\3\2\2\2&$\3\2\2\2&\'\3\2\2\2\'(\3\2\2\2()\b\4\1\2)"+
		"\7\3\2\2\2*+\7\7\2\2+,\7\16\2\2,-\7\b\2\2-.\7\16\2\2./\7\t\2\2/\60\7\3"+
		"\2\2\60\61\5\20\t\2\61\62\7\4\2\2\62\63\7\20\2\2\63\64\b\5\1\2\64:\3\2"+
		"\2\2\65\66\5\n\6\2\66\67\7\20\2\2\678\b\5\1\28:\3\2\2\29*\3\2\2\29\65"+
		"\3\2\2\2:\t\3\2\2\2;<\5\16\b\2<=\5\f\7\2=>\7\24\2\2>?\5\20\t\2?@\b\6\1"+
		"\2@\13\3\2\2\2AB\b\7\1\2BC\7!\2\2CD\b\7\1\2DK\3\2\2\2EF\f\4\2\2FG\7\17"+
		"\2\2GH\7!\2\2HJ\b\7\1\2IE\3\2\2\2JM\3\2\2\2KI\3\2\2\2KL\3\2\2\2L\r\3\2"+
		"\2\2MK\3\2\2\2NO\7\n\2\2OW\b\b\1\2PQ\7\f\2\2QW\b\b\1\2RS\7\13\2\2SW\b"+
		"\b\1\2TU\7\r\2\2UW\b\b\1\2VN\3\2\2\2VP\3\2\2\2VR\3\2\2\2VT\3\2\2\2W\17"+
		"\3\2\2\2XY\5\22\n\2YZ\b\t\1\2Z_\3\2\2\2[\\\5\24\13\2\\]\b\t\1\2]_\3\2"+
		"\2\2^X\3\2\2\2^[\3\2\2\2_\21\3\2\2\2`a\b\n\1\2ab\5\24\13\2bc\b\n\1\2c"+
		"k\3\2\2\2de\f\4\2\2ef\t\2\2\2fg\5\22\n\5gh\b\n\1\2hj\3\2\2\2id\3\2\2\2"+
		"jm\3\2\2\2ki\3\2\2\2kl\3\2\2\2l\23\3\2\2\2mk\3\2\2\2no\b\13\1\2op\7\35"+
		"\2\2pq\5\20\t\2qr\b\13\1\2r|\3\2\2\2st\5\26\f\2tu\b\13\1\2u|\3\2\2\2v"+
		"w\7\3\2\2wx\5\20\t\2xy\7\4\2\2yz\b\13\1\2z|\3\2\2\2{n\3\2\2\2{s\3\2\2"+
		"\2{v\3\2\2\2|\u0089\3\2\2\2}~\f\6\2\2~\177\t\3\2\2\177\u0080\5\24\13\7"+
		"\u0080\u0081\b\13\1\2\u0081\u0088\3\2\2\2\u0082\u0083\f\5\2\2\u0083\u0084"+
		"\t\4\2\2\u0084\u0085\5\24\13\6\u0085\u0086\b\13\1\2\u0086\u0088\3\2\2"+
		"\2\u0087}\3\2\2\2\u0087\u0082\3\2\2\2\u0088\u008b\3\2\2\2\u0089\u0087"+
		"\3\2\2\2\u0089\u008a\3\2\2\2\u008a\25\3\2\2\2\u008b\u0089\3\2\2\2\u008c"+
		"\u008d\5\30\r\2\u008d\u008e\b\f\1\2\u008e\27\3\2\2\2\u008f\u0090\7\36"+
		"\2\2\u0090\u009c\b\r\1\2\u0091\u0092\7\37\2\2\u0092\u009c\b\r\1\2\u0093"+
		"\u0094\7 \2\2\u0094\u009c\b\r\1\2\u0095\u0096\7\"\2\2\u0096\u009c\b\r"+
		"\1\2\u0097\u0098\7#\2\2\u0098\u009c\b\r\1\2\u0099\u009a\7!\2\2\u009a\u009c"+
		"\b\r\1\2\u009b\u008f\3\2\2\2\u009b\u0091\3\2\2\2\u009b\u0093\3\2\2\2\u009b"+
		"\u0095\3\2\2\2\u009b\u0097\3\2\2\2\u009b\u0099\3\2\2\2\u009c\31\3\2\2"+
		"\2\f&9KV^k{\u0087\u0089\u009b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}