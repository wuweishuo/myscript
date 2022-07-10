package variable

type Variables map[string]Variable

func NewVariables() Variables {
	return make(Variables)
}

func (v Variables) Put(name string, val Variable) {
	v[name] = val
}

func (v Variables) Get(name string) Variable {
	return v[name]
}

type Stack struct {
	stack []Variables
}

func NewStack() Stack {
	return Stack{
		stack: []Variables{NewVariables()},
	}
}

func (s Stack) SetVariable(name string, val Variable) {
	for i := len(s.stack) - 1; i >= 0; i-- {
		if s.stack[i].Get(name) != nil {
			s.stack[i].Put(name, val)
			return
		}
	}
	s.stack[len(s.stack)-1].Put(name, val)
}

func (s Stack) GetVariable(name string) Variable {
	for i := len(s.stack) - 1; i >= 0; i-- {
		field := s.stack[i].Get(name)
		if field != NullVariableInstance {
			return field
		}
	}
	return NullVariableInstance
}

func (s Stack) PushVariables() Stack {
	s.stack = append(s.stack, NewVariables())
	return s
}

func (s Stack) PopVariables() Stack {
	s.stack = s.stack[:len(s.stack)-1]
	return s
}
