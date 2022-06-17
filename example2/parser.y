%{
package main
%}

%union {
    Int  int
    Str  string
    Expr Expr
}
%token <Int>  INTEGER
%token <Str>  VARIABLE
%type  <Expr> statement expr
%left '+' '-'
%left '*' '/'

%%

program:
        program statement '\n'
        |
        ;

statement:
        expr                    { yylex.(*Lexer).Statements = append(yylex.(*Lexer).Statements, $1); }
        | VARIABLE '=' expr     { yylex.(*Lexer).Statements = append(yylex.(*Lexer).Statements, &Assignment{$1, $3}); }
        ;

expr:
        INTEGER                 { $$ = &Constant{$1}; }
        | VARIABLE              { $$ = &Variable{$1}; }
        | expr '+' expr         { $$ = &Binary{"+", $1, $3}; }
        | expr '-' expr         { $$ = &Binary{"-", $1, $3}; }
        | expr '*' expr         { $$ = &Binary{"*", $1, $3}; }
        | expr '/' expr         { $$ = &Binary{"/", $1, $3}; }
        | '(' expr ')'          { $$ = $2; }
        ;

%%