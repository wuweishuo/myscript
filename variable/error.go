package variable

import "github.com/pkg/errors"

type ErrorVariable struct {
	val error
}

func NewErrorVariable(start, end int, val error) *ErrorVariable {
	return &ErrorVariable{
		val: errors.Errorf("execute err, start:%d, end:%d, err:%+v", start, end, val),
	}
}

func NewBuildErrorVariable(start, end int) *ErrorVariable {
	return &ErrorVariable{
		val: errors.Errorf("build err, start:%d, end:%d", start, end),
	}
}

func NewUnExpectedTypeErrorVariable(start, end int) *ErrorVariable {
	return &ErrorVariable{
		val: errors.Errorf("unexpected type err, start:%d, end:%d", start, end),
	}
}

func NewUnExpectedKeywordErrorVariable(keyword string, start, end int) *ErrorVariable {
	return &ErrorVariable{
		val: errors.Errorf("unexpected keyword:%s , start:%d, end:%d", keyword, start, end),
	}
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
