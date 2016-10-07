package go_ast_dsl

import (
	"go/token"
	"go/ast"
)

// ContinueStmt
//
type ContinueStmt struct {
	label string
}

func NewContinueStmt() *ContinueStmt {
	return &ContinueStmt{}
}

func (s *ContinueStmt) WithLabel(label string) *ContinueStmt {
	s.label = label
	return s
}

func (s *ContinueStmt) BranchStmt() *ast.BranchStmt {
	stmt := &ast.BranchStmt{Tok: token.CONTINUE}
	if len(s.label) > 0 {
		stmt.Label = &ast.Ident{Name: s.label}
	}
	return stmt
}

func (s *ContinueStmt) Stmt() ast.Stmt {
	return s.BranchStmt()
}
