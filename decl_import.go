package go_ast_dsl

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Import
//
type Import struct {
	Ast *ast.GenDecl
}

func NewImport(pkg string) *Import {
	decl := &ast.GenDecl{Tok: token.IMPORT}
	path := &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", pkg)}
	spec := &ast.ImportSpec{Path: path}
	decl.Specs = append(decl.Specs, spec)
	return &Import{Ast: decl}
}

func (i *Import) WithAlias(alias string) *Import {
	ident := &ast.Ident{Name: alias}
	i.Ast.Specs[0].(*ast.ImportSpec).Name = ident
	return i
}

// Currently Unsupported
//func (i *Import) WithLeadComment(comment *CommentGroup) *Import {
//	i.Ast.Doc = comment.Ast
//	return i
//}

func (i *Import) WithLineComment(comment *CommentGroup) *Import {
	i.Ast.Specs[0].(*ast.ImportSpec).Comment = comment.Ast
	return i
}

// Import Group
//
type ImportGroup struct {
	decl *ast.GenDecl
}

func NewImportGroup() *ImportGroup {
	decl := &ast.GenDecl{Tok: token.IMPORT}
	decl.Lparen = 1 // Forces the parens in the output
	ig := &ImportGroup{decl: decl}
	return ig
}

func (ig *ImportGroup) Import(pkg string) *ImportGroup {
	path := &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", pkg)}
	spec := &ast.ImportSpec{Path: path}
	ig.decl.Specs = append(ig.decl.Specs, spec)
	return ig
}

func (ig *ImportGroup) ImportWithAlias(alias string, pkg string) *ImportGroup {
	path := &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", pkg)}
	ident := &ast.Ident{Name: alias}
	spec := &ast.ImportSpec{Name: ident, Path: path}
	ig.decl.Specs = append(ig.decl.Specs, spec)
	return ig
}
