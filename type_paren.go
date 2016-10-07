package go_ast_dsl

import "go/ast"

// ParenType
//
type ParenType struct {
	typ TypeExprProvider
}

func NewParenType() *ParenType {
	return &ParenType{}
}

func (t *ParenType) WithType(typ TypeExprProvider) *ParenType {
	t.typ = typ
	return t
}

func (t *ParenType) TypeExpr() ast.Expr {
	// &ast.ParenExpr{Lparen: lparen, X: typ, Rparen: rparen}
	return &ast.ParenExpr{X: t.typ.TypeExpr()}
}
