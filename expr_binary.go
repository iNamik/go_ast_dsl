package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

type BinaryTokenProvider interface {
	BinaryToken() token.Token
}

type BinaryOp token.Token

const (
	BINARY_LOR     BinaryOp = BinaryOp(token.LOR)
	BINARY_LAND             = BinaryOp(token.LAND)
	BINARY_EQ               = BinaryOp(token.EQL)
	BINARY_NOT_EQ           = BinaryOp(token.NEQ)
	BINARY_LESS             = BinaryOp(token.LSS)
	BINARY_LEQ              = BinaryOp(token.LEQ)
	BINARY_GREATER          = BinaryOp(token.GTR)
	BINARY_GEQ              = BinaryOp(token.GEQ)
	BINARY_ADD              = BinaryOp(token.ADD)
	BINARY_SUB              = BinaryOp(token.SUB)
	BINARY_OR               = BinaryOp(token.OR)
	BINARY_XOR              = BinaryOp(token.XOR)
	BINARY_MUL              = BinaryOp(token.MUL)
	BINARY_DIV              = BinaryOp(token.QUO)
	BINARY_MOD              = BinaryOp(token.REM)
	BINARY_SHL              = BinaryOp(token.SHL)
	BINARY_SHR              = BinaryOp(token.SHR)
	BINARY_AND              = BinaryOp(token.AND)
	BINARY_AND_NOT          = BinaryOp(token.AND_NOT)
)

func (t BinaryOp) BinaryToken() token.Token {
	return token.Token(t)
}

// BinaryExpr
//
type BinaryExpr struct {
	left  ExprProvider
	op    BinaryTokenProvider
	right ExprProvider
}

func NewBinaryExpr() *BinaryExpr {
	return &BinaryExpr{}
}

func (s *BinaryExpr) WithOperator(op BinaryTokenProvider) *BinaryExpr {
	s.op = op
	return s
}

func (s *BinaryExpr) WithLeftExpr(left ExprProvider) *BinaryExpr {
	s.left = left
	return s
}

func (s *BinaryExpr) WithRightExpr(right ExprProvider) *BinaryExpr {
	s.right = right
	return s
}

func (s *BinaryExpr) Expr() ast.Expr {
	expr := &ast.BinaryExpr{Op: s.op.BinaryToken(), X: s.left.Expr(), Y: s.right.Expr()}
	return expr
}
