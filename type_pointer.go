package go_ast_dsl

import "go/ast"

// PointerType
//
type PointerType struct {
	typ TypeExprProvider
}

func NewPointerType() *PointerType {
	return &PointerType{}
}

func (t *PointerType) WithType(typ TypeExprProvider) *PointerType {
	t.typ = typ
	return t
}

func (t *PointerType) TypeExpr() ast.Expr {
	// &ast.StarExpr{Star: star, X: base}
	return &ast.StarExpr{X: t.typ.TypeExpr()}
}
