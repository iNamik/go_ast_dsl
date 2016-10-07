package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

// TypeDef
//
type TypeDef struct {
	name string
	typ  TypeExprProvider
}

func NewTypeDef() *TypeDef {
	return &TypeDef{}
}

func (t *TypeDef) WithName(name string) *TypeDef {
	t.name = name
	return t
}

func (t *TypeDef) WithType(typ TypeExprProvider) *TypeDef {
	t.typ = typ
	return t
}

func (t *TypeDef) TypeSpec() *ast.TypeSpec {
	ident := &ast.Ident{Name: t.name}
	spec := &ast.TypeSpec{Name: ident, Type: t.typ.TypeExpr()}
	return spec
}

func (t *TypeDef) Spec() ast.Spec {
	return t.TypeSpec()
}

func (t *TypeDef) GenDecl() *ast.GenDecl {
	decl := newTypeDecl()
	AppendToDecl(decl, t)
	return decl
}

func (t *TypeDef) Decl() ast.Decl {
	return t.GenDecl()
}

func newTypeDecl() *ast.GenDecl {
	return NewGenDecl(token.TYPE)
}

func newTypeDeclGroup() *ast.GenDecl {
	return NewGenDeclGroup(token.TYPE)
}

// TypeDefGroup
//
type TypeDefGroup struct {
	list []SpecProvider
}

func NewTypeDefGroup() *TypeDefGroup {
	return &TypeDefGroup{}
}

func (g *TypeDefGroup) AddType(t *TypeDef) *TypeDefGroup {
	g.list = append(g.list, t)
	return g
}

func (g *TypeDefGroup) GenDecl() *ast.GenDecl {
	decl := newTypeDeclGroup()
	AppendToDecl(decl, g.list...)
	return decl
}

func (g *TypeDefGroup) Decl() ast.Decl {
	return g.GenDecl()
}
