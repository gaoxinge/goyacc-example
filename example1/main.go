package main

func main() {
	program := "a = 1 + 1\nb = 4 - 2\nb * a\nb / a\n"
	lexer := NexLexer(program)
	yyParse(lexer)
}
