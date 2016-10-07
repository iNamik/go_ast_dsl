package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

// IncStmt
//
type IncStmt struct {
	expr ExprProvider
}

func NewIncStmt() *IncStmt {
	return &IncStmt{}
}

func (s *IncStmt) WithExpr(exp ExprProvider) *IncStmt {
	s.expr = exp
	return s
}

func (s *IncStmt) Stmt() ast.Stmt {
	stmt := &ast.IncDecStmt{Tok: token.INC, X: s.expr.Expr()}
	return stmt
}
