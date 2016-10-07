package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

func NewGenDecl(tok token.Token) *ast.GenDecl {
	return &ast.GenDecl{Tok: tok}
}

func NewGenDeclGroup(tok token.Token) *ast.GenDecl {
	return &ast.GenDecl{Tok: tok, Lparen: 1} // 1 Forces the parens in the output
}

func AppendToValueSpec(spec *ast.ValueSpec, ident *ast.Ident, expr ast.Expr) *ast.ValueSpec {
	spec.Names = append(spec.Names, ident)
	if nil != expr {
		spec.Values = append(spec.Values, expr)
	}
	return spec
}

func AppendToDecl(decl *ast.GenDecl, list ...SpecProvider) {
	for _, p := range list {
		decl.Specs = append(decl.Specs, p.Spec())
	}
}

type IdentProvider interface {
	Ident() *ast.Ident
}

type ValueSpecProvider interface {
	ValueSpec() *ast.ValueSpec
}

type TypeSpecProvider interface {
	TypeSpec() *ast.TypeSpec
}

type SpecProvider interface {
	Spec() ast.Spec
}

type GenDeclProvider interface {
	GenDecl() *ast.GenDecl
}

type DeclProvider interface {
	Decl() ast.Decl
}

type BasicLitProvider interface {
	BasicLit() *ast.BasicLit
}

type TypeExprProvider interface {
	TypeExpr() ast.Expr
}

type ExprProvider interface {
	Expr() ast.Expr
}

type FieldDeclProvider interface {
	FieldDecl() *ast.Field
}

type FieldListProvider interface {
	FieldList() *ast.FieldList
}

type StmtProvider interface {
	Stmt() ast.Stmt
}

type BlockStmtProvider interface {
	BlockStmt() *ast.BlockStmt
}

