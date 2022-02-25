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
		LP=1, RP=2, L_LLAVE=3, R_LLAVE=4, L_CORCH=5, R_CORCH=6, OUT=7, PRINTLN=8, 
		IF_TOK=9, ELSE=10, ARGS=11, CLASS=12, MAIN=13, PRIVATE=14, PUBLIC=15, 
		STATIC=16, STRINGARGS=17, RETURN_P=18, INTTYPE=19, FLOATTYPE=20, STRINGTYPE=21, 
		VOIDTYPE=22, SYSTEM=23, BOOLTYPE=24, PUNTO=25, COMA=26, PTCOMA=27, AND=28, 
		OR=29, NOT=30, IGUAL=31, DIFERENTE=32, MAYORIGUAL=33, MENORIGUAL=34, MAYOR=35, 
		MENOR=36, MUL=37, DIV=38, ADD=39, SUB=40, NUMBER=41, FLOAT=42, STRING=43, 
		TRUE=44, FALSE=45, ID=46, WHITESPACE=47;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"LP", "RP", "L_LLAVE", "R_LLAVE", "L_CORCH", "R_CORCH", "OUT", "PRINTLN", 
			"IF_TOK", "ELSE", "ARGS", "CLASS", "MAIN", "PRIVATE", "PUBLIC", "STATIC", 
			"STRINGARGS", "RETURN_P", "INTTYPE", "FLOATTYPE", "STRINGTYPE", "VOIDTYPE", 
			"SYSTEM", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", "OR", "NOT", 
			"IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "TRUE", "FALSE", "ID", 
			"WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'['", "']'", "'out'", "'println'", 
			"'if'", "'else'", "'args'", "'class'", "'main'", "'private'", "'public'", 
			"'static'", "'String'", "'return'", "'int'", "'float'", "'string'", "'void'", 
			"'system'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'", "'!'", 
			"'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", 
			null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LP", "RP", "L_LLAVE", "R_LLAVE", "L_CORCH", "R_CORCH", "OUT", 
			"PRINTLN", "IF_TOK", "ELSE", "ARGS", "CLASS", "MAIN", "PRIVATE", "PUBLIC", 
			"STATIC", "STRINGARGS", "RETURN_P", "INTTYPE", "FLOATTYPE", "STRINGTYPE", 
			"VOIDTYPE", "SYSTEM", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND", "OR", 
			"NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "TRUE", "FALSE", 
			"ID", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\61\u0136\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\3\2\3\2\3\3\3\3\3\4\3\4"+
		"\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\33"+
		"\3\33\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3"+
		"!\3\"\3\"\3\"\3#\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\6*\u0103"+
		"\n*\r*\16*\u0104\3+\6+\u0108\n+\r+\16+\u0109\3+\3+\6+\u010e\n+\r+\16+"+
		"\u010f\3,\3,\7,\u0114\n,\f,\16,\u0117\13,\3,\3,\3-\3-\3-\3-\3-\3.\3.\3"+
		".\3.\3.\3.\3/\3/\7/\u0128\n/\f/\16/\u012b\13/\3\60\6\60\u012e\n\60\r\60"+
		"\16\60\u012f\3\60\3\60\3\61\3\61\3\61\2\2\62\3\3\5\4\7\5\t\6\13\7\r\b"+
		"\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26"+
		"+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S"+
		"+U,W-Y.[/]\60_\61a\2\3\2\t\3\2\62;\3\2\60\60\3\2$$\5\2C\\aac|\6\2\62;"+
		"C\\aac|\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u013a\2\3\3\2\2\2\2"+
		"\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2"+
		"\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2"+
		"\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2"+
		"\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2"+
		"\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2"+
		"\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2"+
		"K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3"+
		"\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\3c\3\2\2\2\5e\3\2\2"+
		"\2\7g\3\2\2\2\ti\3\2\2\2\13k\3\2\2\2\rm\3\2\2\2\17o\3\2\2\2\21s\3\2\2"+
		"\2\23{\3\2\2\2\25~\3\2\2\2\27\u0083\3\2\2\2\31\u0088\3\2\2\2\33\u008e"+
		"\3\2\2\2\35\u0093\3\2\2\2\37\u009b\3\2\2\2!\u00a2\3\2\2\2#\u00a9\3\2\2"+
		"\2%\u00b0\3\2\2\2\'\u00b7\3\2\2\2)\u00bb\3\2\2\2+\u00c1\3\2\2\2-\u00c8"+
		"\3\2\2\2/\u00cd\3\2\2\2\61\u00d4\3\2\2\2\63\u00dc\3\2\2\2\65\u00de\3\2"+
		"\2\2\67\u00e0\3\2\2\29\u00e2\3\2\2\2;\u00e5\3\2\2\2=\u00e8\3\2\2\2?\u00ea"+
		"\3\2\2\2A\u00ec\3\2\2\2C\u00ef\3\2\2\2E\u00f2\3\2\2\2G\u00f5\3\2\2\2I"+
		"\u00f7\3\2\2\2K\u00f9\3\2\2\2M\u00fb\3\2\2\2O\u00fd\3\2\2\2Q\u00ff\3\2"+
		"\2\2S\u0102\3\2\2\2U\u0107\3\2\2\2W\u0111\3\2\2\2Y\u011a\3\2\2\2[\u011f"+
		"\3\2\2\2]\u0125\3\2\2\2_\u012d\3\2\2\2a\u0133\3\2\2\2cd\7*\2\2d\4\3\2"+
		"\2\2ef\7+\2\2f\6\3\2\2\2gh\7}\2\2h\b\3\2\2\2ij\7\177\2\2j\n\3\2\2\2kl"+
		"\7]\2\2l\f\3\2\2\2mn\7_\2\2n\16\3\2\2\2op\7q\2\2pq\7w\2\2qr\7v\2\2r\20"+
		"\3\2\2\2st\7r\2\2tu\7t\2\2uv\7k\2\2vw\7p\2\2wx\7v\2\2xy\7n\2\2yz\7p\2"+
		"\2z\22\3\2\2\2{|\7k\2\2|}\7h\2\2}\24\3\2\2\2~\177\7g\2\2\177\u0080\7n"+
		"\2\2\u0080\u0081\7u\2\2\u0081\u0082\7g\2\2\u0082\26\3\2\2\2\u0083\u0084"+
		"\7c\2\2\u0084\u0085\7t\2\2\u0085\u0086\7i\2\2\u0086\u0087\7u\2\2\u0087"+
		"\30\3\2\2\2\u0088\u0089\7e\2\2\u0089\u008a\7n\2\2\u008a\u008b\7c\2\2\u008b"+
		"\u008c\7u\2\2\u008c\u008d\7u\2\2\u008d\32\3\2\2\2\u008e\u008f\7o\2\2\u008f"+
		"\u0090\7c\2\2\u0090\u0091\7k\2\2\u0091\u0092\7p\2\2\u0092\34\3\2\2\2\u0093"+
		"\u0094\7r\2\2\u0094\u0095\7t\2\2\u0095\u0096\7k\2\2\u0096\u0097\7x\2\2"+
		"\u0097\u0098\7c\2\2\u0098\u0099\7v\2\2\u0099\u009a\7g\2\2\u009a\36\3\2"+
		"\2\2\u009b\u009c\7r\2\2\u009c\u009d\7w\2\2\u009d\u009e\7d\2\2\u009e\u009f"+
		"\7n\2\2\u009f\u00a0\7k\2\2\u00a0\u00a1\7e\2\2\u00a1 \3\2\2\2\u00a2\u00a3"+
		"\7u\2\2\u00a3\u00a4\7v\2\2\u00a4\u00a5\7c\2\2\u00a5\u00a6\7v\2\2\u00a6"+
		"\u00a7\7k\2\2\u00a7\u00a8\7e\2\2\u00a8\"\3\2\2\2\u00a9\u00aa\7U\2\2\u00aa"+
		"\u00ab\7v\2\2\u00ab\u00ac\7t\2\2\u00ac\u00ad\7k\2\2\u00ad\u00ae\7p\2\2"+
		"\u00ae\u00af\7i\2\2\u00af$\3\2\2\2\u00b0\u00b1\7t\2\2\u00b1\u00b2\7g\2"+
		"\2\u00b2\u00b3\7v\2\2\u00b3\u00b4\7w\2\2\u00b4\u00b5\7t\2\2\u00b5\u00b6"+
		"\7p\2\2\u00b6&\3\2\2\2\u00b7\u00b8\7k\2\2\u00b8\u00b9\7p\2\2\u00b9\u00ba"+
		"\7v\2\2\u00ba(\3\2\2\2\u00bb\u00bc\7h\2\2\u00bc\u00bd\7n\2\2\u00bd\u00be"+
		"\7q\2\2\u00be\u00bf\7c\2\2\u00bf\u00c0\7v\2\2\u00c0*\3\2\2\2\u00c1\u00c2"+
		"\7u\2\2\u00c2\u00c3\7v\2\2\u00c3\u00c4\7t\2\2\u00c4\u00c5\7k\2\2\u00c5"+
		"\u00c6\7p\2\2\u00c6\u00c7\7i\2\2\u00c7,\3\2\2\2\u00c8\u00c9\7x\2\2\u00c9"+
		"\u00ca\7q\2\2\u00ca\u00cb\7k\2\2\u00cb\u00cc\7f\2\2\u00cc.\3\2\2\2\u00cd"+
		"\u00ce\7u\2\2\u00ce\u00cf\7{\2\2\u00cf\u00d0\7u\2\2\u00d0\u00d1\7v\2\2"+
		"\u00d1\u00d2\7g\2\2\u00d2\u00d3\7o\2\2\u00d3\60\3\2\2\2\u00d4\u00d5\7"+
		"d\2\2\u00d5\u00d6\7q\2\2\u00d6\u00d7\7q\2\2\u00d7\u00d8\7n\2\2\u00d8\u00d9"+
		"\7g\2\2\u00d9\u00da\7c\2\2\u00da\u00db\7p\2\2\u00db\62\3\2\2\2\u00dc\u00dd"+
		"\7\60\2\2\u00dd\64\3\2\2\2\u00de\u00df\7.\2\2\u00df\66\3\2\2\2\u00e0\u00e1"+
		"\7=\2\2\u00e18\3\2\2\2\u00e2\u00e3\7(\2\2\u00e3\u00e4\7(\2\2\u00e4:\3"+
		"\2\2\2\u00e5\u00e6\7~\2\2\u00e6\u00e7\7~\2\2\u00e7<\3\2\2\2\u00e8\u00e9"+
		"\7#\2\2\u00e9>\3\2\2\2\u00ea\u00eb\7?\2\2\u00eb@\3\2\2\2\u00ec\u00ed\7"+
		"#\2\2\u00ed\u00ee\7?\2\2\u00eeB\3\2\2\2\u00ef\u00f0\7@\2\2\u00f0\u00f1"+
		"\7?\2\2\u00f1D\3\2\2\2\u00f2\u00f3\7>\2\2\u00f3\u00f4\7?\2\2\u00f4F\3"+
		"\2\2\2\u00f5\u00f6\7@\2\2\u00f6H\3\2\2\2\u00f7\u00f8\7>\2\2\u00f8J\3\2"+
		"\2\2\u00f9\u00fa\7,\2\2\u00faL\3\2\2\2\u00fb\u00fc\7\61\2\2\u00fcN\3\2"+
		"\2\2\u00fd\u00fe\7-\2\2\u00feP\3\2\2\2\u00ff\u0100\7/\2\2\u0100R\3\2\2"+
		"\2\u0101\u0103\t\2\2\2\u0102\u0101\3\2\2\2\u0103\u0104\3\2\2\2\u0104\u0102"+
		"\3\2\2\2\u0104\u0105\3\2\2\2\u0105T\3\2\2\2\u0106\u0108\t\2\2\2\u0107"+
		"\u0106\3\2\2\2\u0108\u0109\3\2\2\2\u0109\u0107\3\2\2\2\u0109\u010a\3\2"+
		"\2\2\u010a\u010b\3\2\2\2\u010b\u010d\t\3\2\2\u010c\u010e\t\2\2\2\u010d"+
		"\u010c\3\2\2\2\u010e\u010f\3\2\2\2\u010f\u010d\3\2\2\2\u010f\u0110\3\2"+
		"\2\2\u0110V\3\2\2\2\u0111\u0115\7$\2\2\u0112\u0114\n\4\2\2\u0113\u0112"+
		"\3\2\2\2\u0114\u0117\3\2\2\2\u0115\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116"+
		"\u0118\3\2\2\2\u0117\u0115\3\2\2\2\u0118\u0119\7$\2\2\u0119X\3\2\2\2\u011a"+
		"\u011b\7v\2\2\u011b\u011c\7t\2\2\u011c\u011d\7w\2\2\u011d\u011e\7g\2\2"+
		"\u011eZ\3\2\2\2\u011f\u0120\7h\2\2\u0120\u0121\7c\2\2\u0121\u0122\7n\2"+
		"\2\u0122\u0123\7u\2\2\u0123\u0124\7g\2\2\u0124\\\3\2\2\2\u0125\u0129\t"+
		"\5\2\2\u0126\u0128\t\6\2\2\u0127\u0126\3\2\2\2\u0128\u012b\3\2\2\2\u0129"+
		"\u0127\3\2\2\2\u0129\u012a\3\2\2\2\u012a^\3\2\2\2\u012b\u0129\3\2\2\2"+
		"\u012c\u012e\t\7\2\2\u012d\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u012d"+
		"\3\2\2\2\u012f\u0130\3\2\2\2\u0130\u0131\3\2\2\2\u0131\u0132\b\60\2\2"+
		"\u0132`\3\2\2\2\u0133\u0134\7^\2\2\u0134\u0135\t\b\2\2\u0135b\3\2\2\2"+
		"\t\2\u0104\u0109\u010f\u0115\u0129\u012f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}