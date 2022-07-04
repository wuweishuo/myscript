package myscript

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
	parser "myscript/gen"
	"reflect"
)

type Shell struct {
	functions map[string]Function
}

func NewShell() Shell {
	return Shell{
		functions: map[string]Function{},
	}
}

func (c Shell) Compile(text string) (script Script, err error) {
	stream := antlr.NewInputStream(text)
	lexer := parser.NewMyScriptLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMyScriptParser(tokenStream)
	return Script{
		ast:       p.StatementList(),
		functions: c.functions,
	}, nil

}

func (c Shell) PutFunction(name string, fn interface{}) error {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		return errors.New("not support type")
	}
	var argumentTypeList []FieldType
	for i := 0; i < t.NumIn(); i++ {
		fieldType, err := c.convertType(t.In(i))
		if err != nil {
			return err
		}
		if fieldType == ErrorFieldType {
			return errors.New("not support type")
		}
		argumentTypeList = append(argumentTypeList, fieldType)
	}
	var returnTypeList []FieldType
	for i := 0; i < t.NumOut(); i++ {
		fieldType, err := c.convertType(t.Out(i))
		if err != nil {
			return err
		}
		returnTypeList = append(returnTypeList, fieldType)
	}
	c.functions[name] = Function{
		fn:               fn,
		argumentTypeList: argumentTypeList,
		returnTypeList:   returnTypeList,
	}
	return nil
}

func (c Shell) convertType(t reflect.Type) (FieldType, error) {
	kind := t.Kind()
	switch kind {
	case reflect.String:
		return StringFieldType, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return NumberFieldType, nil
	case reflect.Bool:
		return BoolFieldType, nil
	default:
		if t.Implements(reflect.TypeOf(error(nil))) {
			return ErrorFieldType, nil
		}
		return 0, errors.New("not support type")
	}
}

type Script struct {
	ast           parser.IStatementListContext
	functions     map[string]Function
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
	return c.variableStack.SetVariable(name, val)
}

func (c Script) GetResult(name string) Field {
	return c.variableStack.GetVariable(name)
}
