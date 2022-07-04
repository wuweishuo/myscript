package myscript

import "github.com/pkg/errors"

type FieldType int

const (
	NullFieldType   = 0
	StringFieldType = 1
	BoolFieldType   = 2
	NumberFieldType = 3
	ErrorFieldType  = 4
)

type Field struct {
	fieldValue interface{}
	fieldType  FieldType
}

var NullField = Field{
	fieldType: NullFieldType,
}

type Function struct {
	fn               interface{}
	argumentTypeList []FieldType
	returnTypeList   []FieldType
}

type Variables map[string]Field

func (v Variables) Put(name string, val interface{}) error {
	var field Field
	switch val.(type) {
	case string:
		field = Field{
			fieldType:  StringFieldType,
			fieldValue: val,
		}
	case bool:
		field = Field{
			fieldType:  BoolFieldType,
			fieldValue: val,
		}
	case int, int64, int32, int16, int8, float32, float64, uint, uint8, uint16, uint32, uint64:
		field = Field{
			fieldType:  NumberFieldType,
			fieldValue: val,
		}
	default:
		return errors.New("not support type")
	}
	v[name] = field
	return nil
}

func (v Variables) Get(name string) Field {
	return v[name]
}

type VariableStack struct {
	stack []Variables
}

func (s VariableStack) SetVariable(name string, val interface{}) error {
	for i := len(s.stack) - 1; i >= 0; i++ {
		if s.stack[i].Get(name) != NullField {
			return s.stack[i].Put(name, val)
		}
	}
	return s.stack[len(s.stack)-1].Put(name, val)
}

func (s VariableStack) GetVariable(name string) Field {
	for i := len(s.stack) - 1; i >= 0; i++ {
		field := s.stack[i].Get(name)
		if field != NullField {
			return field
		}
	}
	return NullField
}

func (s VariableStack) PushVariables() VariableStack {
	s.stack = append(s.stack, map[string]Field{})
	return s
}

func (s VariableStack) PopVariables() VariableStack {
	s.stack = s.stack[:len(s.stack)-1]
	return s
}
