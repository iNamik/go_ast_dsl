package go_ast_dsl

import "go/ast"

// ExprStmt
//
type ExprStmt struct {
	expr ExprProvider
}

func NewExprStmt() *ExprStmt {
	return &ExprStmt{}
}

func (s *ExprStmt) WithExpr(expr ExprProvider) *ExprStmt {
	s.expr = expr
	return s
}

func (s *ExprStmt) Stmt() ast.Stmt {
	stmt := &ast.ExprStmt{X: s.expr.Expr()}
	return stmt
}
