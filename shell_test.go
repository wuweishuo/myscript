package myscript

import (
	"context"
	"fmt"
	"myscript/variable"
	"testing"
)

func Test_shell(t *testing.T) {
	shell := NewShell()
	script, err := shell.Compile(`
a = add(1, 2)
	`)
	if err != nil {
		panic(err)
	}
	err = script.PutVariable("add", func(i ...variable.Variable) (variable.Variable, error) {
		res := variable.NewNumberVariableFromInt(0)
		for _, v := range i {
			res = res.Add(v.(*variable.NumberVariable))
		}
		return res, nil
	})
	if err != nil {
		panic(err)
	}
	err = script.PutVariable("a", 1)
	if err != nil {
		panic(err)
	}
	err = script.Eval(context.TODO())
	if err != nil {
		panic(err)
	}
	result := script.GetResult("a")
	fmt.Println(result.Value())
}
