package myscript

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"reflect"
)

var NotSupportTypeErr = errors.New("not support type")

type FieldType int

const (
	NullFieldType     = 0
	StringFieldType   = 1
	BoolFieldType     = 2
	NumberFieldType   = 3
	ErrorFieldType    = 4
	FunctionFieldType = 5
)

func GetFieldTypeFromReflect(t reflect.Type) (FieldType, error) {
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
		return 0, NotSupportTypeErr
	}
}

type Variable interface {
	FieldType() FieldType
	Value() interface{}
	IsNull() bool
}

type NullVariable struct{}

var NullVariableInstance = NullVariable{}

func (n NullVariable) IsNull() bool {
	return true
}

func (n NullVariable) FieldType() FieldType {
	return NullFieldType
}

func (n NullVariable) Value() interface{} {
	return nil
}

type StringVariable struct {
	val string
}

func (s StringVariable) IsNull() bool {
	return false
}

func (s StringVariable) FieldType() FieldType {
	return StringFieldType
}

func (s StringVariable) Value() interface{} {
	return s.val
}

func (s StringVariable) GetString() string {
	return s.val
}

type BoolVariable struct {
	val bool
}

func (b BoolVariable) IsNull() bool {
	return false
}

func (b BoolVariable) FieldType() FieldType {
	return BoolFieldType
}

func (b BoolVariable) Value() interface{} {
	return b.val
}

func (b BoolVariable) GetBool() bool {
	return b.val
}

type NumberVariable struct {
	val decimal.Decimal
}

func (n NumberVariable) FieldType() FieldType {
	return NumberFieldType
}

func (n NumberVariable) Value() interface{} {
	return n.val
}

func (n NumberVariable) IsNull() bool {
	return false
}

func (n NumberVariable) GetDecimal() decimal.Decimal {
	return n.val
}

type ErrorVariable struct {
	val error
}

func (e ErrorVariable) FieldType() FieldType {
	return ErrorFieldType
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

type FunctionVariable struct {
	fn               interface{}
	argumentTypeList []FieldType
	returnTypeList   []FieldType
}

func (f FunctionVariable) FieldType() FieldType {
	return FunctionFieldType
}

func (f FunctionVariable) Value() interface{} {
	return f.fn
}

func (f FunctionVariable) IsNull() bool {
	return false
}

func NewVariableFromInterface(val interface{}) (Variable, error) {
	switch v := val.(type) {
	case string:
		return StringVariable{
			val: v,
		}, nil
	case bool:
		return BoolVariable{
			val: v,
		}, nil
	case int, int64, int32, int16, int8, uint, uint8, uint16, uint32, uint64:
		return NumberVariable{
			val: decimal.NewFromInt(val.(int64)),
		}, nil
	case float32, float64:
		return NumberVariable{
			val: decimal.NewFromFloat(val.(float64)),
		}, nil
	default:
		t := reflect.TypeOf(val)
		if t.Kind() != reflect.Func {
			return NullVariableInstance, NotSupportTypeErr
		}
		var argumentTypeList []FieldType
		for i := 0; i < t.NumIn(); i++ {
			fieldType, err := GetFieldTypeFromReflect(t.In(i))
			if err != nil {
				return NullVariableInstance, err
			}
			if fieldType == ErrorFieldType {
				return NullVariableInstance, errors.New("not support type")
			}
			argumentTypeList = append(argumentTypeList, fieldType)
		}
		var returnTypeList []FieldType
		for i := 0; i < t.NumOut(); i++ {
			fieldType, err := GetFieldTypeFromReflect(t.Out(i))
			if err != nil {
				return NullVariableInstance, err
			}
			returnTypeList = append(returnTypeList, fieldType)
		}
		return FunctionVariable{
			fn:               v,
			argumentTypeList: argumentTypeList,
			returnTypeList:   returnTypeList,
		}, nil
	}
}

func NewVariableFromString(fieldType FieldType, val string) (Variable, error) {
	switch fieldType {
	case StringFieldType:
		return StringVariable{
			val: val,
		}, nil
	case BoolFieldType:
		if val == "true" {
			return BoolVariable{
				val: true,
			}, nil
		}
		if val == "false" {
			return BoolVariable{
				val: false,
			}, nil
		}
		return NullVariableInstance, nil
	case NumberFieldType:
		number, err := decimal.NewFromString(val)
		if err != nil {
			return NullVariableInstance, err
		}
		return NumberVariable{
			val: number,
		}, err
	default:
		return NullVariableInstance, NotSupportTypeErr
	}
}

type Function struct {
	fn               interface{}
	argumentTypeList []FieldType
	returnTypeList   []FieldType
}

type Variables map[string]Variable

func (v Variables) Put(name string, val interface{}) error {
	field, err := NewVariableFromInterface(val)
	if err != nil {
		return err
	}
	v[name] = field
	return nil
}

func (v Variables) Get(name string) Variable {
	return v[name]
}

type VariableStack struct {
	stack []Variables
}

func (s VariableStack) SetVariable(name string, val interface{}) error {
	for i := len(s.stack) - 1; i >= 0; i++ {
		if s.stack[i].Get(name) != NullVariableInstance {
			return s.stack[i].Put(name, val)
		}
	}
	return s.stack[len(s.stack)-1].Put(name, val)
}

func (s VariableStack) GetVariable(name string) Variable {
	for i := len(s.stack) - 1; i >= 0; i++ {
		field := s.stack[i].Get(name)
		if field != NullVariableInstance {
			return field
		}
	}
	return NullVariableInstance
}

func (s VariableStack) PushVariables() VariableStack {
	s.stack = append(s.stack, map[string]Variable{})
	return s
}

func (s VariableStack) PopVariables() VariableStack {
	s.stack = s.stack[:len(s.stack)-1]
	return s
}
