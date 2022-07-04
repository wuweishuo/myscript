// Code generated from /Users/bytedance/GolandProjects/myscript/MyScriptParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MyScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMyScriptParserListener is a complete listener for a parse tree produced by MyScriptParser.
type BaseMyScriptParserListener struct{}

var _ MyScriptParserListener = &BaseMyScriptParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMyScriptParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMyScriptParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMyScriptParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMyScriptParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseMyScriptParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseMyScriptParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMyScriptParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMyScriptParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseMyScriptParserListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseMyScriptParserListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseMyScriptParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseMyScriptParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMyScriptParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMyScriptParserListener) ExitBlock(ctx *BlockContext) {}

// EnterFuncStmt is called when production funcStmt is entered.
func (s *BaseMyScriptParserListener) EnterFuncStmt(ctx *FuncStmtContext) {}

// ExitFuncStmt is called when production funcStmt is exited.
func (s *BaseMyScriptParserListener) ExitFuncStmt(ctx *FuncStmtContext) {}

// EnterExpressionSequence is called when production expressionSequence is entered.
func (s *BaseMyScriptParserListener) EnterExpressionSequence(ctx *ExpressionSequenceContext) {}

// ExitExpressionSequence is called when production expressionSequence is exited.
func (s *BaseMyScriptParserListener) ExitExpressionSequence(ctx *ExpressionSequenceContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseMyScriptParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseMyScriptParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseMyScriptParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseMyScriptParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseMyScriptParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseMyScriptParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseMyScriptParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseMyScriptParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseMyScriptParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseMyScriptParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseMyScriptParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseMyScriptParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseMyScriptParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseMyScriptParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseMyScriptParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseMyScriptParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseMyScriptParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseMyScriptParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMyScriptParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMyScriptParserListener) ExitLiteral(ctx *LiteralContext) {}
