package myscript

import (
	"context"
	"fmt"
	"testing"
)

func Test_shell(t *testing.T) {
	shell := NewShell()
	err := shell.PutFunction("add", func(i, j int) int {
		return i + j
	})
	if err != nil {
		panic(err)
	}
	script, err := shell.Compile("a = add(1, 2)")
	if err != nil {
		panic(err)
	}
	err = script.Eval(context.TODO())
	if err != nil {
		panic(err)
	}
	result := script.GetResult("a")
	fmt.Println(result)
}
