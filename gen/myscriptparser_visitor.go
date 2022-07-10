// Code generated from D:/GoProject/myscript\MyScriptParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MyScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MyScriptParser.
type MyScriptParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MyScriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by MyScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MyScriptParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by MyScriptParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by MyScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MyScriptParser#expressionSequence.
	VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{}

	// Visit a parse tree produced by MyScriptParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#AdditiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#RelationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#LogicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#FuncExpression.
	VisitFuncExpression(ctx *FuncExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#LogicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#MultiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by MyScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
