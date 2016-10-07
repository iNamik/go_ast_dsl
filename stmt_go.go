package go_ast_dsl

import (
	"go/ast"
)

// GoStmt
//
type GoStmt struct {
	expr CallExprProvider
}

func NewGoStmt() *GoStmt {
	return &GoStmt{}
}

func (s *GoStmt) WithCallExpr(exp CallExprProvider) *GoStmt {
	s.expr = exp
	return s
}

func (s *GoStmt) Stmt() ast.Stmt {
	stmt := &ast.GoStmt{Call: s.expr.CallExpr()}
	return stmt
}
