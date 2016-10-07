package go_ast_dsl

import "go/ast"

// ConstStmt
//
type ConstStmt struct {
	decl DeclProvider
}

func NewConstStmt() *ConstStmt {
	return &ConstStmt{}
}

func (s *ConstStmt) WithConst(decl *ConstDecl) *ConstStmt {
	s.decl = decl
	return s
}

func (s *ConstStmt) WithConstList(decl *ConstDeclList) *ConstStmt {
	s.decl = decl
	return s
}

func (s *ConstStmt) WithConstGroup(decl *ConstDeclGroup) *ConstStmt {
	s.decl = decl
	return s
}

func (s *ConstStmt) Stmt() ast.Stmt {
	stmt := &ast.DeclStmt{Decl: s.decl.Decl()}
	return stmt
}
