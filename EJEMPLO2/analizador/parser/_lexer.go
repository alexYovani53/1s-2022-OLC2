// Code generated from Lexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 37, 230,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 29, 6, 29, 170, 10, 29, 13, 29, 14, 29, 171, 3,
	30, 6, 30, 175, 10, 30, 13, 30, 14, 30, 176, 3, 30, 3, 30, 6, 30, 181,
	10, 30, 13, 30, 14, 30, 182, 3, 31, 3, 31, 7, 31, 187, 10, 31, 12, 31,
	14, 31, 190, 11, 31, 3, 31, 3, 31, 3, 32, 3, 32, 7, 32, 196, 10, 32, 12,
	32, 14, 32, 199, 11, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 6, 35, 215, 10, 35, 13,
	35, 14, 35, 216, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 39, 3, 39, 2, 2, 40, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 2, 73, 37, 75, 2, 77, 2, 3, 2, 10, 3, 2, 50, 59, 3, 2, 48,
	48, 3, 2, 36, 36, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 9, 2, 34, 35, 37, 37, 45,
	45, 47, 48, 60, 60, 66, 66, 93, 95, 4, 2, 67, 67, 99, 99, 2, 232, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2,
	2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3,
	2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57,
	3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2,
	65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	3, 79, 3, 2, 2, 2, 5, 81, 3, 2, 2, 2, 7, 83, 3, 2, 2, 2, 9, 85, 3, 2, 2,
	2, 11, 87, 3, 2, 2, 2, 13, 94, 3, 2, 2, 2, 15, 98, 3, 2, 2, 2, 17, 106,
	3, 2, 2, 2, 19, 110, 3, 2, 2, 2, 21, 116, 3, 2, 2, 2, 23, 123, 3, 2, 2,
	2, 25, 131, 3, 2, 2, 2, 27, 133, 3, 2, 2, 2, 29, 135, 3, 2, 2, 2, 31, 137,
	3, 2, 2, 2, 33, 140, 3, 2, 2, 2, 35, 143, 3, 2, 2, 2, 37, 145, 3, 2, 2,
	2, 39, 147, 3, 2, 2, 2, 41, 150, 3, 2, 2, 2, 43, 153, 3, 2, 2, 2, 45, 156,
	3, 2, 2, 2, 47, 158, 3, 2, 2, 2, 49, 160, 3, 2, 2, 2, 51, 162, 3, 2, 2,
	2, 53, 164, 3, 2, 2, 2, 55, 166, 3, 2, 2, 2, 57, 169, 3, 2, 2, 2, 59, 174,
	3, 2, 2, 2, 61, 184, 3, 2, 2, 2, 63, 193, 3, 2, 2, 2, 65, 200, 3, 2, 2,
	2, 67, 205, 3, 2, 2, 2, 69, 214, 3, 2, 2, 2, 71, 220, 3, 2, 2, 2, 73, 223,
	3, 2, 2, 2, 75, 226, 3, 2, 2, 2, 77, 228, 3, 2, 2, 2, 79, 80, 7, 42, 2,
	2, 80, 4, 3, 2, 2, 2, 81, 82, 7, 43, 2, 2, 82, 6, 3, 2, 2, 2, 83, 84, 7,
	125, 2, 2, 84, 8, 3, 2, 2, 2, 85, 86, 7, 127, 2, 2, 86, 10, 3, 2, 2, 2,
	87, 88, 7, 117, 2, 2, 88, 89, 7, 123, 2, 2, 89, 90, 7, 117, 2, 2, 90, 91,
	7, 118, 2, 2, 91, 92, 7, 103, 2, 2, 92, 93, 7, 111, 2, 2, 93, 12, 3, 2,
	2, 2, 94, 95, 7, 113, 2, 2, 95, 96, 7, 119, 2, 2, 96, 97, 7, 118, 2, 2,
	97, 14, 3, 2, 2, 2, 98, 99, 7, 114, 2, 2, 99, 100, 7, 116, 2, 2, 100, 101,
	7, 107, 2, 2, 101, 102, 7, 112, 2, 2, 102, 103, 7, 118, 2, 2, 103, 104,
	7, 110, 2, 2, 104, 105, 7, 112, 2, 2, 105, 16, 3, 2, 2, 2, 106, 107, 7,
	107, 2, 2, 107, 108, 7, 112, 2, 2, 108, 109, 7, 118, 2, 2, 109, 18, 3,
	2, 2, 2, 110, 111, 7, 104, 2, 2, 111, 112, 7, 110, 2, 2, 112, 113, 7, 113,
	2, 2, 113, 114, 7, 99, 2, 2, 114, 115, 7, 118, 2, 2, 115, 20, 3, 2, 2,
	2, 116, 117, 7, 117, 2, 2, 117, 118, 7, 118, 2, 2, 118, 119, 7, 116, 2,
	2, 119, 120, 7, 107, 2, 2, 120, 121, 7, 112, 2, 2, 121, 122, 7, 105, 2,
	2, 122, 22, 3, 2, 2, 2, 123, 124, 7, 100, 2, 2, 124, 125, 7, 113, 2, 2,
	125, 126, 7, 113, 2, 2, 126, 127, 7, 110, 2, 2, 127, 128, 7, 103, 2, 2,
	128, 129, 7, 99, 2, 2, 129, 130, 7, 112, 2, 2, 130, 24, 3, 2, 2, 2, 131,
	132, 7, 48, 2, 2, 132, 26, 3, 2, 2, 2, 133, 134, 7, 46, 2, 2, 134, 28,
	3, 2, 2, 2, 135, 136, 7, 61, 2, 2, 136, 30, 3, 2, 2, 2, 137, 138, 7, 40,
	2, 2, 138, 139, 7, 40, 2, 2, 139, 32, 3, 2, 2, 2, 140, 141, 7, 126, 2,
	2, 141, 142, 7, 126, 2, 2, 142, 34, 3, 2, 2, 2, 143, 144, 7, 35, 2, 2,
	144, 36, 3, 2, 2, 2, 145, 146, 7, 63, 2, 2, 146, 38, 3, 2, 2, 2, 147, 148,
	7, 35, 2, 2, 148, 149, 7, 63, 2, 2, 149, 40, 3, 2, 2, 2, 150, 151, 7, 64,
	2, 2, 151, 152, 7, 63, 2, 2, 152, 42, 3, 2, 2, 2, 153, 154, 7, 62, 2, 2,
	154, 155, 7, 63, 2, 2, 155, 44, 3, 2, 2, 2, 156, 157, 7, 64, 2, 2, 157,
	46, 3, 2, 2, 2, 158, 159, 7, 62, 2, 2, 159, 48, 3, 2, 2, 2, 160, 161, 7,
	44, 2, 2, 161, 50, 3, 2, 2, 2, 162, 163, 7, 49, 2, 2, 163, 52, 3, 2, 2,
	2, 164, 165, 7, 45, 2, 2, 165, 54, 3, 2, 2, 2, 166, 167, 7, 47, 2, 2, 167,
	56, 3, 2, 2, 2, 168, 170, 9, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 171, 3,
	2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 58, 3, 2, 2,
	2, 173, 175, 9, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180,
	9, 3, 2, 2, 179, 181, 9, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 182, 3, 2,
	2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 60, 3, 2, 2, 2,
	184, 188, 7, 36, 2, 2, 185, 187, 10, 4, 2, 2, 186, 185, 3, 2, 2, 2, 187,
	190, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 191,
	3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 191, 192, 7, 36, 2, 2, 192, 62, 3, 2,
	2, 2, 193, 197, 9, 5, 2, 2, 194, 196, 9, 6, 2, 2, 195, 194, 3, 2, 2, 2,
	196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198,
	64, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 201, 7, 118, 2, 2, 201, 202,
	7, 116, 2, 2, 202, 203, 7, 119, 2, 2, 203, 204, 7, 103, 2, 2, 204, 66,
	3, 2, 2, 2, 205, 206, 7, 104, 2, 2, 206, 207, 7, 99, 2, 2, 207, 208, 7,
	110, 2, 2, 208, 209, 7, 117, 2, 2, 209, 210, 7, 103, 2, 2, 210, 211, 3,
	2, 2, 2, 211, 212, 5, 71, 36, 2, 212, 68, 3, 2, 2, 2, 213, 215, 9, 7, 2,
	2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216,
	217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 8, 35, 2, 2, 219, 70,
	3, 2, 2, 2, 220, 221, 7, 94, 2, 2, 221, 222, 9, 8, 2, 2, 222, 72, 3, 2,
	2, 2, 223, 224, 5, 75, 38, 2, 224, 225, 5, 77, 39, 2, 225, 74, 3, 2, 2,
	2, 226, 227, 9, 9, 2, 2, 227, 76, 3, 2, 2, 2, 228, 229, 9, 9, 2, 2, 229,
	78, 3, 2, 2, 2, 9, 2, 171, 176, 182, 188, 197, 216, 3, 8, 2, 2,
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
	"", "'('", "')'", "'{'", "'}'", "'system'", "'out'", "'println'", "'int'",
	"'float'", "'string'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'",
	"'!'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'",
	"'-'", "", "", "", "", "'true'",
}

var lexerSymbolicNames = []string{
	"", "LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE",
	"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND",
	"OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR",
	"MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "ID",
	"TRUE", "FALSE", "WHITESPACE", "AB",
}

var lexerRuleNames = []string{
	"LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE",
	"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND",
	"OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR",
	"MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "ID",
	"TRUE", "FALSE", "WHITESPACE", "ESC_SEQ", "AB", "A", "B",
}

type Lexer struct {
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

func NewLexer(input antlr.CharStream) *Lexer {

	l := new(Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lexer tokens.
const (
	LexerLP         = 1
	LexerRP         = 2
	LexerL_LLAVE    = 3
	LexerR_LLAVE    = 4
	LexerSYSTEM     = 5
	LexerOUT        = 6
	LexerPRINTLN    = 7
	LexerINTTYPE    = 8
	LexerFLOATTYPE  = 9
	LexerSTRINGTYPE = 10
	LexerBOOLTYPE   = 11
	LexerPUNTO      = 12
	LexerCOMA       = 13
	LexerPTCOMA     = 14
	LexerAND        = 15
	LexerOR         = 16
	LexerNOT        = 17
	LexerIGUAL      = 18
	LexerDIFERENTE  = 19
	LexerMAYORIGUAL = 20
	LexerMENORIGUAL = 21
	LexerMAYOR      = 22
	LexerMENOR      = 23
	LexerMUL        = 24
	LexerDIV        = 25
	LexerADD        = 26
	LexerSUB        = 27
	LexerNUMBER     = 28
	LexerFLOAT      = 29
	LexerSTRING     = 30
	LexerID         = 31
	LexerTRUE       = 32
	LexerFALSE      = 33
	LexerWHITESPACE = 34
	LexerAB         = 35
)
