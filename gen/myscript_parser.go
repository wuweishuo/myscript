// Code generated from D:/GoProject/myscript\MyScriptParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MyScriptParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MyScriptParser struct {
	*antlr.BaseParser
}

var myscriptparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func myscriptparserParserInit() {
	staticData := &myscriptparserParserStaticData
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
		"statementList", "statement", "assignStmt", "ifStmt", "block", "arithmetic",
		"logicStmt", "relation", "funcStmt", "paramList", "param",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 112, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 4, 0, 24, 8, 0, 11, 0, 12, 0, 25, 1, 1, 1, 1, 1, 1, 3, 1, 31,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 39, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 3, 3, 49, 8, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 62, 8, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 70, 8, 5, 10, 5, 12, 5, 73, 9, 5, 1, 6,
		1, 6, 1, 6, 5, 6, 78, 8, 6, 10, 6, 12, 6, 81, 9, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 95, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 105, 8, 9, 10, 9, 12, 9,
		108, 9, 9, 1, 10, 1, 10, 1, 10, 0, 1, 10, 11, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 0, 5, 1, 0, 9, 10, 1, 0, 11, 12, 1, 0, 19, 20, 1, 0, 13, 18,
		1, 0, 22, 24, 118, 0, 23, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4, 32, 1, 0, 0,
		0, 6, 40, 1, 0, 0, 0, 8, 50, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12, 74, 1,
		0, 0, 0, 14, 94, 1, 0, 0, 0, 16, 96, 1, 0, 0, 0, 18, 101, 1, 0, 0, 0, 20,
		109, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0, 23, 22, 1, 0, 0, 0, 24, 25, 1, 0,
		0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 1, 1, 0, 0, 0, 27, 31,
		3, 4, 2, 0, 28, 31, 3, 6, 3, 0, 29, 31, 3, 16, 8, 0, 30, 27, 1, 0, 0, 0,
		30, 28, 1, 0, 0, 0, 30, 29, 1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 32, 33, 5, 23,
		0, 0, 33, 38, 5, 7, 0, 0, 34, 39, 3, 10, 5, 0, 35, 39, 3, 14, 7, 0, 36,
		39, 5, 22, 0, 0, 37, 39, 5, 24, 0, 0, 38, 34, 1, 0, 0, 0, 38, 35, 1, 0,
		0, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41,
		5, 2, 0, 0, 41, 42, 3, 12, 6, 0, 42, 48, 3, 8, 4, 0, 43, 46, 5, 1, 0, 0,
		44, 47, 3, 6, 3, 0, 45, 47, 3, 8, 4, 0, 46, 44, 1, 0, 0, 0, 46, 45, 1,
		0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 43, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49,
		7, 1, 0, 0, 0, 50, 51, 5, 5, 0, 0, 51, 52, 3, 0, 0, 0, 52, 53, 5, 6, 0,
		0, 53, 9, 1, 0, 0, 0, 54, 55, 6, 5, -1, 0, 55, 56, 5, 3, 0, 0, 56, 57,
		3, 10, 5, 0, 57, 58, 5, 4, 0, 0, 58, 62, 1, 0, 0, 0, 59, 62, 5, 22, 0,
		0, 60, 62, 5, 23, 0, 0, 61, 54, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 60,
		1, 0, 0, 0, 62, 71, 1, 0, 0, 0, 63, 64, 10, 5, 0, 0, 64, 65, 7, 0, 0, 0,
		65, 70, 3, 10, 5, 6, 66, 67, 10, 4, 0, 0, 67, 68, 7, 1, 0, 0, 68, 70, 3,
		10, 5, 5, 69, 63, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71,
		69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 11, 1, 0, 0, 0, 73, 71, 1, 0, 0,
		0, 74, 79, 3, 14, 7, 0, 75, 76, 7, 2, 0, 0, 76, 78, 3, 14, 7, 0, 77, 75,
		1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0,
		80, 13, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 5, 21, 0, 0, 83, 95, 3,
		14, 7, 0, 84, 85, 3, 10, 5, 0, 85, 86, 7, 3, 0, 0, 86, 87, 3, 10, 5, 0,
		87, 95, 1, 0, 0, 0, 88, 89, 5, 3, 0, 0, 89, 90, 3, 14, 7, 0, 90, 91, 5,
		4, 0, 0, 91, 95, 1, 0, 0, 0, 92, 95, 5, 24, 0, 0, 93, 95, 5, 23, 0, 0,
		94, 82, 1, 0, 0, 0, 94, 84, 1, 0, 0, 0, 94, 88, 1, 0, 0, 0, 94, 92, 1,
		0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 15, 1, 0, 0, 0, 96, 97, 5, 23, 0, 0, 97,
		98, 5, 3, 0, 0, 98, 99, 3, 18, 9, 0, 99, 100, 5, 4, 0, 0, 100, 17, 1, 0,
		0, 0, 101, 106, 3, 20, 10, 0, 102, 103, 5, 8, 0, 0, 103, 105, 3, 20, 10,
		0, 104, 102, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 19, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 7,
		4, 0, 0, 110, 21, 1, 0, 0, 0, 11, 25, 30, 38, 46, 48, 61, 69, 71, 79, 94,
		106,
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

// MyScriptParserInit initializes any static state used to implement MyScriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMyScriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyScriptParserInit() {
	staticData := &myscriptparserParserStaticData
	staticData.once.Do(myscriptparserParserInit)
}

// NewMyScriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMyScriptParser(input antlr.TokenStream) *MyScriptParser {
	MyScriptParserInit()
	this := new(MyScriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &myscriptparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MyScriptParser.g4"

	return this
}

// MyScriptParser tokens.
const (
	MyScriptParserEOF               = antlr.TokenEOF
	MyScriptParserELSE              = 1
	MyScriptParserIF                = 2
	MyScriptParserL_PAREN           = 3
	MyScriptParserR_PAREN           = 4
	MyScriptParserL_CURLY           = 5
	MyScriptParserR_CURLY           = 6
	MyScriptParserASSIGN            = 7
	MyScriptParserCOMMA             = 8
	MyScriptParserMUL               = 9
	MyScriptParserDIV               = 10
	MyScriptParserADD               = 11
	MyScriptParserSUB               = 12
	MyScriptParserEQUALS            = 13
	MyScriptParserNOT_EQUALS        = 14
	MyScriptParserLESS              = 15
	MyScriptParserLESS_OR_EQUALS    = 16
	MyScriptParserGREATER           = 17
	MyScriptParserGREATER_OR_EQUALS = 18
	MyScriptParserLOGICAL_OR        = 19
	MyScriptParserLOGICAL_AND       = 20
	MyScriptParserEXCLAMATION       = 21
	MyScriptParserNUMBER            = 22
	MyScriptParserIDENTIFIER        = 23
	MyScriptParserBOOL              = 24
	MyScriptParserWHITESPACE        = 25
)

// MyScriptParser rules.
const (
	MyScriptParserRULE_statementList = 0
	MyScriptParserRULE_statement     = 1
	MyScriptParserRULE_assignStmt    = 2
	MyScriptParserRULE_ifStmt        = 3
	MyScriptParserRULE_block         = 4
	MyScriptParserRULE_arithmetic    = 5
	MyScriptParserRULE_logicStmt     = 6
	MyScriptParserRULE_relation      = 7
	MyScriptParserRULE_funcStmt      = 8
	MyScriptParserRULE_paramList     = 9
	MyScriptParserRULE_param         = 10
)

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *MyScriptParser) StatementList() (localctx IStatementListContext) {
	this := p
	_ = this

	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MyScriptParserRULE_statementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MyScriptParserIF || _la == MyScriptParserIDENTIFIER {
		{
			p.SetState(22)
			p.Statement()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) FuncStmt() IFuncStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *MyScriptParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MyScriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.IfStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.FuncStmt()
		}

	}

	return localctx
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_assignStmt
	return p
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *AssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserASSIGN, 0)
}

func (s *AssignStmtContext) Arithmetic() IArithmeticContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *AssignStmtContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *AssignStmtContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserNUMBER, 0)
}

func (s *AssignStmtContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MyScriptParserBOOL, 0)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (p *MyScriptParser) AssignStmt() (localctx IAssignStmtContext) {
	this := p
	_ = this

	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyScriptParserRULE_assignStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(MyScriptParserIDENTIFIER)
	}
	{
		p.SetState(33)
		p.Match(MyScriptParserASSIGN)
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(34)
			p.arithmetic(0)
		}

	case 2:
		{
			p.SetState(35)
			p.Relation()
		}

	case 3:
		{
			p.SetState(36)
			p.Match(MyScriptParserNUMBER)
		}

	case 4:
		{
			p.SetState(37)
			p.Match(MyScriptParserBOOL)
		}

	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIF, 0)
}

func (s *IfStmtContext) LogicStmt() ILogicStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicStmtContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MyScriptParserELSE, 0)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *MyScriptParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyScriptParserRULE_ifStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(MyScriptParserIF)
	}
	{
		p.SetState(41)
		p.LogicStmt()
	}
	{
		p.SetState(42)
		p.Block()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MyScriptParserELSE {
		{
			p.SetState(43)
			p.Match(MyScriptParserELSE)
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MyScriptParserIF:
			{
				p.SetState(44)
				p.IfStmt()
			}

		case MyScriptParserL_CURLY:
			{
				p.SetState(45)
				p.Block()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MyScriptParserL_CURLY, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MyScriptParserR_CURLY, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *MyScriptParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(MyScriptParserL_CURLY)
	}
	{
		p.SetState(51)
		p.StatementList()
	}
	{
		p.SetState(52)
		p.Match(MyScriptParserR_CURLY)
	}

	return localctx
}

// IArithmeticContext is an interface to support dynamic dispatch.
type IArithmeticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsArithmeticContext differentiates from other interfaces.
	IsArithmeticContext()
}

type ArithmeticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyArithmeticContext() *ArithmeticContext {
	var p = new(ArithmeticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_arithmetic
	return p
}

func (*ArithmeticContext) IsArithmeticContext() {}

func NewArithmeticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticContext {
	var p = new(ArithmeticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_arithmetic

	return p
}

func (s *ArithmeticContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserL_PAREN, 0)
}

func (s *ArithmeticContext) AllArithmetic() []IArithmeticContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmeticContext); ok {
			len++
		}
	}

	tst := make([]IArithmeticContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmeticContext); ok {
			tst[i] = t.(IArithmeticContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticContext) Arithmetic(i int) IArithmeticContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *ArithmeticContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserR_PAREN, 0)
}

func (s *ArithmeticContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserNUMBER, 0)
}

func (s *ArithmeticContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *ArithmeticContext) MUL() antlr.TerminalNode {
	return s.GetToken(MyScriptParserMUL, 0)
}

func (s *ArithmeticContext) DIV() antlr.TerminalNode {
	return s.GetToken(MyScriptParserDIV, 0)
}

func (s *ArithmeticContext) ADD() antlr.TerminalNode {
	return s.GetToken(MyScriptParserADD, 0)
}

func (s *ArithmeticContext) SUB() antlr.TerminalNode {
	return s.GetToken(MyScriptParserSUB, 0)
}

func (s *ArithmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterArithmetic(s)
	}
}

func (s *ArithmeticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitArithmetic(s)
	}
}

func (p *MyScriptParser) Arithmetic() (localctx IArithmeticContext) {
	return p.arithmetic(0)
}

func (p *MyScriptParser) arithmetic(_p int) (localctx IArithmeticContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArithmeticContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithmeticContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, MyScriptParserRULE_arithmetic, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MyScriptParserL_PAREN:
		{
			p.SetState(55)
			p.Match(MyScriptParserL_PAREN)
		}
		{
			p.SetState(56)
			p.arithmetic(0)
		}
		{
			p.SetState(57)
			p.Match(MyScriptParserR_PAREN)
		}

	case MyScriptParserNUMBER:
		{
			p.SetState(59)
			p.Match(MyScriptParserNUMBER)
		}

	case MyScriptParserIDENTIFIER:
		{
			p.SetState(60)
			p.Match(MyScriptParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_arithmetic)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(64)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MyScriptParserMUL || _la == MyScriptParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(65)
					p.arithmetic(6)
				}

			case 2:
				localctx = NewArithmeticContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_arithmetic)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(67)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MyScriptParserADD || _la == MyScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(68)
					p.arithmetic(5)
				}

			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// ILogicStmtContext is an interface to support dynamic dispatch.
type ILogicStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsLogicStmtContext differentiates from other interfaces.
	IsLogicStmtContext()
}

type LogicStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyLogicStmtContext() *LogicStmtContext {
	var p = new(LogicStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_logicStmt
	return p
}

func (*LogicStmtContext) IsLogicStmtContext() {}

func NewLogicStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicStmtContext {
	var p = new(LogicStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_logicStmt

	return p
}

func (s *LogicStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicStmtContext) GetOp() antlr.Token { return s.op }

func (s *LogicStmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicStmtContext) AllRelation() []IRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationContext); ok {
			len++
		}
	}

	tst := make([]IRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationContext); ok {
			tst[i] = t.(IRelationContext)
			i++
		}
	}

	return tst
}

func (s *LogicStmtContext) Relation(i int) IRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *LogicStmtContext) AllLOGICAL_OR() []antlr.TerminalNode {
	return s.GetTokens(MyScriptParserLOGICAL_OR)
}

func (s *LogicStmtContext) LOGICAL_OR(i int) antlr.TerminalNode {
	return s.GetToken(MyScriptParserLOGICAL_OR, i)
}

func (s *LogicStmtContext) AllLOGICAL_AND() []antlr.TerminalNode {
	return s.GetTokens(MyScriptParserLOGICAL_AND)
}

func (s *LogicStmtContext) LOGICAL_AND(i int) antlr.TerminalNode {
	return s.GetToken(MyScriptParserLOGICAL_AND, i)
}

func (s *LogicStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterLogicStmt(s)
	}
}

func (s *LogicStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitLogicStmt(s)
	}
}

func (p *MyScriptParser) LogicStmt() (localctx ILogicStmtContext) {
	this := p
	_ = this

	localctx = NewLogicStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyScriptParserRULE_logicStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Relation()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MyScriptParserLOGICAL_OR || _la == MyScriptParserLOGICAL_AND {
		{
			p.SetState(75)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LogicStmtContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MyScriptParserLOGICAL_OR || _la == MyScriptParserLOGICAL_AND) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LogicStmtContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(76)
			p.Relation()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) GetOp() antlr.Token { return s.op }

func (s *RelationContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(MyScriptParserEXCLAMATION, 0)
}

func (s *RelationContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *RelationContext) AllArithmetic() []IArithmeticContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmeticContext); ok {
			len++
		}
	}

	tst := make([]IArithmeticContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmeticContext); ok {
			tst[i] = t.(IArithmeticContext)
			i++
		}
	}

	return tst
}

func (s *RelationContext) Arithmetic(i int) IArithmeticContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *RelationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserEQUALS, 0)
}

func (s *RelationContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserNOT_EQUALS, 0)
}

func (s *RelationContext) LESS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserLESS, 0)
}

func (s *RelationContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserLESS_OR_EQUALS, 0)
}

func (s *RelationContext) GREATER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserGREATER, 0)
}

func (s *RelationContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserGREATER_OR_EQUALS, 0)
}

func (s *RelationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserL_PAREN, 0)
}

func (s *RelationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserR_PAREN, 0)
}

func (s *RelationContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MyScriptParserBOOL, 0)
}

func (s *RelationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *MyScriptParser) Relation() (localctx IRelationContext) {
	this := p
	_ = this

	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyScriptParserRULE_relation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(MyScriptParserEXCLAMATION)
		}
		{
			p.SetState(83)
			p.Relation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.arithmetic(0)
		}
		{
			p.SetState(85)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*RelationContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MyScriptParserEQUALS)|(1<<MyScriptParserNOT_EQUALS)|(1<<MyScriptParserLESS)|(1<<MyScriptParserLESS_OR_EQUALS)|(1<<MyScriptParserGREATER)|(1<<MyScriptParserGREATER_OR_EQUALS))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*RelationContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(86)
			p.arithmetic(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(MyScriptParserL_PAREN)
		}

		{
			p.SetState(89)
			p.Relation()
		}

		{
			p.SetState(90)
			p.Match(MyScriptParserR_PAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Match(MyScriptParserBOOL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Match(MyScriptParserIDENTIFIER)
		}

	}

	return localctx
}

// IFuncStmtContext is an interface to support dynamic dispatch.
type IFuncStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncStmtContext differentiates from other interfaces.
	IsFuncStmtContext()
}

type FuncStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncStmtContext() *FuncStmtContext {
	var p = new(FuncStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_funcStmt
	return p
}

func (*FuncStmtContext) IsFuncStmtContext() {}

func NewFuncStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStmtContext {
	var p = new(FuncStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_funcStmt

	return p
}

func (s *FuncStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *FuncStmtContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserL_PAREN, 0)
}

func (s *FuncStmtContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FuncStmtContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserR_PAREN, 0)
}

func (s *FuncStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterFuncStmt(s)
	}
}

func (s *FuncStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitFuncStmt(s)
	}
}

func (p *MyScriptParser) FuncStmt() (localctx IFuncStmtContext) {
	this := p
	_ = this

	localctx = NewFuncStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyScriptParserRULE_funcStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(MyScriptParserIDENTIFIER)
	}
	{
		p.SetState(97)
		p.Match(MyScriptParserL_PAREN)
	}
	{
		p.SetState(98)
		p.ParamList()
	}
	{
		p.SetState(99)
		p.Match(MyScriptParserR_PAREN)
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyScriptParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyScriptParserCOMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *MyScriptParser) ParamList() (localctx IParamListContext) {
	this := p
	_ = this

	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MyScriptParserRULE_paramList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Param()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MyScriptParserCOMMA {
		{
			p.SetState(102)
			p.Match(MyScriptParserCOMMA)
		}
		{
			p.SetState(103)
			p.Param()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *ParamContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MyScriptParserBOOL, 0)
}

func (s *ParamContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserNUMBER, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *MyScriptParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MyScriptParserRULE_param)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MyScriptParserNUMBER)|(1<<MyScriptParserIDENTIFIER)|(1<<MyScriptParserBOOL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *MyScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ArithmeticContext = nil
		if localctx != nil {
			t = localctx.(*ArithmeticContext)
		}
		return p.Arithmetic_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyScriptParser) Arithmetic_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
