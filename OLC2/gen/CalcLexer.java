// Generated from C:/Users/Yovani/Desktop/OLC2/analizador\CalcLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CalcLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		SYSTEM=1, OUT=2, PRINTLN=3, NUMBER=4, STRING=5, PUNTO=6, PTCOMA=7, DIFERENTE=8, 
		MAYORIGUAL=9, MENORIGUAL=10, MAYOR=11, MENOR=12, MUL=13, DIV=14, ADD=15, 
		SUB=16, LP=17, RP=18, L_LLAVE=19, R_LLAVE=20, WHITESPACE=21;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"SYSTEM", "OUT", "PRINTLN", "NUMBER", "STRING", "PUNTO", "PTCOMA", "DIFERENTE", 
			"MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", 
			"LP", "RP", "L_LLAVE", "R_LLAVE", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'system'", "'out'", "'println'", null, null, "'.'", "';'", "'!'", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", 
			"'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "SYSTEM", "OUT", "PRINTLN", "NUMBER", "STRING", "PUNTO", "PTCOMA", 
			"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", 
			"ADD", "SUB", "LP", "RP", "L_LLAVE", "R_LLAVE", "WHITESPACE"
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


	public CalcLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "CalcLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\27z\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\6\5D\n\5"+
		"\r\5\16\5E\3\6\3\6\7\6J\n\6\f\6\16\6M\13\6\3\6\3\6\3\7\3\7\3\b\3\b\3\t"+
		"\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3"+
		"\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\6\26r"+
		"\n\26\r\26\16\26s\3\26\3\26\3\27\3\27\3\27\2\2\30\3\3\5\4\7\5\t\6\13\7"+
		"\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\2\3\2\6\3\2\62;\3\2$$\6\2\13\f\17\17\"\"^^\t\2\"#%%--/\60<<"+
		"BB]_\2{\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2"+
		"\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3"+
		"\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2"+
		"\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\3/\3\2\2\2\5"+
		"\66\3\2\2\2\7:\3\2\2\2\tC\3\2\2\2\13G\3\2\2\2\rP\3\2\2\2\17R\3\2\2\2\21"+
		"T\3\2\2\2\23V\3\2\2\2\25Y\3\2\2\2\27\\\3\2\2\2\31^\3\2\2\2\33`\3\2\2\2"+
		"\35b\3\2\2\2\37d\3\2\2\2!f\3\2\2\2#h\3\2\2\2%j\3\2\2\2\'l\3\2\2\2)n\3"+
		"\2\2\2+q\3\2\2\2-w\3\2\2\2/\60\7u\2\2\60\61\7{\2\2\61\62\7u\2\2\62\63"+
		"\7v\2\2\63\64\7g\2\2\64\65\7o\2\2\65\4\3\2\2\2\66\67\7q\2\2\678\7w\2\2"+
		"89\7v\2\29\6\3\2\2\2:;\7r\2\2;<\7t\2\2<=\7k\2\2=>\7p\2\2>?\7v\2\2?@\7"+
		"n\2\2@A\7p\2\2A\b\3\2\2\2BD\t\2\2\2CB\3\2\2\2DE\3\2\2\2EC\3\2\2\2EF\3"+
		"\2\2\2F\n\3\2\2\2GK\7$\2\2HJ\n\3\2\2IH\3\2\2\2JM\3\2\2\2KI\3\2\2\2KL\3"+
		"\2\2\2LN\3\2\2\2MK\3\2\2\2NO\7$\2\2O\f\3\2\2\2PQ\7\60\2\2Q\16\3\2\2\2"+
		"RS\7=\2\2S\20\3\2\2\2TU\7#\2\2U\22\3\2\2\2VW\7@\2\2WX\7?\2\2X\24\3\2\2"+
		"\2YZ\7>\2\2Z[\7?\2\2[\26\3\2\2\2\\]\7@\2\2]\30\3\2\2\2^_\7>\2\2_\32\3"+
		"\2\2\2`a\7,\2\2a\34\3\2\2\2bc\7\61\2\2c\36\3\2\2\2de\7-\2\2e \3\2\2\2"+
		"fg\7/\2\2g\"\3\2\2\2hi\7*\2\2i$\3\2\2\2jk\7+\2\2k&\3\2\2\2lm\7}\2\2m("+
		"\3\2\2\2no\7\177\2\2o*\3\2\2\2pr\t\4\2\2qp\3\2\2\2rs\3\2\2\2sq\3\2\2\2"+
		"st\3\2\2\2tu\3\2\2\2uv\b\26\2\2v,\3\2\2\2wx\7^\2\2xy\t\5\2\2y.\3\2\2\2"+
		"\6\2EKs\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}