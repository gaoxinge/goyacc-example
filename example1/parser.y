%{
package main

import (
    "fmt"
)
%}

%union {
    Int int
    Str string
}
%token <Int> INTEGER
%token <Str> VARIABLE
%type  <Int> expr
%left '+' '-'
%left '*' '/'

%%

program:
        program statement '\n'
        |
        ;

statement:
        expr                    { fmt.Println($1); }
        | VARIABLE '=' expr     { yylex.(*Lexer).Sym[$1] = $3; }
        ;

expr:
        INTEGER
        | VARIABLE              { $$ = yylex.(*Lexer).Sym[$1]; }
        | expr '+' expr         { $$ = $1 + $3; }
        | expr '-' expr         { $$ = $1 - $3; }
        | expr '*' expr         { $$ = $1 * $3; }
        | expr '/' expr         { $$ = $1 / $3; }
        | '(' expr ')'          { $$ = $2; }
        ;

%%