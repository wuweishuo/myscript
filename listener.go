package myscript

import (
	"fmt"
	parser "myscript/gen"
	"myscript/variable"
)

type TempVariables struct {
	stack []variable.Variable
}

func (v *TempVariables) Push(field variable.Variable) {
	v.stack = append(v.stack, field)
}

func (v *TempVariables) Pop() variable.Variable {
	if len(v.stack) == 0 {
		return variable.NullVariableInstance
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

func (l *Listener) EnterBlock(c *parser.BlockContext) {
	l.script.variableStack.PushVariables()
}

func (l *Listener) ExitFuncExpression(ctx *parser.FuncExpressionContext) {
	argsExpr := ctx.ExpressionSequence()
	argsCount := argsExpr.GetChildCount()/2 + 1
	args := make([]variable.Variable, argsCount)
	for i := argsCount - 1; i >= 0; i-- {
		args[i] = l.tempVariables.Pop()
	}
	fn := l.script.variableStack.GetVariable(ctx.IDENTIFIER().GetText())
	if f, ok := fn.(*variable.FunctionVariable); ok {
		res, err := f.Call(args...)
		if err != nil {
			panic(err)
		}
		l.tempVariables.Push(res)
		return
	}
	panic(variable.NotSupportTypeErr)
}

func (l *Listener) ExitAssignStmt(c *parser.AssignStmtContext) {
	variableName := c.IDENTIFIER().GetText()
	v := l.tempVariables.Pop()
	l.script.variableStack.SetVariable(variableName, v)
}

func (l *Listener) ExitIfStmt(c *parser.IfStmtContext) {
	fmt.Println(c.GetText())
}

func (l *Listener) ExitBlock(c *parser.BlockContext) {
	l.script.variableStack.PopVariables()
}

func (l *Listener) ExitAdditiveExpression(c *parser.AdditiveExpressionContext) {
	right, rok := l.tempVariables.Pop().(*variable.NumberVariable)
	left, lok := l.tempVariables.Pop().(*variable.NumberVariable)
	if !rok || !lok {
		panic(variable.NotSupportTypeErr)
	}

	if c.ADD() != nil {
		l.tempVariables.Push(left.Add(right))
	}
	if c.SUB() != nil {
		l.tempVariables.Push(left.Sub(right))
	}
}

func (l *Listener) ExitRelationalExpression(c *parser.RelationalExpressionContext) {
	right, rok := l.tempVariables.Pop().(*variable.NumberVariable)
	left, lok := l.tempVariables.Pop().(*variable.NumberVariable)
	if !rok || !lok {
		panic(variable.NotSupportTypeErr)
	}

	if c.GREATER() != nil {
		l.tempVariables.Push(left.Greater(right))
	}
	if c.GREATER_OR_EQUALS() != nil {
		l.tempVariables.Push(left.GreaterOrEqual(right))
	}
	if c.LESS() != nil {
		l.tempVariables.Push(left.Less(right))
	}
	if c.LESS_OR_EQUALS() != nil {
		l.tempVariables.Push(left.LessOrEqual(right))
	}
}

func (l *Listener) ExitLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	right, rok := l.tempVariables.Pop().(*variable.BoolVariable)
	left, lok := l.tempVariables.Pop().(*variable.BoolVariable)
	if !rok || !lok {
		panic(variable.NotSupportTypeErr)
	}
	l.tempVariables.Push(left.And(right))
}

func (l *Listener) ExitLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	right, rok := l.tempVariables.Pop().(*variable.BoolVariable)
	left, lok := l.tempVariables.Pop().(*variable.BoolVariable)
	if !rok || !lok {
		panic(variable.NotSupportTypeErr)
	}
	l.tempVariables.Push(left.Or(right))
}

func (l *Listener) ExitNotExpression(c *parser.NotExpressionContext) {
	v, ok := l.tempVariables.Pop().(*variable.BoolVariable)
	if !ok {
		panic(variable.NotSupportTypeErr)
	}
	l.tempVariables.Push(v.Not())
}

func (l *Listener) ExitEqualityExpression(c *parser.EqualityExpressionContext) {
	right := l.tempVariables.Pop()
	left := l.tempVariables.Pop()
	if c.EQUALS() != nil {
		l.tempVariables.Push(left.Equal(right))
	}
	if c.NOT_EQUALS() != nil {
		l.tempVariables.Push(left.NotEqual(right))
	}
}

func (l *Listener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	right, rok := l.tempVariables.Pop().(*variable.NumberVariable)
	left, lok := l.tempVariables.Pop().(*variable.NumberVariable)
	if !rok || !lok {
		panic(variable.NotSupportTypeErr)
	}

	if c.MUL() != nil {
		l.tempVariables.Push(left.Mul(right))
	}
	if c.DIV() != nil {
		l.tempVariables.Push(left.Div(right))
	}
	if c.MOD() != nil {
		l.tempVariables.Push(left.Mod(right))
	}
}

func (l *Listener) ExitIdentifierExpression(c *parser.IdentifierExpressionContext) {
	variableName := c.IDENTIFIER().GetText()
	v := l.script.variableStack.GetVariable(variableName)
	l.tempVariables.Push(v)
}

func (l *Listener) ExitLiteral(c *parser.LiteralContext) {
	if c.NUMBER() != nil {
		v, err := variable.NewNumberVariableFromString(c.NUMBER().GetText())
		if err != nil {
			panic(err)
		}
		l.tempVariables.Push(v)
	}
	if c.BOOL() != nil {
		v, err := variable.NewBoolVariableFromString(c.BOOL().GetText())
		if err != nil {
			panic(err)
		}
		l.tempVariables.Push(v)
	}
}
