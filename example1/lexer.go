package main

import (
	"regexp"
	"strconv"
)

type Lexer struct {
	input string
	Sym   map[string]int
}

func NexLexer(program string) *Lexer {
	lexer := &Lexer{
		input: program,
		Sym:   make(map[string]int),
	}
	return lexer
}

func (lexer *Lexer) Lex(lval *yySymType) int {
	for ; len(lexer.input) > 0 && (lexer.input[0] == ' ' || lexer.input[0] == '\t'); lexer.input = lexer.input[1:] {
		// pass
	}

	if len(lexer.input) == 0 {
		return 0
	}

	switch lexer.input[0] {
	case '+':
		lexer.input = lexer.input[1:]
		return '+'
	case '-':
		lexer.input = lexer.input[1:]
		return '-'
	case '*':
		lexer.input = lexer.input[1:]
		return '*'
	case '/':
		lexer.input = lexer.input[1:]
		return '/'
	case '=':
		lexer.input = lexer.input[1:]
		return '='
	case '\n':
		lexer.input = lexer.input[1:]
		return '\n'
	}

	var v string

	v = regexp.MustCompile(`^[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?`).FindString(lexer.input)
	if len(v) > 0 {
		lexer.input = lexer.input[len(v):]
		lval.Int, _ = strconv.Atoi(v)
		return INTEGER
	}

	v = regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9]*`).FindString(lexer.input)
	if len(v) > 0 {
		lexer.input = lexer.input[len(v):]
		lval.Str = v
		return VARIABLE
	}

	return 0
}

func (lexer *Lexer) Error(e string) {
	panic(e)
}
