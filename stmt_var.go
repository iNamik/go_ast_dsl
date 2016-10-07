package go_ast_dsl

import "go/ast"

// VarStmt
//
type VarStmt struct {
	decl DeclProvider
}

func NewVarStmt() *VarStmt {
	return &VarStmt{}
}

func (s *VarStmt) WithVar(decl *Var) *VarStmt {
	s.decl = decl
	return s
}

func (s *VarStmt) WithVarList(decl *VarList) *VarStmt {
	s.decl = decl
	return s
}

func (s *VarStmt) WithVarGroup(decl *VarGroup) *VarStmt {
	s.decl = decl
	return s
}

func (s *VarStmt) Stmt() ast.Stmt {
	stmt := &ast.DeclStmt{Decl: s.decl.Decl()}
	return stmt
}
