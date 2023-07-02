package jlib

import (
	"log"

	"github.com/ZombiiTheCoder/JsonParser/lexer"
	"github.com/ZombiiTheCoder/JsonParser/mapper"
	"github.com/ZombiiTheCoder/JsonParser/parser"
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

func Access(j map[string]any, text string) map[string]any {

	if j[text] == nil {
		log.Fatalf("Value \"%v\" Does Not Exist In Map Or Value in map is not a nested map[string]any\n", text)
	}

	return j[text].(map[string]any)

}