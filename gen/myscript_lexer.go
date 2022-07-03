// Code generated from D:/GoProject/myscript\MyScriptLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MyScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var myscriptlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func myscriptlexerLexerInit() {
	staticData := &myscriptlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'else'", "'if'", "'('", "')'", "'{'", "'}'", "'='", "','", "'*'",
		"'/'", "'+'", "'-'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'||'",
		"'&&'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "ELSE", "IF", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "ASSIGN",
		"COMMA", "MUL", "DIV", "ADD", "SUB", "EQUALS", "NOT_EQUALS", "LESS",
		"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "LOGICAL_OR", "LOGICAL_AND",
		"EXCLAMATION", "NUMBER", "IDENTIFIER", "BOOL", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"ELSE", "IF", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "ASSIGN",
		"COMMA", "MUL", "DIV", "ADD", "SUB", "EQUALS", "NOT_EQUALS", "LESS",
		"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "LOGICAL_OR", "LOGICAL_AND",
		"EXCLAMATION", "NUMBER", "IDENTIFIER", "BOOL", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 137, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 4, 21, 105, 8, 21, 11, 21, 12,
		21, 106, 1, 22, 4, 22, 110, 8, 22, 11, 22, 12, 22, 111, 1, 22, 5, 22, 115,
		8, 22, 10, 22, 12, 22, 118, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 3, 23, 129, 8, 23, 1, 24, 4, 24, 132, 8, 24, 11,
		24, 12, 24, 133, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 2, 0, 48, 57,
		95, 95, 3, 0, 9, 10, 13, 13, 32, 32, 141, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 1, 51, 1, 0, 0, 0, 3, 56, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0,
		0, 0, 9, 63, 1, 0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 69,
		1, 0, 0, 0, 17, 71, 1, 0, 0, 0, 19, 73, 1, 0, 0, 0, 21, 75, 1, 0, 0, 0,
		23, 77, 1, 0, 0, 0, 25, 79, 1, 0, 0, 0, 27, 82, 1, 0, 0, 0, 29, 85, 1,
		0, 0, 0, 31, 87, 1, 0, 0, 0, 33, 90, 1, 0, 0, 0, 35, 92, 1, 0, 0, 0, 37,
		95, 1, 0, 0, 0, 39, 98, 1, 0, 0, 0, 41, 101, 1, 0, 0, 0, 43, 104, 1, 0,
		0, 0, 45, 109, 1, 0, 0, 0, 47, 128, 1, 0, 0, 0, 49, 131, 1, 0, 0, 0, 51,
		52, 5, 101, 0, 0, 52, 53, 5, 108, 0, 0, 53, 54, 5, 115, 0, 0, 54, 55, 5,
		101, 0, 0, 55, 2, 1, 0, 0, 0, 56, 57, 5, 105, 0, 0, 57, 58, 5, 102, 0,
		0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 40, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62, 5,
		41, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 123, 0, 0, 64, 10, 1, 0, 0, 0,
		65, 66, 5, 125, 0, 0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 61, 0, 0, 68, 14,
		1, 0, 0, 0, 69, 70, 5, 44, 0, 0, 70, 16, 1, 0, 0, 0, 71, 72, 5, 42, 0,
		0, 72, 18, 1, 0, 0, 0, 73, 74, 5, 47, 0, 0, 74, 20, 1, 0, 0, 0, 75, 76,
		5, 43, 0, 0, 76, 22, 1, 0, 0, 0, 77, 78, 5, 45, 0, 0, 78, 24, 1, 0, 0,
		0, 79, 80, 5, 61, 0, 0, 80, 81, 5, 61, 0, 0, 81, 26, 1, 0, 0, 0, 82, 83,
		5, 33, 0, 0, 83, 84, 5, 61, 0, 0, 84, 28, 1, 0, 0, 0, 85, 86, 5, 60, 0,
		0, 86, 30, 1, 0, 0, 0, 87, 88, 5, 60, 0, 0, 88, 89, 5, 61, 0, 0, 89, 32,
		1, 0, 0, 0, 90, 91, 5, 62, 0, 0, 91, 34, 1, 0, 0, 0, 92, 93, 5, 62, 0,
		0, 93, 94, 5, 61, 0, 0, 94, 36, 1, 0, 0, 0, 95, 96, 5, 124, 0, 0, 96, 97,
		5, 124, 0, 0, 97, 38, 1, 0, 0, 0, 98, 99, 5, 38, 0, 0, 99, 100, 5, 38,
		0, 0, 100, 40, 1, 0, 0, 0, 101, 102, 5, 33, 0, 0, 102, 42, 1, 0, 0, 0,
		103, 105, 7, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 44, 1, 0, 0, 0, 108, 110, 7,
		1, 0, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 116, 1, 0, 0, 0, 113, 115, 7, 2, 0, 0, 114,
		113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 46, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 5, 116,
		0, 0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 117, 0, 0, 122, 129, 5, 101,
		0, 0, 123, 124, 5, 102, 0, 0, 124, 125, 5, 97, 0, 0, 125, 126, 5, 108,
		0, 0, 126, 127, 5, 115, 0, 0, 127, 129, 5, 101, 0, 0, 128, 119, 1, 0, 0,
		0, 128, 123, 1, 0, 0, 0, 129, 48, 1, 0, 0, 0, 130, 132, 7, 3, 0, 0, 131,
		130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 6, 24, 0, 0, 136, 50, 1, 0,
		0, 0, 6, 0, 106, 111, 116, 128, 133, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MyScriptLexerInit initializes any static state used to implement MyScriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMyScriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyScriptLexerInit() {
	staticData := &myscriptlexerLexerStaticData
	staticData.once.Do(myscriptlexerLexerInit)
}

// NewMyScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMyScriptLexer(input antlr.CharStream) *MyScriptLexer {
	MyScriptLexerInit()
	l := new(MyScriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &myscriptlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "MyScriptLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MyScriptLexer tokens.
const (
	MyScriptLexerELSE              = 1
	MyScriptLexerIF                = 2
	MyScriptLexerL_PAREN           = 3
	MyScriptLexerR_PAREN           = 4
	MyScriptLexerL_CURLY           = 5
	MyScriptLexerR_CURLY           = 6
	MyScriptLexerASSIGN            = 7
	MyScriptLexerCOMMA             = 8
	MyScriptLexerMUL               = 9
	MyScriptLexerDIV               = 10
	MyScriptLexerADD               = 11
	MyScriptLexerSUB               = 12
	MyScriptLexerEQUALS            = 13
	MyScriptLexerNOT_EQUALS        = 14
	MyScriptLexerLESS              = 15
	MyScriptLexerLESS_OR_EQUALS    = 16
	MyScriptLexerGREATER           = 17
	MyScriptLexerGREATER_OR_EQUALS = 18
	MyScriptLexerLOGICAL_OR        = 19
	MyScriptLexerLOGICAL_AND       = 20
	MyScriptLexerEXCLAMATION       = 21
	MyScriptLexerNUMBER            = 22
	MyScriptLexerIDENTIFIER        = 23
	MyScriptLexerBOOL              = 24
	MyScriptLexerWHITESPACE        = 25
)
