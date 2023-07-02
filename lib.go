package jsonparser

import (
	"jsonparser/lexer"
	"jsonparser/mapper"
	"jsonparser/parser"
)

func MapJson(text string) map[string]any {

	lex := lexer.Lexer{}
	lex.Init(text)
	Tokens := lex.Tokenize()

	par := parser.Parser{}
	par.Init(Tokens)
	Ast := par.Parse()

	mapp := mapper.Mapper{}
	return mapp.Eval(Ast).(map[string]any)

}