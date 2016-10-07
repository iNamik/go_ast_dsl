package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

type Literal struct {
	kind  token.Token
	value string
}

func NewIntLiteral(value string) *Literal {
	return &Literal{kind: token.INT, value: value}
}

func NewFloatLiteral(value string) *Literal {
	return &Literal{kind: token.FLOAT, value: value}
}

func NewImagLiteral(value string) *Literal {
	return &Literal{kind: token.IMAG, value: value}
}

func NewCharLiteral(value string) *Literal {
	return &Literal{kind: token.CHAR, value: value}
}

func NewStringLiteral(value string) *Literal {
	return &Literal{kind: token.STRING, value: value}
}

func (l *Literal) BasicLit() *ast.BasicLit {
	return &ast.BasicLit{Kind: l.kind, Value: l.value}
}

func (l *Literal) Expr() ast.Expr {
	return l.BasicLit()
}
