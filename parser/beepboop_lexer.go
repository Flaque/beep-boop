// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 123,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 7, 3, 50, 10, 3, 12, 3, 14, 3, 53, 11, 3, 3, 3, 3, 3,
	3, 4, 6, 4, 58, 10, 4, 13, 4, 14, 4, 59, 3, 5, 6, 5, 63, 10, 5, 13, 5,
	14, 5, 64, 3, 5, 3, 5, 3, 6, 6, 6, 70, 10, 6, 13, 6, 14, 6, 71, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 6, 15, 91, 10, 15, 13, 15, 14, 15, 92,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 2, 2, 23, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 3, 2, 6, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 3, 2, 50,
	59, 4, 2, 67, 92, 99, 124, 2, 127, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3,
	2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37,
	3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 3,
	45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 57, 3, 2, 2, 2, 9, 62, 3, 2, 2, 2,
	11, 69, 3, 2, 2, 2, 13, 73, 3, 2, 2, 2, 15, 75, 3, 2, 2, 2, 17, 77, 3,
	2, 2, 2, 19, 79, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2, 23, 83, 3, 2, 2, 2, 25,
	85, 3, 2, 2, 2, 27, 87, 3, 2, 2, 2, 29, 90, 3, 2, 2, 2, 31, 94, 3, 2, 2,
	2, 33, 97, 3, 2, 2, 2, 35, 100, 3, 2, 2, 2, 37, 104, 3, 2, 2, 2, 39, 109,
	3, 2, 2, 2, 41, 116, 3, 2, 2, 2, 43, 120, 3, 2, 2, 2, 45, 46, 7, 38, 2,
	2, 46, 4, 3, 2, 2, 2, 47, 51, 7, 37, 2, 2, 48, 50, 10, 2, 2, 2, 49, 48,
	3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 8, 3, 2, 2, 55, 6, 3, 2,
	2, 2, 56, 58, 9, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 8, 3, 2, 2, 2, 61, 63, 9, 3, 2, 2,
	62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3,
	2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 8, 5, 3, 2, 67, 10, 3, 2, 2, 2, 68,
	70, 9, 4, 2, 2, 69, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2,
	2, 71, 72, 3, 2, 2, 2, 72, 12, 3, 2, 2, 2, 73, 74, 7, 45, 2, 2, 74, 14,
	3, 2, 2, 2, 75, 76, 7, 47, 2, 2, 76, 16, 3, 2, 2, 2, 77, 78, 7, 44, 2,
	2, 78, 18, 3, 2, 2, 2, 79, 80, 7, 49, 2, 2, 80, 20, 3, 2, 2, 2, 81, 82,
	7, 63, 2, 2, 82, 22, 3, 2, 2, 2, 83, 84, 7, 126, 2, 2, 84, 24, 3, 2, 2,
	2, 85, 86, 7, 42, 2, 2, 86, 26, 3, 2, 2, 2, 87, 88, 7, 43, 2, 2, 88, 28,
	3, 2, 2, 2, 89, 91, 9, 5, 2, 2, 90, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2,
	92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 30, 3, 2, 2, 2, 94, 95, 7,
	107, 2, 2, 95, 96, 7, 104, 2, 2, 96, 32, 3, 2, 2, 2, 97, 98, 7, 102, 2,
	2, 98, 99, 7, 113, 2, 2, 99, 34, 3, 2, 2, 2, 100, 101, 7, 103, 2, 2, 101,
	102, 7, 112, 2, 2, 102, 103, 7, 102, 2, 2, 103, 36, 3, 2, 2, 2, 104, 105,
	7, 104, 2, 2, 105, 106, 7, 119, 2, 2, 106, 107, 7, 112, 2, 2, 107, 108,
	7, 101, 2, 2, 108, 38, 3, 2, 2, 2, 109, 110, 7, 116, 2, 2, 110, 111, 7,
	103, 2, 2, 111, 112, 7, 118, 2, 2, 112, 113, 7, 119, 2, 2, 113, 114, 7,
	116, 2, 2, 114, 115, 7, 112, 2, 2, 115, 40, 3, 2, 2, 2, 116, 117, 7, 104,
	2, 2, 117, 118, 7, 113, 2, 2, 118, 119, 7, 116, 2, 2, 119, 42, 3, 2, 2,
	2, 120, 121, 7, 107, 2, 2, 121, 122, 7, 117, 2, 2, 122, 44, 3, 2, 2, 2,
	8, 2, 51, 59, 64, 71, 92, 4, 8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'$'", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'='", "'|'", "'('",
	"')'", "", "'if'", "'do'", "'end'", "'func'", "'return'", "'for'", "'is'",
}

var lexerSymbolicNames = []string{
	"", "", "COMMENT", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "MULT",
	"DIVIDE", "ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", "IF", "DO", "END",
	"FUNC", "RETURN", "FOR", "IS",
}

var lexerRuleNames = []string{
	"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "MULT",
	"DIVIDE", "ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", "IF", "DO", "END",
	"FUNC", "RETURN", "FOR", "IS",
}

type BeepBoopLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewBeepBoopLexer(input antlr.CharStream) *BeepBoopLexer {

	l := new(BeepBoopLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "BeepBoop.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BeepBoopLexer tokens.
const (
	BeepBoopLexerT__0       = 1
	BeepBoopLexerCOMMENT    = 2
	BeepBoopLexerNEWLINE    = 3
	BeepBoopLexerWHITESPACE = 4
	BeepBoopLexerINT        = 5
	BeepBoopLexerPLUS       = 6
	BeepBoopLexerMINUS      = 7
	BeepBoopLexerMULT       = 8
	BeepBoopLexerDIVIDE     = 9
	BeepBoopLexerASSIGN     = 10
	BeepBoopLexerPIPE       = 11
	BeepBoopLexerLPAREN     = 12
	BeepBoopLexerRPAREN     = 13
	BeepBoopLexerSTRING     = 14
	BeepBoopLexerIF         = 15
	BeepBoopLexerDO         = 16
	BeepBoopLexerEND        = 17
	BeepBoopLexerFUNC       = 18
	BeepBoopLexerRETURN     = 19
	BeepBoopLexerFOR        = 20
	BeepBoopLexerIS         = 21
)
