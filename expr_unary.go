package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

type UnaryTokenProvider interface {
	UnaryToken() token.Token
}

// UnaryOp
//
type UnaryOp token.Token

const (
	UNARY_PLUS  UnaryOp = UnaryOp(token.ADD)
	UNARY_MINUS         = UnaryOp(token.SUB)
	UNARY_STAR          = UnaryOp(token.MUL)
	UNARY_NOT           = UnaryOp(token.NOT)
	UNARY_XOR           = UnaryOp(token.XOR)
	UNARY_AND           = UnaryOp(token.AND)
	UNARY_RECV          = UnaryOp(token.ARROW)
)

func (t UnaryOp) UnaryToken() token.Token {
	return token.Token(t)
}

// UnaryExpr
//
type UnaryExpr struct {
	op   UnaryTokenProvider
	expr ExprProvider
}

func NewUnaryExpr() *UnaryExpr {
	return &UnaryExpr{}
}

func (s *UnaryExpr) WithOperator(op UnaryTokenProvider) *UnaryExpr {
	s.op = op
	return s
}
func (s *UnaryExpr) WithExpr(expr ExprProvider) *UnaryExpr {
	s.expr = expr
	return s
}

func (s *UnaryExpr) Expr() ast.Expr {
	expr := &ast.UnaryExpr{Op: s.op.UnaryToken(), X: s.expr.Expr()}
	return expr
}
