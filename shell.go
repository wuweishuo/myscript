package myscript

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
	parser "myscript/gen"
	"myscript/variable"
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
		ast:           p.StatementList(),
		variableStack: variable.NewStack(),
	}, nil

}

type Script struct {
	ast           parser.IStatementListContext
	variableStack variable.Stack
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
	v, err := variable.NewVariableFromInterface(val)
	if err != nil {
		return err
	}
	c.variableStack.SetVariable(name, v)
	return nil
}

func (c Script) GetResult(name string) variable.Variable {
	return c.variableStack.GetVariable(name)
}
