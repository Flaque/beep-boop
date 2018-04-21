// Generated from /home/flaque/projects/go/src/github.com/Flaque/boop/parser/BeepBoop.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BoopLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, COMMENT=2, NEWLINE=3, WHITESPACE=4, EQUALS=5, ASSIGN=6, PIPE=7, 
		LTHAN=8, GTHAN=9, LTHAN_EQUALS=10, GTHAN_EQUALS=11, LPAREN=12, RPAREN=13, 
		LSQUIG=14, RSQUIG=15, LBLOCK=16, RBLOCK=17, IF=18, DO=19, END=20, FUNC=21, 
		RETURN=22, FOR=23, PLUS=24, SUB=25, DIV=26, MULT=27, TRUE=28, FALSE=29, 
		OR=30, AND=31, INT=32, QUOTED=33, LETTERS=34, STRING=35;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE", 
		"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", 
		"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "IF", "DO", "END", "FUNC", "RETURN", 
		"FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", 
		"QUOTED", "STRINGORNEW", "LETTERS", "STRING", "ESC"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "':'", null, null, null, "'=='", "'='", "'|'", "'<'", "'>'", "'<='", 
		"'>='", "'('", "')'", "'{'", "'}'", "'['", "']'", "'if'", "'do'", "'end'", 
		"'func'", "'return'", "'for'", "'add'", "'sub'", "'div'", "'mult'", "'true'", 
		"'false'", "'or'", "'and'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE", 
		"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", 
		"LSQUIG", "RSQUIG", "LBLOCK", "RBLOCK", "IF", "DO", "END", "FUNC", "RETURN", 
		"FOR", "PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", 
		"QUOTED", "LETTERS", "STRING"
	};
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


	public BoopLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "BeepBoop.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2%\u00e5\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\3\2\3\2\3\3\3\3\7\3R\n\3\f\3\16\3"+
		"U\13\3\3\3\3\3\3\4\3\4\6\4[\n\4\r\4\16\4\\\5\4_\n\4\3\5\6\5b\n\5\r\5\16"+
		"\5c\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\13"+
		"\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22"+
		"\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3\31"+
		"\3\31\3\31\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34"+
		"\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37"+
		"\3\37\3 \3 \3 \3 \3!\6!\u00c3\n!\r!\16!\u00c4\3\"\3\"\6\"\u00c9\n\"\r"+
		"\"\16\"\u00ca\7\"\u00cd\n\"\f\"\16\"\u00d0\13\"\3\"\3\"\3#\3#\5#\u00d6"+
		"\n#\3$\6$\u00d9\n$\r$\16$\u00da\3%\3%\5%\u00df\n%\3&\3&\3&\5&\u00e4\n"+
		"&\2\2\'\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17"+
		"\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\35"+
		"9\36;\37= ?!A\"C#E\2G$I%K\2\3\2\b\4\2\f\f\17\17\4\2\13\13\"\"\3\2\62;"+
		"\4\2C\\c|\4\2$$^^\n\2$$\61\61^^ddhhppttvv\2\u00ed\2\3\3\2\2\2\2\5\3\2"+
		"\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\3M\3\2\2\2\5O\3\2\2"+
		"\2\7^\3\2\2\2\ta\3\2\2\2\13g\3\2\2\2\rj\3\2\2\2\17l\3\2\2\2\21n\3\2\2"+
		"\2\23p\3\2\2\2\25r\3\2\2\2\27u\3\2\2\2\31x\3\2\2\2\33z\3\2\2\2\35|\3\2"+
		"\2\2\37~\3\2\2\2!\u0080\3\2\2\2#\u0082\3\2\2\2%\u0084\3\2\2\2\'\u0087"+
		"\3\2\2\2)\u008a\3\2\2\2+\u008e\3\2\2\2-\u0093\3\2\2\2/\u009a\3\2\2\2\61"+
		"\u009e\3\2\2\2\63\u00a2\3\2\2\2\65\u00a6\3\2\2\2\67\u00aa\3\2\2\29\u00af"+
		"\3\2\2\2;\u00b4\3\2\2\2=\u00ba\3\2\2\2?\u00bd\3\2\2\2A\u00c2\3\2\2\2C"+
		"\u00c6\3\2\2\2E\u00d5\3\2\2\2G\u00d8\3\2\2\2I\u00de\3\2\2\2K\u00e0\3\2"+
		"\2\2MN\7<\2\2N\4\3\2\2\2OS\7%\2\2PR\n\2\2\2QP\3\2\2\2RU\3\2\2\2SQ\3\2"+
		"\2\2ST\3\2\2\2TV\3\2\2\2US\3\2\2\2VW\b\3\2\2W\6\3\2\2\2X_\t\2\2\2Y[\t"+
		"\2\2\2ZY\3\2\2\2[\\\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]_\3\2\2\2^X\3\2\2\2"+
		"^Z\3\2\2\2_\b\3\2\2\2`b\t\3\2\2a`\3\2\2\2bc\3\2\2\2ca\3\2\2\2cd\3\2\2"+
		"\2de\3\2\2\2ef\b\5\3\2f\n\3\2\2\2gh\7?\2\2hi\7?\2\2i\f\3\2\2\2jk\7?\2"+
		"\2k\16\3\2\2\2lm\7~\2\2m\20\3\2\2\2no\7>\2\2o\22\3\2\2\2pq\7@\2\2q\24"+
		"\3\2\2\2rs\7>\2\2st\7?\2\2t\26\3\2\2\2uv\7@\2\2vw\7?\2\2w\30\3\2\2\2x"+
		"y\7*\2\2y\32\3\2\2\2z{\7+\2\2{\34\3\2\2\2|}\7}\2\2}\36\3\2\2\2~\177\7"+
		"\177\2\2\177 \3\2\2\2\u0080\u0081\7]\2\2\u0081\"\3\2\2\2\u0082\u0083\7"+
		"_\2\2\u0083$\3\2\2\2\u0084\u0085\7k\2\2\u0085\u0086\7h\2\2\u0086&\3\2"+
		"\2\2\u0087\u0088\7f\2\2\u0088\u0089\7q\2\2\u0089(\3\2\2\2\u008a\u008b"+
		"\7g\2\2\u008b\u008c\7p\2\2\u008c\u008d\7f\2\2\u008d*\3\2\2\2\u008e\u008f"+
		"\7h\2\2\u008f\u0090\7w\2\2\u0090\u0091\7p\2\2\u0091\u0092\7e\2\2\u0092"+
		",\3\2\2\2\u0093\u0094\7t\2\2\u0094\u0095\7g\2\2\u0095\u0096\7v\2\2\u0096"+
		"\u0097\7w\2\2\u0097\u0098\7t\2\2\u0098\u0099\7p\2\2\u0099.\3\2\2\2\u009a"+
		"\u009b\7h\2\2\u009b\u009c\7q\2\2\u009c\u009d\7t\2\2\u009d\60\3\2\2\2\u009e"+
		"\u009f\7c\2\2\u009f\u00a0\7f\2\2\u00a0\u00a1\7f\2\2\u00a1\62\3\2\2\2\u00a2"+
		"\u00a3\7u\2\2\u00a3\u00a4\7w\2\2\u00a4\u00a5\7d\2\2\u00a5\64\3\2\2\2\u00a6"+
		"\u00a7\7f\2\2\u00a7\u00a8\7k\2\2\u00a8\u00a9\7x\2\2\u00a9\66\3\2\2\2\u00aa"+
		"\u00ab\7o\2\2\u00ab\u00ac\7w\2\2\u00ac\u00ad\7n\2\2\u00ad\u00ae\7v\2\2"+
		"\u00ae8\3\2\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1\7t\2\2\u00b1\u00b2\7w\2"+
		"\2\u00b2\u00b3\7g\2\2\u00b3:\3\2\2\2\u00b4\u00b5\7h\2\2\u00b5\u00b6\7"+
		"c\2\2\u00b6\u00b7\7n\2\2\u00b7\u00b8\7u\2\2\u00b8\u00b9\7g\2\2\u00b9<"+
		"\3\2\2\2\u00ba\u00bb\7q\2\2\u00bb\u00bc\7t\2\2\u00bc>\3\2\2\2\u00bd\u00be"+
		"\7c\2\2\u00be\u00bf\7p\2\2\u00bf\u00c0\7f\2\2\u00c0@\3\2\2\2\u00c1\u00c3"+
		"\t\4\2\2\u00c2\u00c1\3\2\2\2\u00c3\u00c4\3\2\2\2\u00c4\u00c2\3\2\2\2\u00c4"+
		"\u00c5\3\2\2\2\u00c5B\3\2\2\2\u00c6\u00ce\7$\2\2\u00c7\u00c9\5E#\2\u00c8"+
		"\u00c7\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\u00c8\3\2\2\2\u00ca\u00cb\3\2"+
		"\2\2\u00cb\u00cd\3\2\2\2\u00cc\u00c8\3\2\2\2\u00cd\u00d0\3\2\2\2\u00ce"+
		"\u00cc\3\2\2\2\u00ce\u00cf\3\2\2\2\u00cf\u00d1\3\2\2\2\u00d0\u00ce\3\2"+
		"\2\2\u00d1\u00d2\7$\2\2\u00d2D\3\2\2\2\u00d3\u00d6\5I%\2\u00d4\u00d6\5"+
		"\7\4\2\u00d5\u00d3\3\2\2\2\u00d5\u00d4\3\2\2\2\u00d6F\3\2\2\2\u00d7\u00d9"+
		"\t\5\2\2\u00d8\u00d7\3\2\2\2\u00d9\u00da\3\2\2\2\u00da\u00d8\3\2\2\2\u00da"+
		"\u00db\3\2\2\2\u00dbH\3\2\2\2\u00dc\u00df\5K&\2\u00dd\u00df\n\6\2\2\u00de"+
		"\u00dc\3\2\2\2\u00de\u00dd\3\2\2\2\u00dfJ\3\2\2\2\u00e0\u00e3\7^\2\2\u00e1"+
		"\u00e4\t\7\2\2\u00e2\u00e4\5G$\2\u00e3\u00e1\3\2\2\2\u00e3\u00e2\3\2\2"+
		"\2\u00e4L\3\2\2\2\16\2S\\^c\u00c4\u00ca\u00ce\u00d5\u00da\u00de\u00e3"+
		"\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}