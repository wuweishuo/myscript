package variable

type NullVariable struct{}

var NullVariableInstance = NullVariable{}

func (n NullVariable) IsNull() bool {
	return true
}

func (n NullVariable) Value() interface{} {
	return nil
}

func (n NullVariable) Equal(v Variable) Variable {
	if v.IsNull() {
		return True
	}
	return False
}

func (n NullVariable) NotEqual(v Variable) Variable {
	if v.IsNull() {
		return False
	}
	return True
}
