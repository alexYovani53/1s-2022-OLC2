// Generated from e:\Documentos\Cursos Sistemas\PracticasFinales\OLC2-LABORATORIO\EJEMPLO2\analizador\parser\Lexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Lexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LP=1, RP=2, L_LLAVE=3, R_LLAVE=4, SYSTEM=5, OUT=6, PRINTLN=7, INTTYPE=8, 
		FLOATTYPE=9, STRINGTYPE=10, BOOLTYPE=11, PUNTO=12, COMA=13, PTCOMA=14, 
		AND=15, OR=16, NOT=17, IGUAL=18, DIFERENTE=19, MAYORIGUAL=20, MENORIGUAL=21, 
		MAYOR=22, MENOR=23, MUL=24, DIV=25, ADD=26, SUB=27, NUMBER=28, FLOAT=29, 
		STRING=30, ID=31, TRUE=32, FALSE=33, WHITESPACE=34, AB=35;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE", 
			"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", 
			"OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "ID", 
			"TRUE", "FALSE", "WHITESPACE", "ESC_SEQ", "AB", "A", "B"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'system'", "'out'", "'println'", "'int'", 
			"'float'", "'string'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'", 
			"'!'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", 
			"'-'", null, null, null, null, "'true'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE", 
			"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", 
			"OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "ID", 
			"TRUE", "FALSE", "WHITESPACE", "AB"
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


	public Lexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2%\u00e6\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\3\2\3\2\3\3\3\3\3\4\3\4\3"+
		"\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3"+
		"\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3"+
		"\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3"+
		"\32\3\33\3\33\3\34\3\34\3\35\6\35\u00aa\n\35\r\35\16\35\u00ab\3\36\6\36"+
		"\u00af\n\36\r\36\16\36\u00b0\3\36\3\36\6\36\u00b5\n\36\r\36\16\36\u00b6"+
		"\3\37\3\37\7\37\u00bb\n\37\f\37\16\37\u00be\13\37\3\37\3\37\3 \3 \7 \u00c4"+
		"\n \f \16 \u00c7\13 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3"+
		"#\6#\u00d7\n#\r#\16#\u00d8\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3\'\3\'\2\2("+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37"+
		"= ?!A\"C#E$G\2I%K\2M\2\3\2\n\3\2\62;\3\2\60\60\3\2$$\5\2C\\aac|\6\2\62"+
		";C\\aac|\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\4\2CCcc\2\u00e8\2\3\3"+
		"\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2"+
		"\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3"+
		"\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61"+
		"\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2"+
		"\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2I\3\2\2\2\3O"+
		"\3\2\2\2\5Q\3\2\2\2\7S\3\2\2\2\tU\3\2\2\2\13W\3\2\2\2\r^\3\2\2\2\17b\3"+
		"\2\2\2\21j\3\2\2\2\23n\3\2\2\2\25t\3\2\2\2\27{\3\2\2\2\31\u0083\3\2\2"+
		"\2\33\u0085\3\2\2\2\35\u0087\3\2\2\2\37\u0089\3\2\2\2!\u008c\3\2\2\2#"+
		"\u008f\3\2\2\2%\u0091\3\2\2\2\'\u0093\3\2\2\2)\u0096\3\2\2\2+\u0099\3"+
		"\2\2\2-\u009c\3\2\2\2/\u009e\3\2\2\2\61\u00a0\3\2\2\2\63\u00a2\3\2\2\2"+
		"\65\u00a4\3\2\2\2\67\u00a6\3\2\2\29\u00a9\3\2\2\2;\u00ae\3\2\2\2=\u00b8"+
		"\3\2\2\2?\u00c1\3\2\2\2A\u00c8\3\2\2\2C\u00cd\3\2\2\2E\u00d6\3\2\2\2G"+
		"\u00dc\3\2\2\2I\u00df\3\2\2\2K\u00e2\3\2\2\2M\u00e4\3\2\2\2OP\7*\2\2P"+
		"\4\3\2\2\2QR\7+\2\2R\6\3\2\2\2ST\7}\2\2T\b\3\2\2\2UV\7\177\2\2V\n\3\2"+
		"\2\2WX\7u\2\2XY\7{\2\2YZ\7u\2\2Z[\7v\2\2[\\\7g\2\2\\]\7o\2\2]\f\3\2\2"+
		"\2^_\7q\2\2_`\7w\2\2`a\7v\2\2a\16\3\2\2\2bc\7r\2\2cd\7t\2\2de\7k\2\2e"+
		"f\7p\2\2fg\7v\2\2gh\7n\2\2hi\7p\2\2i\20\3\2\2\2jk\7k\2\2kl\7p\2\2lm\7"+
		"v\2\2m\22\3\2\2\2no\7h\2\2op\7n\2\2pq\7q\2\2qr\7c\2\2rs\7v\2\2s\24\3\2"+
		"\2\2tu\7u\2\2uv\7v\2\2vw\7t\2\2wx\7k\2\2xy\7p\2\2yz\7i\2\2z\26\3\2\2\2"+
		"{|\7d\2\2|}\7q\2\2}~\7q\2\2~\177\7n\2\2\177\u0080\7g\2\2\u0080\u0081\7"+
		"c\2\2\u0081\u0082\7p\2\2\u0082\30\3\2\2\2\u0083\u0084\7\60\2\2\u0084\32"+
		"\3\2\2\2\u0085\u0086\7.\2\2\u0086\34\3\2\2\2\u0087\u0088\7=\2\2\u0088"+
		"\36\3\2\2\2\u0089\u008a\7(\2\2\u008a\u008b\7(\2\2\u008b \3\2\2\2\u008c"+
		"\u008d\7~\2\2\u008d\u008e\7~\2\2\u008e\"\3\2\2\2\u008f\u0090\7#\2\2\u0090"+
		"$\3\2\2\2\u0091\u0092\7?\2\2\u0092&\3\2\2\2\u0093\u0094\7#\2\2\u0094\u0095"+
		"\7?\2\2\u0095(\3\2\2\2\u0096\u0097\7@\2\2\u0097\u0098\7?\2\2\u0098*\3"+
		"\2\2\2\u0099\u009a\7>\2\2\u009a\u009b\7?\2\2\u009b,\3\2\2\2\u009c\u009d"+
		"\7@\2\2\u009d.\3\2\2\2\u009e\u009f\7>\2\2\u009f\60\3\2\2\2\u00a0\u00a1"+
		"\7,\2\2\u00a1\62\3\2\2\2\u00a2\u00a3\7\61\2\2\u00a3\64\3\2\2\2\u00a4\u00a5"+
		"\7-\2\2\u00a5\66\3\2\2\2\u00a6\u00a7\7/\2\2\u00a78\3\2\2\2\u00a8\u00aa"+
		"\t\2\2\2\u00a9\u00a8\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ab"+
		"\u00ac\3\2\2\2\u00ac:\3\2\2\2\u00ad\u00af\t\2\2\2\u00ae\u00ad\3\2\2\2"+
		"\u00af\u00b0\3\2\2\2\u00b0\u00ae\3\2\2\2\u00b0\u00b1\3\2\2\2\u00b1\u00b2"+
		"\3\2\2\2\u00b2\u00b4\t\3\2\2\u00b3\u00b5\t\2\2\2\u00b4\u00b3\3\2\2\2\u00b5"+
		"\u00b6\3\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7<\3\2\2\2"+
		"\u00b8\u00bc\7$\2\2\u00b9\u00bb\n\4\2\2\u00ba\u00b9\3\2\2\2\u00bb\u00be"+
		"\3\2\2\2\u00bc\u00ba\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\u00bf\3\2\2\2\u00be"+
		"\u00bc\3\2\2\2\u00bf\u00c0\7$\2\2\u00c0>\3\2\2\2\u00c1\u00c5\t\5\2\2\u00c2"+
		"\u00c4\t\6\2\2\u00c3\u00c2\3\2\2\2\u00c4\u00c7\3\2\2\2\u00c5\u00c3\3\2"+
		"\2\2\u00c5\u00c6\3\2\2\2\u00c6@\3\2\2\2\u00c7\u00c5\3\2\2\2\u00c8\u00c9"+
		"\7v\2\2\u00c9\u00ca\7t\2\2\u00ca\u00cb\7w\2\2\u00cb\u00cc\7g\2\2\u00cc"+
		"B\3\2\2\2\u00cd\u00ce\7h\2\2\u00ce\u00cf\7c\2\2\u00cf\u00d0\7n\2\2\u00d0"+
		"\u00d1\7u\2\2\u00d1\u00d2\7g\2\2\u00d2\u00d3\3\2\2\2\u00d3\u00d4\5G$\2"+
		"\u00d4D\3\2\2\2\u00d5\u00d7\t\7\2\2\u00d6\u00d5\3\2\2\2\u00d7\u00d8\3"+
		"\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00da\3\2\2\2\u00da"+
		"\u00db\b#\2\2\u00dbF\3\2\2\2\u00dc\u00dd\7^\2\2\u00dd\u00de\t\b\2\2\u00de"+
		"H\3\2\2\2\u00df\u00e0\5K&\2\u00e0\u00e1\5M\'\2\u00e1J\3\2\2\2\u00e2\u00e3"+
		"\t\t\2\2\u00e3L\3\2\2\2\u00e4\u00e5\t\t\2\2\u00e5N\3\2\2\2\t\2\u00ab\u00b0"+
		"\u00b6\u00bc\u00c5\u00d8\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}