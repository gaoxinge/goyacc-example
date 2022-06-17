package main

import "fmt"

func evalExpr(expr Expr, sym map[string]int) int {
	switch e := expr.(type) {
	case *Constant:
		return e.Value
	case *Variable:
		return sym[e.Name]
	case *Assignment:
		value := evalExpr(e.Expr, sym)
		sym[e.Name] = value
		return -1
	case *Binary:
		lhs := evalExpr(e.Left, sym)
		rhs := evalExpr(e.Right, sym)
		switch e.Op {
		case "+":
			return lhs + rhs
		case "-":
			return lhs - rhs
		case "*":
			return lhs * rhs
		case "/":
			return lhs / rhs
		}
	}
	return -1
}

func eval(statements []Expr) {
	sym := make(map[string]int)
	for _, statement := range statements {
		value := evalExpr(statement, sym)
		if value != -1 {
			fmt.Println(value)
		}
	}
}

func main() {
	program := "a = 1 + 1\nb = 4 - 2\nb * a\nb / a\n"
	lexer := NexLexer(program)
	yyParse(lexer)
	eval(lexer.Statements)
}
