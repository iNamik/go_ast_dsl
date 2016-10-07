package go_ast_dsl

import "fmt"
import "go/ast"

type CommentGroup struct {
	Ast *ast.CommentGroup
}

func NewCommentGroup() *CommentGroup {
	group := &ast.CommentGroup{}
	return &CommentGroup{Ast: group}
}

func (c *CommentGroup) addBlock(comment string) *CommentGroup {
	c.Ast.List = append(c.Ast.List, &ast.Comment{Text: fmt.Sprintf("/*%s*/", comment)})
	return c
}

func (c *CommentGroup) addLine(comment string) *CommentGroup {
	c.Ast.List = append(c.Ast.List, &ast.Comment{Text: fmt.Sprintf("//%s", comment)})
	return c
}

func (c *CommentGroup) Empty() bool {
	return len(c.Ast.List) > 0
}
