package go_ast_dsl

import (
	"go/token"
	"go/ast"
)

// FallthroughStmt
//
type FallthroughStmt struct {}

func NewFallthroughStmt() *FallthroughStmt {
	return &FallthroughStmt{}
}

func (s *FallthroughStmt) BranchStmt() *ast.BranchStmt {
	return &ast.BranchStmt{Tok: token.FALLTHROUGH}
}

func (s *FallthroughStmt) Stmt() ast.Stmt {
	return s.BranchStmt()
}
