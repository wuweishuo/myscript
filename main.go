package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "myscript/gen"
)

func main() {
	stream := antlr.NewInputStream("!(2<3)")
	lexer := parser.NewMyScriptLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMyScriptParser(tokenStream)
	c := &Listener{}
	antlr.ParseTreeWalkerDefault.Walk(c, p.Statement())
}
