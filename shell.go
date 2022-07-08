package myscript

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
	parser "myscript/gen"
)

type Shell struct{}

func NewShell() Shell {
	return Shell{}
}

func (c Shell) Compile(text string) (script Script, err error) {
	stream := antlr.NewInputStream(text)
	lexer := parser.NewMyScriptLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMyScriptParser(tokenStream)
	return Script{
		ast: p.StatementList(),
	}, nil

}

type Script struct {
	ast           parser.IStatementListContext
	variableStack VariableStack
}

func (c Script) Eval(ctx context.Context) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	antlr.ParseTreeWalkerDefault.Walk(&Listener{script: c}, c.ast)
	return nil
}

func (c Script) PutVariable(name string, val interface{}) error {
	variable, err := NewVariableFromInterface(val)
	if err != nil {
		return err
	}
	return c.variableStack.SetVariable(name, variable)
}

func (c Script) GetResult(name string) Variable {
	return c.variableStack.GetVariable(name)
}
