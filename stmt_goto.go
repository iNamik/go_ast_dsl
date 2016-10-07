package go_ast_dsl

import (
	"go/token"
	"go/ast"
)

// GotoStmt
//
type GotoStmt struct {
	label string
}

func NewGotoStmt() *GotoStmt {
	return &GotoStmt{}
}

func (s *GotoStmt) WithLabel(label string) *GotoStmt {
	s.label = label
	return s
}

func (s *GotoStmt) BranchStmt() *ast.BranchStmt {
	stmt := &ast.BranchStmt{Tok: token.GOTO}
	stmt.Label = &ast.Ident{Name: s.label}
	return stmt
}

func (s *GotoStmt) Stmt() ast.Stmt {
	return s.BranchStmt()
}
