package go_ast_dsl

import "go/ast"

// MapType
//
type MapType struct {
	key TypeExprProvider
	val TypeExprProvider
}

func NewMapType() *MapType {
	return &MapType{}
}

func (t *MapType) WithKeyType(typ TypeExprProvider) *MapType {
	t.key = typ
	return t
}

func (t *MapType) WithValueType(typ TypeExprProvider) *MapType {
	t.val = typ
	return t
}

func (t *MapType) TypeExpr() ast.Expr {
	// &ast.MapType{Map: pos, Key: key, Value: value}
	return &ast.MapType{Key: t.key.TypeExpr(), Value: t.val.TypeExpr()}
}
