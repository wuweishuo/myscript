package variable

import (
	"github.com/pkg/errors"
)

var NotSupportTypeErr = errors.New("not support type")

type Variable interface {
	Value() interface{}
	IsNull() bool
	Equal(v Variable) Variable
	NotEqual(v Variable) Variable
}

func NewVariableFromInterface(val interface{}) (Variable, error) {
	switch v := val.(type) {
	case string:
		return NewStringVariable(v), nil
	case bool:
		return NewBoolVariable(v), nil
	case int:
		return NewNumberVariableFromInt(int64(v)), nil
	case int64:
		return NewNumberVariableFromInt(v), nil
	case float32:
		return NewNumberVariableFromFloat(float64(v)), nil
	case float64:
		return NewNumberVariableFromFloat(v), nil
	case Function:
		return NewFunctionVariable(val.(Function))
	default:
		return nil, NotSupportTypeErr
	}
}
