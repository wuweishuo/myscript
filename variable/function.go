package variable

type Function = func(args ...Variable) (Variable, error)

type FunctionVariable struct {
	fn Function
}

func NewFunctionVariable(fn Function) (*FunctionVariable, error) {
	return &FunctionVariable{
		fn: fn,
	}, nil
}

func (f FunctionVariable) Value() interface{} {
	return f.fn
}

func (f FunctionVariable) IsNull() bool {
	return false
}

func (f FunctionVariable) Equal(v Variable) Variable {
	return False
}

func (f FunctionVariable) NotEqual(v Variable) Variable {
	return True
}

func (f FunctionVariable) Call(args ...Variable) (Variable, error) {
	return f.fn(args...)
}
