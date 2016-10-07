package go_ast_dsl

import "go/ast"

// TypeRef
//
type TypeRef struct {
	name string
	pkg  string
}

func NewTypeRef() *TypeRef {
	return &TypeRef{}
}

func (t *TypeRef) WithName(name string) *TypeRef {
	t.name = name
	return t
}

func (t *TypeRef) WithPackage(pkg string) *TypeRef {
	t.pkg = pkg
	return t
}

func (t *TypeRef) TypeExpr() ast.Expr {
	name := &ast.Ident{Name: t.name}
	if len(t.pkg) > 0 {
		pkg := &ast.Ident{Name: t.pkg}
		return &ast.SelectorExpr{X: pkg, Sel: name}
	}
	return name
}

func (t *TypeRef) Expr() ast.Expr {
	return t.TypeExpr()
}
