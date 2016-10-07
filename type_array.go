package go_ast_dsl

import "go/ast"

// ArrayType
//
type ArrayType struct {
	len ExprProvider
	typ TypeExprProvider
}

func NewArrayType() *ArrayType {
	return &ArrayType{}
}

func (t *ArrayType) WithLen(len ExprProvider) *ArrayType {
	t.len = len
	return t
}

func (t *ArrayType) WithType(typ TypeExprProvider) *ArrayType {
	t.typ = typ
	return t
}

func (t *ArrayType) TypeExpr() ast.Expr {
	// &ast.ArrayType{Lbrack: lbrack, Len: len, Elt: elt}
	var len ast.Expr
	if nil != t.len {
		len = t.len.Expr()
	}
	return &ast.ArrayType{Len: len, Elt: t.typ.TypeExpr()}
}
