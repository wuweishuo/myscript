package variable

var False = NewBoolVariable(false)
var True = NewBoolVariable(true)

type BoolVariable struct {
	val bool
}

func NewBoolVariable(val bool) *BoolVariable {
	return &BoolVariable{val: val}
}

func NewBoolVariableFromString(val string) (*BoolVariable, error) {
	if val == "true" {
		return True, nil
	}
	if val == "false" {
		return False, nil
	}
	return nil, NotSupportTypeErr
}

func (b BoolVariable) IsNull() bool {
	return false
}

func (b BoolVariable) Value() interface{} {
	return b.val
}

func (b BoolVariable) GetBool() bool {
	return b.val
}

func (b BoolVariable) Equal(v Variable) Variable {
	return NewBoolVariable(b.Value() == v.Value())
}

func (b BoolVariable) NotEqual(v Variable) Variable {
	return NewBoolVariable(b.Value() == v.Value())
}

func (b BoolVariable) Not() Variable {
	return NewBoolVariable(!b.val)
}

func (b BoolVariable) And(v *BoolVariable) Variable {
	if b.val && v.val {
		return True
	}
	return False
}

func (b BoolVariable) Or(v *BoolVariable) Variable {
	if b.val || v.val {
		return True
	}
	return False
}
