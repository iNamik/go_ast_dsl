package go_ast_dsl

import (
	"go/token"
	"go/ast"
)

// BreakStmt
//
type BreakStmt struct {
	label string
}

func NewBreakStmt() *BreakStmt {
	return &BreakStmt{}
}

func (s *BreakStmt) WithLabel(label string) *BreakStmt {
	s.label = label
	return s
}

func (s *BreakStmt) BranchStmt() *ast.BranchStmt {
	stmt := &ast.BranchStmt{Tok: token.BREAK}
	if len(s.label) > 0 {
		stmt.Label = &ast.Ident{Name: s.label}
	}
	return stmt
}

func (s *BreakStmt) Stmt() ast.Stmt {
	return s.BranchStmt()
}
