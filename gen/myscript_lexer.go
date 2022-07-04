// Code generated from /Users/bytedance/GolandProjects/myscript/MyScriptLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

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
	modeNames []string
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
    "'/'", "'+'", "'-'", "'%'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", 
    "'||'", "'&&'", "'!'", "'null'",
  }
  staticData.symbolicNames = []string{
    "", "ELSE", "IF", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "ASSIGN", 
    "COMMA", "MUL", "DIV", "ADD", "SUB", "MOD", "EQUALS", "NOT_EQUALS", 
    "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "LOGICAL_OR", 
    "LOGICAL_AND", "EXCLAMATION", "NULL", "BOOL", "NUMBER", "IDENTIFIER", 
    "WHITESPACE",
  }
  staticData.ruleNames = []string{
    "ELSE", "IF", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "ASSIGN", 
    "COMMA", "MUL", "DIV", "ADD", "SUB", "MOD", "EQUALS", "NOT_EQUALS", 
    "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "LOGICAL_OR", 
    "LOGICAL_AND", "EXCLAMATION", "NULL", "BOOL", "NUMBER", "IDENTIFIER", 
    "WHITESPACE",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 27, 148, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 
	2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 
	8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 
	1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 
	17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 
	1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 
	23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 124, 8, 23, 1, 24, 4, 24, 127, 8, 
	24, 11, 24, 12, 24, 128, 1, 25, 4, 25, 132, 8, 25, 11, 25, 12, 25, 133, 
	1, 25, 5, 25, 137, 8, 25, 10, 25, 12, 25, 140, 9, 25, 1, 26, 4, 26, 143, 
	8, 26, 11, 26, 12, 26, 144, 1, 26, 1, 26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 
	4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 
	14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 
	23, 47, 24, 49, 25, 51, 26, 53, 27, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 
	97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 
	32, 152, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 
	0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 
	1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 
	23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 
	0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 
	0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 
	0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 
	0, 0, 0, 1, 55, 1, 0, 0, 0, 3, 60, 1, 0, 0, 0, 5, 63, 1, 0, 0, 0, 7, 65, 
	1, 0, 0, 0, 9, 67, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13, 71, 1, 0, 0, 0, 
	15, 73, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 77, 1, 0, 0, 0, 21, 79, 1, 
	0, 0, 0, 23, 81, 1, 0, 0, 0, 25, 83, 1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 
	88, 1, 0, 0, 0, 31, 91, 1, 0, 0, 0, 33, 93, 1, 0, 0, 0, 35, 96, 1, 0, 0, 
	0, 37, 98, 1, 0, 0, 0, 39, 101, 1, 0, 0, 0, 41, 104, 1, 0, 0, 0, 43, 107, 
	1, 0, 0, 0, 45, 109, 1, 0, 0, 0, 47, 123, 1, 0, 0, 0, 49, 126, 1, 0, 0, 
	0, 51, 131, 1, 0, 0, 0, 53, 142, 1, 0, 0, 0, 55, 56, 5, 101, 0, 0, 56, 
	57, 5, 108, 0, 0, 57, 58, 5, 115, 0, 0, 58, 59, 5, 101, 0, 0, 59, 2, 1, 
	0, 0, 0, 60, 61, 5, 105, 0, 0, 61, 62, 5, 102, 0, 0, 62, 4, 1, 0, 0, 0, 
	63, 64, 5, 40, 0, 0, 64, 6, 1, 0, 0, 0, 65, 66, 5, 41, 0, 0, 66, 8, 1, 
	0, 0, 0, 67, 68, 5, 123, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70, 5, 125, 0, 0, 
	70, 12, 1, 0, 0, 0, 71, 72, 5, 61, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 
	44, 0, 0, 74, 16, 1, 0, 0, 0, 75, 76, 5, 42, 0, 0, 76, 18, 1, 0, 0, 0, 
	77, 78, 5, 47, 0, 0, 78, 20, 1, 0, 0, 0, 79, 80, 5, 43, 0, 0, 80, 22, 1, 
	0, 0, 0, 81, 82, 5, 45, 0, 0, 82, 24, 1, 0, 0, 0, 83, 84, 5, 37, 0, 0, 
	84, 26, 1, 0, 0, 0, 85, 86, 5, 61, 0, 0, 86, 87, 5, 61, 0, 0, 87, 28, 1, 
	0, 0, 0, 88, 89, 5, 33, 0, 0, 89, 90, 5, 61, 0, 0, 90, 30, 1, 0, 0, 0, 
	91, 92, 5, 60, 0, 0, 92, 32, 1, 0, 0, 0, 93, 94, 5, 60, 0, 0, 94, 95, 5, 
	61, 0, 0, 95, 34, 1, 0, 0, 0, 96, 97, 5, 62, 0, 0, 97, 36, 1, 0, 0, 0, 
	98, 99, 5, 62, 0, 0, 99, 100, 5, 61, 0, 0, 100, 38, 1, 0, 0, 0, 101, 102, 
	5, 124, 0, 0, 102, 103, 5, 124, 0, 0, 103, 40, 1, 0, 0, 0, 104, 105, 5, 
	38, 0, 0, 105, 106, 5, 38, 0, 0, 106, 42, 1, 0, 0, 0, 107, 108, 5, 33, 
	0, 0, 108, 44, 1, 0, 0, 0, 109, 110, 5, 110, 0, 0, 110, 111, 5, 117, 0, 
	0, 111, 112, 5, 108, 0, 0, 112, 113, 5, 108, 0, 0, 113, 46, 1, 0, 0, 0, 
	114, 115, 5, 116, 0, 0, 115, 116, 5, 114, 0, 0, 116, 117, 5, 117, 0, 0, 
	117, 124, 5, 101, 0, 0, 118, 119, 5, 102, 0, 0, 119, 120, 5, 97, 0, 0, 
	120, 121, 5, 108, 0, 0, 121, 122, 5, 115, 0, 0, 122, 124, 5, 101, 0, 0, 
	123, 114, 1, 0, 0, 0, 123, 118, 1, 0, 0, 0, 124, 48, 1, 0, 0, 0, 125, 127, 
	7, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0, 
	0, 0, 128, 129, 1, 0, 0, 0, 129, 50, 1, 0, 0, 0, 130, 132, 7, 1, 0, 0, 
	131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 
	134, 1, 0, 0, 0, 134, 138, 1, 0, 0, 0, 135, 137, 7, 2, 0, 0, 136, 135, 
	1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 
	0, 0, 139, 52, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 143, 7, 3, 0, 0, 
	142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 
	145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 6, 26, 0, 0, 147, 54, 
	1, 0, 0, 0, 6, 0, 123, 128, 133, 138, 144, 1, 6, 0, 0,
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
	MyScriptLexerELSE = 1
	MyScriptLexerIF = 2
	MyScriptLexerL_PAREN = 3
	MyScriptLexerR_PAREN = 4
	MyScriptLexerL_CURLY = 5
	MyScriptLexerR_CURLY = 6
	MyScriptLexerASSIGN = 7
	MyScriptLexerCOMMA = 8
	MyScriptLexerMUL = 9
	MyScriptLexerDIV = 10
	MyScriptLexerADD = 11
	MyScriptLexerSUB = 12
	MyScriptLexerMOD = 13
	MyScriptLexerEQUALS = 14
	MyScriptLexerNOT_EQUALS = 15
	MyScriptLexerLESS = 16
	MyScriptLexerLESS_OR_EQUALS = 17
	MyScriptLexerGREATER = 18
	MyScriptLexerGREATER_OR_EQUALS = 19
	MyScriptLexerLOGICAL_OR = 20
	MyScriptLexerLOGICAL_AND = 21
	MyScriptLexerEXCLAMATION = 22
	MyScriptLexerNULL = 23
	MyScriptLexerBOOL = 24
	MyScriptLexerNUMBER = 25
	MyScriptLexerIDENTIFIER = 26
	MyScriptLexerWHITESPACE = 27
)

