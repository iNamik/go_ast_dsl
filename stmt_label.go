package go_ast_dsl

import "go/ast"

// LabelStmt
//
type LabelStmt struct {
	name string
	stmt StmtProvider
}

func NewLabelStmt() *LabelStmt {
	return &LabelStmt{}
}

func (s *LabelStmt) WithName(name string) *LabelStmt {
	s.name = name
	return s
}

func (s *LabelStmt) WithStmt(stmt StmtProvider) *LabelStmt {
	s.stmt = stmt
	return s
}

func (s *LabelStmt) Stmt() ast.Stmt {
	ident := &ast.Ident{Name: s.name}
	stmt := &ast.LabeledStmt{Label: ident}
	if nil != s.stmt {
		stmt.Stmt = s.stmt.Stmt() // Yikes!
	} else {
		stmt.Stmt = &ast.EmptyStmt{Implicit: true}
	}
	return stmt
}
