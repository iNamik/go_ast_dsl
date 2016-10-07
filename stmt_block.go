package go_ast_dsl

import "go/ast"

// BlockStmt
//
type BlockStmt struct {
	list []StmtProvider
}

func NewBlockStmt() *BlockStmt {
	return &BlockStmt{}
}

func (s *BlockStmt) Add(stmt StmtProvider) *BlockStmt {
	s.list = append(s.list, stmt)
	return s
}

func (s *BlockStmt) BlockStmt() *ast.BlockStmt {
	stmt := &ast.BlockStmt{}
	for _, result := range s.list {
		stmt.List = append(stmt.List, result.Stmt())
	}
	return stmt
}

func (s *BlockStmt) Stmt() ast.Stmt {
	return s.BlockStmt()
}
