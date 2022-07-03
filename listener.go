package main

import parser "myscript/gen"

type Listener struct {
	*parser.BaseMyScriptParserListener
}

func (l *Listener) EnterStatement(c *parser.StatementContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) EnterNumber(c *parser.NumberContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) EnterMulDiv(c *parser.MulDivContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) EnterAddSub(c *parser.AddSubContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) EnterOperant(c *parser.OperantContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) EnterLogical(c *parser.LogicalContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) ExitStatement(c *parser.StatementContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) ExitNumber(c *parser.NumberContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) ExitMulDiv(c *parser.MulDivContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) ExitAddSub(c *parser.AddSubContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) ExitOperant(c *parser.OperantContext) {
	//TODO implement me
	panic("implement me")
}

func (l *Listener) ExitLogical(c *parser.LogicalContext) {
	//TODO implement me
	panic("implement me")
}
