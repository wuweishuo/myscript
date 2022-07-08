package myscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "myscript/gen"
)

type TempVariables struct {
	stack []Variable
}

func (v *TempVariables) Push(field Variable) {
	v.stack = append(v.stack, field)
}

func (v *TempVariables) Pop() Variable {
	if len(v.stack) == 0 {
		return NullVariableInstance
	}
	field := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	return field
}

type Listener struct {
	parser.BaseMyScriptParserListener
	script        Script
	tempVariables TempVariables
}

func (l Listener) VisitTerminal(node antlr.TerminalNode) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) VisitErrorNode(node antlr.ErrorNode) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterStatementList(c *parser.StatementListContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterStatement(c *parser.StatementContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterAssignStmt(c *parser.AssignStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterIfStmt(c *parser.IfStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterBlock(c *parser.BlockContext) {
	l.script.variableStack.PushVariables()
}

func (l Listener) EnterFuncStmt(c *parser.FuncStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterExpressionSequence(c *parser.ExpressionSequenceContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterParenthesizedExpression(c *parser.ParenthesizedExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterAdditiveExpression(c *parser.AdditiveExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterRelationalExpression(c *parser.RelationalExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterLiteralExpression(c *parser.LiteralExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterNotExpression(c *parser.NotExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterEqualityExpression(c *parser.EqualityExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterIdentifierExpression(c *parser.IdentifierExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) EnterLiteral(c *parser.LiteralContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitStatementList(c *parser.StatementListContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitStatement(c *parser.StatementContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitAssignStmt(c *parser.AssignStmtContext) {
	variableName := c.IDENTIFIER().GetText()
	variable := l.tempVariables.Pop()
	err := l.script.variableStack.SetVariable(variableName, variable)
	if err != nil {
		panic(err)
	}
}

func (l Listener) ExitIfStmt(c *parser.IfStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitBlock(c *parser.BlockContext) {
	l.script.variableStack.PopVariables()
}

func (l Listener) ExitFuncStmt(c *parser.FuncStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitExpressionSequence(c *parser.ExpressionSequenceContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitAdditiveExpression(c *parser.AdditiveExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitRelationalExpression(c *parser.RelationalExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitNotExpression(c *parser.NotExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitEqualityExpression(c *parser.EqualityExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	right := l.tempVariables.Pop()
	left := l.tempVariables.Pop()
	if right.FieldType() != NumberFieldType || left.FieldType() != NullFieldType {
		panic(NotSupportTypeErr)
	}

	//TODO implement me
	panic("implement me")
}

func (l Listener) ExitIdentifierExpression(c *parser.IdentifierExpressionContext) {
	variableName := c.IDENTIFIER().GetText()
	variable := l.script.variableStack.GetVariable(variableName)
	l.tempVariables.Push(variable)
}

func (l Listener) ExitLiteral(c *parser.LiteralContext) {
	if c.NUMBER() != nil {
		variable, err := NewVariableFromString(NumberFieldType, c.NUMBER().GetText())
		if err != nil {
			panic(err)
		}
		l.tempVariables.Push(variable)
	}
	if c.BOOL() != nil {
		variable, err := NewVariableFromString(BoolFieldType, c.BOOL().GetText())
		if err != nil {
			panic(err)
		}
		l.tempVariables.Push(variable)
	}
}
