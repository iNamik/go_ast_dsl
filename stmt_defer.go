package go_ast_dsl

import (
	"go/ast"
)

// DeferStmt
//
type DeferStmt struct {
	expr CallExprProvider
}

func NewDeferStmt() *DeferStmt {
	return &DeferStmt{}
}

func (s *DeferStmt) WithCallExpr(exp CallExprProvider) *DeferStmt {
	s.expr = exp
	return s
}

func (s *DeferStmt) Stmt() ast.Stmt {
	stmt := &ast.DeferStmt{Call: s.expr.CallExpr()}
	return stmt
}
