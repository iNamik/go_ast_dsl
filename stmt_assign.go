package go_ast_dsl

import (
	"go/ast"
	"go/token"
)

type AssignTokenProvider interface {
	AssignToken() token.Token
}

type AssignOp token.Token

const (
	ASSIGN_EQUALS  AssignOp = AssignOp(token.ASSIGN)
	ASSIGN_DEFINE           = AssignOp(token.DEFINE)
	ASSIGN_ADD              = AssignOp(token.ADD_ASSIGN)
	ASSIGN_SUB              = AssignOp(token.SUB_ASSIGN)
	ASSIGN_MUL              = AssignOp(token.MUL_ASSIGN)
	ASSIGN_DIV              = AssignOp(token.QUO_ASSIGN)
	ASSIGN_MOD              = AssignOp(token.REM_ASSIGN)
	ASSIGN_AND              = AssignOp(token.AND_ASSIGN)
	ASSIGN_OR               = AssignOp(token.OR_ASSIGN)
	ASSIGN_XOR              = AssignOp(token.XOR_ASSIGN)
	ASSIGN_SHAL             = AssignOp(token.SHL_ASSIGN)
	ASSIGN_SHR              = AssignOp(token.SHR_ASSIGN)
	ASSIGN_AND_NOT          = AssignOp(token.AND_NOT_ASSIGN)
)

func (t AssignOp) AssignToken() token.Token {
	return token.Token(t)
}

// AssignStmt
//
type AssignStmt struct {
	lhs ExprProvider
	op  AssignTokenProvider
	rhs ExprProvider
}

func NewAssignStmt() *AssignStmt {
	return &AssignStmt{op: ASSIGN_EQUALS}
}

func (s *AssignStmt) WithLhsExpr(lhs ExprProvider) *AssignStmt {
	s.lhs = lhs
	return s
}

func (s *AssignStmt) WithRhsExpr(rhs ExprProvider) *AssignStmt {
	s.rhs = rhs
	return s
}

func (s *AssignStmt) WithOperator(op AssignTokenProvider) *AssignStmt {
	s.op = op
	return s
}

func (s *AssignStmt) Stmt() ast.Stmt {
	stmt := &ast.AssignStmt{}
	stmt.Lhs = append(stmt.Lhs, s.lhs.Expr())
	stmt.Tok = s.op.AssignToken()
	stmt.Rhs = append(stmt.Rhs, s.rhs.Expr())
	return stmt
}

// AssignListStmt
//
type AssignListStmt struct {
	lhs []ExprProvider
	op  AssignTokenProvider
	rhs []ExprProvider
}

func NewAssignListStmt() *AssignListStmt {
	return &AssignListStmt{op: ASSIGN_EQUALS}
}

func (s *AssignListStmt) AddLhsExpr(lhs ExprProvider) *AssignListStmt {
	s.lhs = append(s.lhs, lhs)
	return s
}

func (s *AssignListStmt) AddRhsExpr(rhs ExprProvider) *AssignListStmt {
	s.rhs = append(s.rhs, rhs)
	return s
}

func (s *AssignListStmt) WithOperator(op AssignTokenProvider) *AssignListStmt {
	s.op = op
	return s
}

func (s *AssignListStmt) Stmt() ast.Stmt {
	stmt := &ast.AssignStmt{}
	for _, lhs := range s.lhs {
		stmt.Lhs = append(stmt.Lhs, lhs.Expr())
	}
	stmt.Tok = s.op.AssignToken()
	for _, rhs := range s.rhs {
		stmt.Rhs = append(stmt.Rhs, rhs.Expr())
	}
	return stmt
}
