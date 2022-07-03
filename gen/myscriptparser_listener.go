// Code generated from D:/GoProject/myscript\MyScriptParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MyScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MyScriptParserListener is a complete listener for a parse tree produced by MyScriptParser.
type MyScriptParserListener interface {
	antlr.ParseTreeListener

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterArithmetic is called when entering the arithmetic production.
	EnterArithmetic(c *ArithmeticContext)

	// EnterLogicStmt is called when entering the logicStmt production.
	EnterLogicStmt(c *LogicStmtContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterFuncStmt is called when entering the funcStmt production.
	EnterFuncStmt(c *FuncStmtContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitArithmetic is called when exiting the arithmetic production.
	ExitArithmetic(c *ArithmeticContext)

	// ExitLogicStmt is called when exiting the logicStmt production.
	ExitLogicStmt(c *LogicStmtContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitFuncStmt is called when exiting the funcStmt production.
	ExitFuncStmt(c *FuncStmtContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)
}
