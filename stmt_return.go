package go_ast_dsl

import "go/ast"

// ReturnStmt
//
type ReturnStmt struct {
	list []ExprProvider
}

func NewReturnStmt() *ReturnStmt {
	return &ReturnStmt{}
}

func (s *ReturnStmt) AddResult(result ExprProvider) *ReturnStmt {
	s.list = append(s.list, result)
	return s
}

func (s *ReturnStmt) Stmt() ast.Stmt {
	stmt := &ast.ReturnStmt{}
	for _, result := range s.list {
		stmt.Results = append(stmt.Results, result.Expr())
	}
	return stmt
}
