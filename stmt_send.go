package go_ast_dsl

import "go/ast"

// SendStmt
//
type SendStmt struct {
	chn ExprProvider
	val ExprProvider
}

func NewSendStmt() *SendStmt {
	return &SendStmt{}
}

func (s *SendStmt) WithChanExpr(expr ExprProvider) *SendStmt {
	s.chn = expr
	return s
}

func (s *SendStmt) WithValExpr(expr ExprProvider) *SendStmt {
	s.val = expr
	return s
}

func (s *SendStmt) Stmt() ast.Stmt {
	stmt := &ast.SendStmt{Chan: s.chn.Expr(), Value: s.val.Expr()}
	return stmt
}
