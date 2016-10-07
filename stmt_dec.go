package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

// DecStmt
//
type DecStmt struct {
	expr ExprProvider
}

func NewDecStmt() *DecStmt {
	return &DecStmt{}
}

func (s *DecStmt) WithExpr(exp ExprProvider) *DecStmt {
	s.expr = exp
	return s
}

func (s *DecStmt) Stmt() ast.Stmt {
	stmt := &ast.IncDecStmt{Tok: token.DEC, X: s.expr.Expr()}
	return stmt
}
