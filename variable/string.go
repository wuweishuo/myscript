package variable

type StringVariable struct {
	val string
}

func NewStringVariable(val string) *StringVariable {
	return &StringVariable{val: val}
}

func (s StringVariable) IsNull() bool {
	return false
}

func (s StringVariable) Value() interface{} {
	return s.val
}

func (s StringVariable) GetString() string {
	return s.val
}

func (s StringVariable) Equal(v Variable) Variable {
	if s.val == v.Value() {
		return True
	}
	return False
}

func (s StringVariable) NotEqual(v Variable) Variable {
	if s.val == v.Value() {
		return False
	}
	return True
}
