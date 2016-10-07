package go_ast_dsl

import "go/ast"

// TypeStmt
//
type TypeStmt struct {
	decl DeclProvider
}

func NewTypeStmt() *TypeStmt {
	return &TypeStmt{}
}

func (s *TypeStmt) WithType(decl *TypeDef) *TypeStmt {
	s.decl = decl
	return s
}

func (s *TypeStmt) WithTypeGroup(decl *TypeDefGroup) *TypeStmt {
	s.decl = decl
	return s
}
func (s *TypeStmt) Stmt() ast.Stmt {
	stmt := &ast.DeclStmt{Decl: s.decl.Decl()}
	return stmt
}
