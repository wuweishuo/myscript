package myscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "myscript/gen"
	"myscript/variable"
)

type Visitor struct {
	parser.BaseMyScriptParserVisitor
	variableStack variable.Stack
}

func NewVisitor(variableStack variable.Stack) *Visitor {
	return &Visitor{
		variableStack: variableStack,
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, tree := range node.GetChildren() {
		if parserTree, ok := tree.(antlr.ParseTree); ok {
			res := parserTree.Accept(v)
			if _, ok := res.(*variable.ErrorVariable); ok {
				return res
			}
		}
	}
	return variable.Null
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return variable.Null
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return variable.Null
}

func (v *Visitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	res := v.VisitChildren(ctx)
	if _, ok := res.(*variable.ErrorVariable); ok {
		return res
	}
	return variable.Null
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	res := v.VisitChildren(ctx)
	if _, ok := res.(*variable.ErrorVariable); ok {
		return res
	}
	return variable.Null
}

func (v *Visitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	variableName := ctx.IDENTIFIER().GetText()
	res := ctx.SingleExpression().Accept(v).(variable.Variable)
	if _, ok := res.(*variable.ErrorVariable); ok {
		return res
	}
	v.variableStack.SetVariable(variableName, res)
	return variable.Null
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	res := ctx.SingleExpression().Accept(v)
	switch t := res.(type) {
	case *variable.ErrorVariable:
		return res
	case *variable.BoolVariable:
		if t.GetBool() {
			return ctx.Block(0).Accept(v)
		} else {
			if ctx.IfStmt() != nil {
				return ctx.IfStmt().Accept(v)
			} else {
				return ctx.Block(1).Accept(v)
			}
		}
	default:
		return variable.NewBuildErrorVariable(ctx.SingleExpression().GetStart().GetStart(), ctx.SingleExpression().GetStop().GetStop())
	}
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.variableStack.PushVariables()
	res := v.VisitChildren(ctx.StatementList())
	if _, ok := res.(*variable.ErrorVariable); ok {
		return res
	}
	v.variableStack.PopVariables()
	return variable.Null
}

func (v *Visitor) VisitParenthesizedExpression(ctx *parser.ParenthesizedExpressionContext) interface{} {
	return ctx.SingleExpression().Accept(v)
}

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	left := ctx.SingleExpression(0).Accept(v)
	if e, ok := left.(*variable.ErrorVariable); ok {
		return e
	}
	number1, ok := left.(*variable.NumberVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(0).GetStart().GetStart(), ctx.SingleExpression(0).GetStop().GetStop())
	}

	right := ctx.SingleExpression(1).Accept(v)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	number2, ok := right.(*variable.NumberVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(1).GetStart().GetStart(), ctx.SingleExpression(1).GetStop().GetStop())
	}
	op := ctx.GetChild(1).(antlr.TerminalNode)
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserADD {
		return number1.Add(number2)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserSUB {
		return number1.Sub(number2)
	}
	return variable.NewUnExpectedKeywordErrorVariable(op.GetSymbol().GetText(), op.GetSymbol().GetStart(), op.GetSymbol().GetStop())
}

func (v *Visitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	left := ctx.SingleExpression(0).Accept(v)
	if e, ok := left.(*variable.ErrorVariable); ok {
		return e
	}
	number1, ok := left.(*variable.NumberVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(0).GetStart().GetStart(), ctx.SingleExpression(0).GetStop().GetStop())
	}

	right := ctx.SingleExpression(1).Accept(v)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	number2, ok := right.(*variable.NumberVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(1).GetStart().GetStart(), ctx.SingleExpression(1).GetStop().GetStop())
	}
	op := ctx.GetChild(1).(antlr.TerminalNode)
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserGREATER {
		return number1.Greater(number2)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserGREATER_OR_EQUALS {
		return number1.GreaterOrEqual(number2)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserLESS {
		return number1.Less(number2)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserLESS_OR_EQUALS {
		return number1.LessOrEqual(number2)
	}
	return variable.NewUnExpectedKeywordErrorVariable(op.GetSymbol().GetText(), op.GetSymbol().GetStart(), op.GetSymbol().GetStop())
}

func (v *Visitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	left := ctx.SingleExpression(0).Accept(v)
	if e, ok := left.(*variable.ErrorVariable); ok {
		return e
	}
	bool1, ok := left.(*variable.BoolVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(0).GetStart().GetStart(), ctx.SingleExpression(0).GetStop().GetStop())
	}

	right := ctx.SingleExpression(1).Accept(v)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	bool2, ok := right.(*variable.BoolVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(1).GetStart().GetStart(), ctx.SingleExpression(1).GetStop().GetStop())
	}
	return bool1.And(bool2)
}

func (v *Visitor) VisitFuncExpression(ctx *parser.FuncExpressionContext) interface{} {
	var args []variable.Variable
	for _, singleExpr := range ctx.ExpressionSequence().(*parser.ExpressionSequenceContext).AllSingleExpression() {
		arg := singleExpr.Accept(v).(variable.Variable)
		if e, ok := arg.(*variable.ErrorVariable); ok {
			return e
		}
		args = append(args, arg)
	}
	fn := v.variableStack.GetVariable(ctx.IDENTIFIER().GetText())
	if e, ok := fn.(*variable.ErrorVariable); ok {
		return e
	}
	fnVariable, ok := fn.(*variable.FunctionVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.IDENTIFIER().GetSymbol().GetStart(), ctx.IDENTIFIER().GetSymbol().GetStop())
	}
	call, err := fnVariable.Call(args...)
	if err != nil {
		return variable.NewErrorVariable(ctx.GetStart().GetStart(), ctx.GetStop().GetStop(), err)
	}
	return call
}

func (v *Visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {
	return ctx.Literal().Accept(v)
}

func (v *Visitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	left := ctx.SingleExpression(0).Accept(v)
	if e, ok := left.(*variable.ErrorVariable); ok {
		return e
	}
	bool1, ok := left.(*variable.BoolVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(0).GetStart().GetStart(), ctx.SingleExpression(0).GetStop().GetStop())
	}

	right := ctx.SingleExpression(1).Accept(v)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	bool2, ok := right.(*variable.BoolVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(1).GetStart().GetStart(), ctx.SingleExpression(1).GetStop().GetStop())
	}
	return bool1.Or(bool2)
}

func (v *Visitor) VisitNotExpression(ctx *parser.NotExpressionContext) interface{} {
	right := ctx.SingleExpression().Accept(v)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	bool1, ok := right.(*variable.BoolVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression().GetStart().GetStart(), ctx.SingleExpression().GetStop().GetStop())
	}
	return bool1.Not()
}

func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	left := ctx.SingleExpression(0).Accept(v).(variable.Variable)
	if e, ok := left.(*variable.ErrorVariable); ok {
		return e
	}

	right := ctx.SingleExpression(1).Accept(v).(variable.Variable)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	op := ctx.GetChild(1).(antlr.TerminalNode)
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserEQUALS {
		return left.Equal(right)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserDIV {
		return left.NotEqual(right)
	}
	return variable.NewUnExpectedKeywordErrorVariable(op.GetSymbol().GetText(), op.GetSymbol().GetStart(), op.GetSymbol().GetStop())
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	left := ctx.SingleExpression(0).Accept(v)
	if e, ok := left.(*variable.ErrorVariable); ok {
		return e
	}
	number1, ok := left.(*variable.NumberVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(0).GetStart().GetStart(), ctx.SingleExpression(0).GetStop().GetStop())
	}

	right := ctx.SingleExpression(1).Accept(v)
	if e, ok := right.(*variable.ErrorVariable); ok {
		return e
	}
	number2, ok := right.(*variable.NumberVariable)
	if !ok {
		return variable.NewUnExpectedTypeErrorVariable(ctx.SingleExpression(1).GetStart().GetStart(), ctx.SingleExpression(1).GetStop().GetStop())
	}
	op := ctx.GetChild(1).(antlr.TerminalNode)
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserMUL {
		return number1.Mul(number2)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserDIV {
		return number1.Div(number2)
	}
	if op.GetSymbol().GetTokenType() == parser.MyScriptParserMOD {
		return number1.Mod(number2)
	}
	return variable.NewUnExpectedKeywordErrorVariable(op.GetSymbol().GetText(), op.GetSymbol().GetStart(), op.GetSymbol().GetStop())
}

func (v *Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	variableName := ctx.IDENTIFIER().GetText()
	return v.variableStack.GetVariable(variableName)
}

func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.NUMBER() != nil {
		v, err := variable.NewNumberVariableFromString(ctx.NUMBER().GetText())
		if err != nil {
			return variable.NewErrorVariable(ctx.GetStart().GetStart(), ctx.GetStop().GetStop(), err)
		}
		return v
	}
	if ctx.BOOL() != nil {
		v, err := variable.NewBoolVariableFromString(ctx.BOOL().GetText())
		if err != nil {
			return variable.NewErrorVariable(ctx.GetStart().GetStart(), ctx.GetStop().GetStop(), err)
		}
		return v
	}
	return variable.NewBuildErrorVariable(ctx.GetStart().GetStart(), ctx.GetStop().GetStop())
}
