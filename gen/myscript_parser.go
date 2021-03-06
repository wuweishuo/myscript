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
		"'/'", "'+'", "'-'", "'%'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
		"'||'", "'&&'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "ELSE", "IF", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "ASSIGN",
		"COMMA", "MUL", "DIV", "ADD", "SUB", "MOD", "EQUALS", "NOT_EQUALS",
		"LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "LOGICAL_OR",
		"LOGICAL_AND", "EXCLAMATION", "BOOL", "NUMBER", "IDENTIFIER", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"statementList", "statement", "assignStmt", "ifStmt", "block", "expressionSequence",
		"singleExpression", "literal",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 18, 8, 0, 11, 0, 12,
		0, 19, 1, 1, 1, 1, 3, 1, 24, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 36, 8, 3, 3, 3, 38, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 47, 8, 5, 10, 5, 12, 5, 50, 9, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 66, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 86, 8, 6,
		10, 6, 12, 6, 89, 9, 6, 1, 7, 1, 7, 1, 7, 0, 1, 12, 8, 0, 2, 4, 6, 8, 10,
		12, 14, 0, 5, 2, 0, 9, 10, 13, 13, 1, 0, 11, 12, 1, 0, 16, 19, 1, 0, 14,
		15, 1, 0, 23, 24, 99, 0, 17, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 25, 1, 0,
		0, 0, 6, 29, 1, 0, 0, 0, 8, 39, 1, 0, 0, 0, 10, 43, 1, 0, 0, 0, 12, 65,
		1, 0, 0, 0, 14, 90, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0,
		18, 19, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1, 1, 0,
		0, 0, 21, 24, 3, 4, 2, 0, 22, 24, 3, 6, 3, 0, 23, 21, 1, 0, 0, 0, 23, 22,
		1, 0, 0, 0, 24, 3, 1, 0, 0, 0, 25, 26, 5, 25, 0, 0, 26, 27, 5, 7, 0, 0,
		27, 28, 3, 12, 6, 0, 28, 5, 1, 0, 0, 0, 29, 30, 5, 2, 0, 0, 30, 31, 3,
		12, 6, 0, 31, 37, 3, 8, 4, 0, 32, 35, 5, 1, 0, 0, 33, 36, 3, 6, 3, 0, 34,
		36, 3, 8, 4, 0, 35, 33, 1, 0, 0, 0, 35, 34, 1, 0, 0, 0, 36, 38, 1, 0, 0,
		0, 37, 32, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 7, 1, 0, 0, 0, 39, 40, 5,
		5, 0, 0, 40, 41, 3, 0, 0, 0, 41, 42, 5, 6, 0, 0, 42, 9, 1, 0, 0, 0, 43,
		48, 3, 12, 6, 0, 44, 45, 5, 8, 0, 0, 45, 47, 3, 12, 6, 0, 46, 44, 1, 0,
		0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 11,
		1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 52, 6, 6, -1, 0, 52, 53, 5, 22, 0,
		0, 53, 66, 3, 12, 6, 11, 54, 66, 3, 14, 7, 0, 55, 56, 5, 25, 0, 0, 56,
		57, 5, 3, 0, 0, 57, 58, 3, 10, 5, 0, 58, 59, 5, 4, 0, 0, 59, 66, 1, 0,
		0, 0, 60, 66, 5, 25, 0, 0, 61, 62, 5, 3, 0, 0, 62, 63, 3, 12, 6, 0, 63,
		64, 5, 4, 0, 0, 64, 66, 1, 0, 0, 0, 65, 51, 1, 0, 0, 0, 65, 54, 1, 0, 0,
		0, 65, 55, 1, 0, 0, 0, 65, 60, 1, 0, 0, 0, 65, 61, 1, 0, 0, 0, 66, 87,
		1, 0, 0, 0, 67, 68, 10, 10, 0, 0, 68, 69, 7, 0, 0, 0, 69, 86, 3, 12, 6,
		11, 70, 71, 10, 9, 0, 0, 71, 72, 7, 1, 0, 0, 72, 86, 3, 12, 6, 10, 73,
		74, 10, 8, 0, 0, 74, 75, 7, 2, 0, 0, 75, 86, 3, 12, 6, 9, 76, 77, 10, 7,
		0, 0, 77, 78, 7, 3, 0, 0, 78, 86, 3, 12, 6, 8, 79, 80, 10, 6, 0, 0, 80,
		81, 5, 21, 0, 0, 81, 86, 3, 12, 6, 7, 82, 83, 10, 5, 0, 0, 83, 84, 5, 20,
		0, 0, 84, 86, 3, 12, 6, 6, 85, 67, 1, 0, 0, 0, 85, 70, 1, 0, 0, 0, 85,
		73, 1, 0, 0, 0, 85, 76, 1, 0, 0, 0, 85, 79, 1, 0, 0, 0, 85, 82, 1, 0, 0,
		0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 13,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 7, 4, 0, 0, 91, 15, 1, 0, 0, 0,
		8, 19, 23, 35, 37, 48, 65, 85, 87,
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
	MyScriptParserMOD               = 13
	MyScriptParserEQUALS            = 14
	MyScriptParserNOT_EQUALS        = 15
	MyScriptParserLESS              = 16
	MyScriptParserLESS_OR_EQUALS    = 17
	MyScriptParserGREATER           = 18
	MyScriptParserGREATER_OR_EQUALS = 19
	MyScriptParserLOGICAL_OR        = 20
	MyScriptParserLOGICAL_AND       = 21
	MyScriptParserEXCLAMATION       = 22
	MyScriptParserBOOL              = 23
	MyScriptParserNUMBER            = 24
	MyScriptParserIDENTIFIER        = 25
	MyScriptParserWHITESPACE        = 26
)

// MyScriptParser rules.
const (
	MyScriptParserRULE_statementList      = 0
	MyScriptParserRULE_statement          = 1
	MyScriptParserRULE_assignStmt         = 2
	MyScriptParserRULE_ifStmt             = 3
	MyScriptParserRULE_block              = 4
	MyScriptParserRULE_expressionSequence = 5
	MyScriptParserRULE_singleExpression   = 6
	MyScriptParserRULE_literal            = 7
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

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MyScriptParserIF || _la == MyScriptParserIDENTIFIER {
		{
			p.SetState(16)
			p.Statement()
		}

		p.SetState(19)
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

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MyScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.AssignStmt()
		}

	case MyScriptParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.IfStmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *AssignStmtContext) SingleExpression() ISingleExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(25)
		p.Match(MyScriptParserIDENTIFIER)
	}
	{
		p.SetState(26)
		p.Match(MyScriptParserASSIGN)
	}
	{
		p.SetState(27)
		p.singleExpression(0)
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

func (s *IfStmtContext) SingleExpression() ISingleExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(29)
		p.Match(MyScriptParserIF)
	}
	{
		p.SetState(30)
		p.singleExpression(0)
	}
	{
		p.SetState(31)
		p.Block()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MyScriptParserELSE {
		{
			p.SetState(32)
			p.Match(MyScriptParserELSE)
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MyScriptParserIF:
			{
				p.SetState(33)
				p.IfStmt()
			}

		case MyScriptParserL_CURLY:
			{
				p.SetState(34)
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

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(39)
		p.Match(MyScriptParserL_CURLY)
	}
	{
		p.SetState(40)
		p.StatementList()
	}
	{
		p.SetState(41)
		p.Match(MyScriptParserR_CURLY)
	}

	return localctx
}

// IExpressionSequenceContext is an interface to support dynamic dispatch.
type IExpressionSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionSequenceContext differentiates from other interfaces.
	IsExpressionSequenceContext()
}

type ExpressionSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSequenceContext() *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_expressionSequence
	return p
}

func (*ExpressionSequenceContext) IsExpressionSequenceContext() {}

func NewExpressionSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_expressionSequence

	return p
}

func (s *ExpressionSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSequenceContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionSequenceContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *ExpressionSequenceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyScriptParserCOMMA)
}

func (s *ExpressionSequenceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyScriptParserCOMMA, i)
}

func (s *ExpressionSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterExpressionSequence(s)
	}
}

func (s *ExpressionSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitExpressionSequence(s)
	}
}

func (s *ExpressionSequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitExpressionSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyScriptParser) ExpressionSequence() (localctx IExpressionSequenceContext) {
	this := p
	_ = this

	localctx = NewExpressionSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyScriptParserRULE_expressionSequence)
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
		p.SetState(43)
		p.singleExpression(0)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MyScriptParserCOMMA {
		{
			p.SetState(44)
			p.Match(MyScriptParserCOMMA)
		}
		{
			p.SetState(45)
			p.singleExpression(0)
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) CopyFrom(ctx *SingleExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenthesizedExpressionContext struct {
	*SingleExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserL_PAREN, 0)
}

func (s *ParenthesizedExpressionContext) SingleExpression() ISingleExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ParenthesizedExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserR_PAREN, 0)
}

func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitParenthesizedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExpressionContext struct {
	*SingleExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *AdditiveExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(MyScriptParserADD, 0)
}

func (s *AdditiveExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(MyScriptParserSUB, 0)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalExpressionContext struct {
	*SingleExpressionContext
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *RelationalExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserGREATER, 0)
}

func (s *RelationalExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserLESS, 0)
}

func (s *RelationalExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserGREATER_OR_EQUALS, 0)
}

func (s *RelationalExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserLESS_OR_EQUALS, 0)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalAndExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *LogicalAndExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(MyScriptParserLOGICAL_AND, 0)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncExpressionContext struct {
	*SingleExpressionContext
}

func NewFuncExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncExpressionContext {
	var p = new(FuncExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *FuncExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *FuncExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserL_PAREN, 0)
}

func (s *FuncExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *FuncExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MyScriptParserR_PAREN, 0)
}

func (s *FuncExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterFuncExpression(s)
	}
}

func (s *FuncExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitFuncExpression(s)
	}
}

func (s *FuncExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitFuncExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalOrExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *LogicalOrExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(MyScriptParserLOGICAL_OR, 0)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExpressionContext struct {
	*SingleExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(MyScriptParserEXCLAMATION, 0)
}

func (s *NotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityExpressionContext struct {
	*SingleExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *EqualityExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserEQUALS, 0)
}

func (s *EqualityExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(MyScriptParserNOT_EQUALS, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicativeExpressionContext struct {
	*SingleExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
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

	return t.(ISingleExpressionContext)
}

func (s *MultiplicativeExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(MyScriptParserMUL, 0)
}

func (s *MultiplicativeExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(MyScriptParserDIV, 0)
}

func (s *MultiplicativeExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(MyScriptParserMOD, 0)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*SingleExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserIDENTIFIER, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyScriptParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *MyScriptParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, MyScriptParserRULE_singleExpression, _p)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(52)
			p.Match(MyScriptParserEXCLAMATION)
		}
		{
			p.SetState(53)
			p.singleExpression(11)
		}

	case 2:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Literal()
		}

	case 3:
		localctx = NewFuncExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.Match(MyScriptParserIDENTIFIER)
		}
		{
			p.SetState(56)
			p.Match(MyScriptParserL_PAREN)
		}
		{
			p.SetState(57)
			p.ExpressionSequence()
		}
		{
			p.SetState(58)
			p.Match(MyScriptParserR_PAREN)
		}

	case 4:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Match(MyScriptParserIDENTIFIER)
		}

	case 5:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(MyScriptParserL_PAREN)
		}
		{
			p.SetState(62)
			p.singleExpression(0)
		}
		{
			p.SetState(63)
			p.Match(MyScriptParserR_PAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_singleExpression)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(68)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MyScriptParserMUL)|(1<<MyScriptParserDIV)|(1<<MyScriptParserMOD))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(69)
					p.singleExpression(11)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_singleExpression)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(71)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MyScriptParserADD || _la == MyScriptParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(72)
					p.singleExpression(10)
				}

			case 3:
				localctx = NewRelationalExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_singleExpression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(74)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MyScriptParserLESS)|(1<<MyScriptParserLESS_OR_EQUALS)|(1<<MyScriptParserGREATER)|(1<<MyScriptParserGREATER_OR_EQUALS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(75)
					p.singleExpression(9)
				}

			case 4:
				localctx = NewEqualityExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_singleExpression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(77)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MyScriptParserEQUALS || _la == MyScriptParserNOT_EQUALS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(78)
					p.singleExpression(8)
				}

			case 5:
				localctx = NewLogicalAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_singleExpression)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(80)
					p.Match(MyScriptParserLOGICAL_AND)
				}
				{
					p.SetState(81)
					p.singleExpression(7)
				}

			case 6:
				localctx = NewLogicalOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MyScriptParserRULE_singleExpression)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(83)
					p.Match(MyScriptParserLOGICAL_OR)
				}
				{
					p.SetState(84)
					p.singleExpression(6)
				}

			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MyScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MyScriptParserBOOL, 0)
}

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MyScriptParserNUMBER, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyScriptParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyScriptParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MyScriptParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyScriptParserRULE_literal)
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
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MyScriptParserBOOL || _la == MyScriptParserNUMBER) {
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
	case 6:
		var t *SingleExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SingleExpressionContext)
		}
		return p.SingleExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyScriptParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
