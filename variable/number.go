package variable

import "github.com/shopspring/decimal"

type NumberVariable struct {
	val decimal.Decimal
}

func NewNumberVariableFromString(val string) (*NumberVariable, error) {
	number, err := decimal.NewFromString(val)
	if err != nil {
		return nil, err
	}
	return NewNumberVariable(number), nil
}

func NewNumberVariableFromInt(val int64) *NumberVariable {
	return NewNumberVariable(decimal.NewFromInt(val))
}

func NewNumberVariableFromFloat(val float64) *NumberVariable {
	return NewNumberVariable(decimal.NewFromFloat(val))
}

func NewNumberVariable(val decimal.Decimal) *NumberVariable {
	return &NumberVariable{val: val}
}

func (n *NumberVariable) Value() interface{} {
	return n.val
}

func (n *NumberVariable) IsNull() bool {
	return false
}

func (n *NumberVariable) GetDecimal() decimal.Decimal {
	return n.val
}

func (n *NumberVariable) Mul(other *NumberVariable) *NumberVariable {
	return NewNumberVariable(n.val.Mul(other.val))
}

func (n *NumberVariable) Div(other *NumberVariable) *NumberVariable {
	return NewNumberVariable(n.val.Div(other.val))
}

func (n *NumberVariable) Mod(other *NumberVariable) *NumberVariable {
	return NewNumberVariable(n.val.Mod(other.val))
}

func (n *NumberVariable) Add(other *NumberVariable) *NumberVariable {
	return NewNumberVariable(n.val.Add(other.val))
}

func (n *NumberVariable) Sub(other *NumberVariable) *NumberVariable {
	return NewNumberVariable(n.val.Sub(other.val))
}

func (n *NumberVariable) Greater(other *NumberVariable) Variable {
	return NewBoolVariable(n.val.GreaterThan(other.val))
}

func (n *NumberVariable) GreaterOrEqual(other *NumberVariable) Variable {
	return NewBoolVariable(n.val.GreaterThanOrEqual(other.val))
}

func (n *NumberVariable) Less(other *NumberVariable) Variable {
	return NewBoolVariable(n.val.LessThan(other.val))
}

func (n *NumberVariable) LessOrEqual(other *NumberVariable) Variable {
	return NewBoolVariable(n.val.LessThanOrEqual(other.val))
}

func (n NumberVariable) Equal(v Variable) Variable {
	if n.val.Equal(v.(*NumberVariable).GetDecimal()) {
		return True
	}
	return False
}

func (n NumberVariable) NotEqual(v Variable) Variable {
	if n.val.Equal(v.(*NumberVariable).GetDecimal()) {
		return False
	}
	return True
}
