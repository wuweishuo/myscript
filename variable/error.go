package variable

type ErrorVariable struct {
	val error
}

func (e ErrorVariable) Value() interface{} {
	return e.val
}

func (e ErrorVariable) IsNull() bool {
	return false
}

func (e ErrorVariable) GetError() error {
	return e.val
}

func (e ErrorVariable) Equal(v Variable) Variable {
	return NewBoolVariable(e.Value() == v.Value())
}

func (e ErrorVariable) NotEqual(v Variable) Variable {
	return NewBoolVariable(e.Value() == v.Value())
}
