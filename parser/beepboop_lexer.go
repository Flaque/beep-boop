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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 135,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 54, 10, 3, 12, 3, 14,
	3, 57, 11, 3, 3, 3, 3, 3, 3, 4, 6, 4, 62, 10, 4, 13, 4, 14, 4, 63, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 71, 10, 6, 13, 6, 14, 6, 72, 3, 6, 3, 6,
	3, 7, 6, 7, 78, 10, 7, 13, 7, 14, 7, 79, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 6, 16, 99, 10, 16, 13, 16, 14, 16, 100, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 2, 2, 25,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 3, 2, 6, 4, 2, 12, 12, 15, 15, 3, 2, 11, 11,
	3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 2, 139, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 51, 3,
	2, 2, 2, 7, 61, 3, 2, 2, 2, 9, 65, 3, 2, 2, 2, 11, 70, 3, 2, 2, 2, 13,
	77, 3, 2, 2, 2, 15, 81, 3, 2, 2, 2, 17, 83, 3, 2, 2, 2, 19, 85, 3, 2, 2,
	2, 21, 87, 3, 2, 2, 2, 23, 89, 3, 2, 2, 2, 25, 91, 3, 2, 2, 2, 27, 93,
	3, 2, 2, 2, 29, 95, 3, 2, 2, 2, 31, 98, 3, 2, 2, 2, 33, 102, 3, 2, 2, 2,
	35, 105, 3, 2, 2, 2, 37, 108, 3, 2, 2, 2, 39, 112, 3, 2, 2, 2, 41, 117,
	3, 2, 2, 2, 43, 124, 3, 2, 2, 2, 45, 128, 3, 2, 2, 2, 47, 131, 3, 2, 2,
	2, 49, 50, 7, 38, 2, 2, 50, 4, 3, 2, 2, 2, 51, 55, 7, 37, 2, 2, 52, 54,
	10, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 8,
	3, 2, 2, 59, 6, 3, 2, 2, 2, 60, 62, 9, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62,
	63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 8, 3, 2, 2,
	2, 65, 66, 7, 34, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 8, 5, 3, 2, 68, 10,
	3, 2, 2, 2, 69, 71, 9, 3, 2, 2, 70, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 8,
	6, 3, 2, 75, 12, 3, 2, 2, 2, 76, 78, 9, 4, 2, 2, 77, 76, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 14, 3, 2, 2,
	2, 81, 82, 7, 45, 2, 2, 82, 16, 3, 2, 2, 2, 83, 84, 7, 47, 2, 2, 84, 18,
	3, 2, 2, 2, 85, 86, 7, 44, 2, 2, 86, 20, 3, 2, 2, 2, 87, 88, 7, 49, 2,
	2, 88, 22, 3, 2, 2, 2, 89, 90, 7, 63, 2, 2, 90, 24, 3, 2, 2, 2, 91, 92,
	7, 126, 2, 2, 92, 26, 3, 2, 2, 2, 93, 94, 7, 42, 2, 2, 94, 28, 3, 2, 2,
	2, 95, 96, 7, 43, 2, 2, 96, 30, 3, 2, 2, 2, 97, 99, 9, 5, 2, 2, 98, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 32, 3, 2, 2, 2, 102, 103, 7, 107, 2, 2, 103, 104, 7, 104, 2, 2,
	104, 34, 3, 2, 2, 2, 105, 106, 7, 102, 2, 2, 106, 107, 7, 113, 2, 2, 107,
	36, 3, 2, 2, 2, 108, 109, 7, 103, 2, 2, 109, 110, 7, 112, 2, 2, 110, 111,
	7, 102, 2, 2, 111, 38, 3, 2, 2, 2, 112, 113, 7, 104, 2, 2, 113, 114, 7,
	119, 2, 2, 114, 115, 7, 112, 2, 2, 115, 116, 7, 101, 2, 2, 116, 40, 3,
	2, 2, 2, 117, 118, 7, 116, 2, 2, 118, 119, 7, 103, 2, 2, 119, 120, 7, 118,
	2, 2, 120, 121, 7, 119, 2, 2, 121, 122, 7, 116, 2, 2, 122, 123, 7, 112,
	2, 2, 123, 42, 3, 2, 2, 2, 124, 125, 7, 104, 2, 2, 125, 126, 7, 113, 2,
	2, 126, 127, 7, 116, 2, 2, 127, 44, 3, 2, 2, 2, 128, 129, 7, 107, 2, 2,
	129, 130, 7, 117, 2, 2, 130, 46, 3, 2, 2, 2, 131, 132, 11, 2, 2, 2, 132,
	133, 3, 2, 2, 2, 133, 134, 8, 24, 2, 2, 134, 48, 3, 2, 2, 2, 8, 2, 55,
	63, 72, 79, 100, 4, 8, 2, 2, 2, 3, 2,
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
	"", "'$'", "", "", "' '", "", "", "'+'", "'-'", "'*'", "'/'", "'='", "'|'",
	"'('", "')'", "", "'if'", "'do'", "'end'", "'func'", "'return'", "'for'",
	"'is'",
}

var lexerSymbolicNames = []string{
	"", "", "COMMENT", "NEWLINE", "WHITESPACE", "TABSPACE", "INT", "PLUS",
	"MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING",
	"IF", "DO", "END", "FUNC", "RETURN", "FOR", "IS", "UNKNOWN",
}

var lexerRuleNames = []string{
	"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "TABSPACE", "INT", "PLUS",
	"MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING",
	"IF", "DO", "END", "FUNC", "RETURN", "FOR", "IS", "UNKNOWN",
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
	BeepBoopLexerTABSPACE   = 5
	BeepBoopLexerINT        = 6
	BeepBoopLexerPLUS       = 7
	BeepBoopLexerMINUS      = 8
	BeepBoopLexerMULT       = 9
	BeepBoopLexerDIVIDE     = 10
	BeepBoopLexerASSIGN     = 11
	BeepBoopLexerPIPE       = 12
	BeepBoopLexerLPAREN     = 13
	BeepBoopLexerRPAREN     = 14
	BeepBoopLexerSTRING     = 15
	BeepBoopLexerIF         = 16
	BeepBoopLexerDO         = 17
	BeepBoopLexerEND        = 18
	BeepBoopLexerFUNC       = 19
	BeepBoopLexerRETURN     = 20
	BeepBoopLexerFOR        = 21
	BeepBoopLexerIS         = 22
	BeepBoopLexerUNKNOWN    = 23
)
