package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

// ConstDecl
// see:
//   parser/parseGenDecl
//   parser/parseValueSpec
//
type ConstDecl struct {
	name  string
	typ   TypeExprProvider
	value ExprProvider
	iota  int
}

func NewConstDecl() *ConstDecl { return &ConstDecl{} }

func (c *ConstDecl) WithName(name string) *ConstDecl {
	c.name = name
	return c
}

func (c *ConstDecl) WithType(typ TypeExprProvider) *ConstDecl {
	c.typ = typ
	return c
}

func (c *ConstDecl) WithValue(value ExprProvider) *ConstDecl {
	c.value = value
	return c
}

func (c *ConstDecl) WithIota(iota int) *ConstDecl {
	c.iota = iota
	return c
}

func (c *ConstDecl) ValueSpec() *ast.ValueSpec {
	spec := &ast.ValueSpec{}
	if nil != c.typ {
		spec.Type = c.typ.TypeExpr()
	}
	AppendToValueSpec(spec, c.Ident(), c.Expr())
	return spec
}

func (c *ConstDecl) Spec() ast.Spec {
	return c.ValueSpec()
}

func (c *ConstDecl) Ident() *ast.Ident {
	return &ast.Ident{Name: c.name}
}

func (c *ConstDecl) Expr() ast.Expr {
	if nil == c.value {
		return nil
	}
	return c.value.Expr()
}

func (c *ConstDecl) GenDecl() *ast.GenDecl {
	decl := newConstDecl()
	AppendToDecl(decl, c)
	return decl
}

func (c *ConstDecl) Decl() ast.Decl {
	return c.GenDecl()
}

// ConstDeclList
//
type ConstDeclList struct {
	list []*ConstDecl
	typ  TypeExprProvider
}

func NewConstDeclList() *ConstDeclList {
	return &ConstDeclList{}
}

func (l *ConstDeclList) WithType(typ TypeExprProvider) *ConstDeclList {
	l.typ = typ
	return l
}

func (l *ConstDeclList) Add(c *ConstDecl) *ConstDeclList {
	l.list = append(l.list, c)
	return l
}

func (l *ConstDeclList) ValueSpec() *ast.ValueSpec {
	spec := &ast.ValueSpec{}
	if nil != l.typ {
		spec.Type = l.typ.TypeExpr()
	}
	for _, c := range l.list {
		AppendToValueSpec(spec, c.Ident(), c.Expr())
	}
	return spec
}

func (l *ConstDeclList) Spec() ast.Spec {
	return l.ValueSpec()
}

func (l *ConstDeclList) GenDecl() *ast.GenDecl {
	decl := newConstDecl()
	AppendToDecl(decl, l)
	return decl
}

func (l *ConstDeclList) Decl() ast.Decl {
	return l.GenDecl()
}

// ConstDeclGroup
//
type ConstDeclGroup struct {
	list []SpecProvider
}

func NewConstDeclGroup() *ConstDeclGroup {
	return &ConstDeclGroup{}
}

func (g *ConstDeclGroup) AddConst(c *ConstDecl) *ConstDeclGroup {
	g.list = append(g.list, c)
	return g
}

func (g *ConstDeclGroup) AddList(l *ConstDeclList) *ConstDeclGroup {
	g.list = append(g.list, l)
	return g
}

func (g *ConstDeclGroup) GenDecl() *ast.GenDecl {
	decl := newConstDeclGroup()
	AppendToDecl(decl, g.list...)
	return decl
}

func (g *ConstDeclGroup) Decl() ast.Decl {
	return g.GenDecl()
}

func newConstDecl() *ast.GenDecl {
	return NewGenDecl(token.CONST)
}

func newConstDeclGroup() *ast.GenDecl {
	return NewGenDeclGroup(token.CONST)
}
