package go_ast_dsl

import (
	"go/ast"
)

// CallExprProvider
//
type CallExprProvider interface {
	CallExpr() *ast.CallExpr
}

// CallExpr
//
type CallExpr struct {
	fun   ExprProvider
	args []ExprProvider
	ellipses bool
}

func NewCallExpr() *CallExpr {
	return &CallExpr{}
}

func (s *CallExpr) WithFunction(exp ExprProvider) *CallExpr {
	s.fun = exp
	return s
}
func (s *CallExpr) AddArg(arg ExprProvider) *CallExpr {
	s.args = append(s.args, arg)
	return s
}

func (s *CallExpr) AddArgWithEllipses(arg ExprProvider) *CallExpr {
	s.args = append(s.args, arg)
	s.ellipses = true
	return s
}

func (s *CallExpr) CallExpr() *ast.CallExpr {
	expr := &ast.CallExpr{Fun: s.fun.Expr()}
	for _, arg := range s.args {
		expr.Args = append(expr.Args, arg.Expr())
	}
	if s.ellipses {
		expr.Ellipsis = 1
	}
	return expr
}

func (s *CallExpr) Expr() ast.Expr {
	return s.CallExpr()
}
