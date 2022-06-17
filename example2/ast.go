package main

type Expr interface {

}

type Constant struct {
	Value int
}

type Variable struct {
	Name string
}

type Assignment struct {
	Name string
	Expr Expr
}

type Binary struct {
	Op    string
	Left  Expr
	Right Expr
}
