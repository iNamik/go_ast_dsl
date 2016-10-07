package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

// Var
//
type Var struct {
	name  string
	typ   TypeExprProvider
	value ExprProvider
}

func NewVar() *Var {
	return &Var{}
}

func (v *Var) WithName(name string) *Var {
	v.name = name
	return v
}

func (v *Var) WithType(typ TypeExprProvider) *Var {
	v.typ = typ
	return v
}

func (v *Var) WithValue(value ExprProvider) *Var {
	v.value = value
	return v
}

func (v *Var) ValueSpec() *ast.ValueSpec {
	spec := &ast.ValueSpec{}
	if nil != v.typ {
		spec.Type = v.typ.TypeExpr()
	}
	AppendToValueSpec(spec, v.Ident(), v.Expr())
	return spec
}

func (v *Var) Spec() ast.Spec {
	return v.ValueSpec()
}

func (v *Var) Ident() *ast.Ident {
	return &ast.Ident{Name: v.name}
}

func (v *Var) Expr() ast.Expr {
	if nil == v.value {
		return nil
	}
	return v.value.Expr()
}

func (v *Var) GenDecl() *ast.GenDecl {
	decl := newVarDecl()
	AppendToDecl(decl, v)
	return decl
}

func (v *Var) Decl() ast.Decl {
	return v.GenDecl()
}

// VarList
//
type VarList struct {
	list []*Var
	typ  TypeExprProvider
}

func NewVarList() *VarList {
	return &VarList{}
}

func (l *VarList) WithType(typ TypeExprProvider) *VarList {
	l.typ = typ
	return l
}

func (l *VarList) Add(v *Var) *VarList {
	l.list = append(l.list, v)
	return l
}

func (l *VarList) ValueSpec() *ast.ValueSpec {
	spec := &ast.ValueSpec{}
	if nil != l.typ {
		spec.Type = l.typ.TypeExpr()
	}
	for _, v := range l.list {
		AppendToValueSpec(spec, v.Ident(), v.Expr())
	}
	return spec
}

func (l *VarList) Spec() ast.Spec {
	return l.ValueSpec()
}

func (l *VarList) GenDecl() *ast.GenDecl {
	decl := newVarDecl()
	AppendToDecl(decl, l)
	return decl
}

func (l *VarList) Decl() ast.Decl {
	return l.GenDecl()
}

// Var Group
//
type VarGroup struct {
	list []SpecProvider
}

func NewVarGroup() *VarGroup {
	return &VarGroup{}
}

func (g *VarGroup) AddVar(v *Var) *VarGroup {
	g.list = append(g.list, v)
	return g
}

func (g *VarGroup) AddList(l *VarList) *VarGroup {
	g.list = append(g.list, l)
	return g
}

func (g *VarGroup) GenDecl() *ast.GenDecl {
	decl := newVarDeclGroup()
	AppendToDecl(decl, g.list...)
	return decl
}

func (g *VarGroup) Decl() ast.Decl {
	return g.GenDecl()
}

func newVarDecl() *ast.GenDecl {
	return NewGenDecl(token.VAR)
}

func newVarDeclGroup() *ast.GenDecl {
	return NewGenDeclGroup(token.VAR)
}
