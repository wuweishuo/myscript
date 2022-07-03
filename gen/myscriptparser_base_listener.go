// Code generated from D:/GoProject/myscript\MyScriptParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

// EnterArithmetic is called when production arithmetic is entered.
func (s *BaseMyScriptParserListener) EnterArithmetic(ctx *ArithmeticContext) {}

// ExitArithmetic is called when production arithmetic is exited.
func (s *BaseMyScriptParserListener) ExitArithmetic(ctx *ArithmeticContext) {}

// EnterLogicStmt is called when production logicStmt is entered.
func (s *BaseMyScriptParserListener) EnterLogicStmt(ctx *LogicStmtContext) {}

// ExitLogicStmt is called when production logicStmt is exited.
func (s *BaseMyScriptParserListener) ExitLogicStmt(ctx *LogicStmtContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseMyScriptParserListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseMyScriptParserListener) ExitRelation(ctx *RelationContext) {}

// EnterFuncStmt is called when production funcStmt is entered.
func (s *BaseMyScriptParserListener) EnterFuncStmt(ctx *FuncStmtContext) {}

// ExitFuncStmt is called when production funcStmt is exited.
func (s *BaseMyScriptParserListener) ExitFuncStmt(ctx *FuncStmtContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseMyScriptParserListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseMyScriptParserListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseMyScriptParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseMyScriptParserListener) ExitParam(ctx *ParamContext) {}
