package go_ast_dsl

import "go/ast"

// ChanType
//
type ChanType struct {
	send bool
	recv bool
	typ  TypeExprProvider
}

func NewChanType() *ChanType {
	return &ChanType{}
}

func (t *ChanType) WithSend() *ChanType {
	t.send = true
	return t
}

func (t *ChanType) WithRecv() *ChanType {
	t.recv = true
	return t
}

func (t *ChanType) WithSendAndReceive() *ChanType {
	t.send = true
	t.recv = true
	return t
}

func (t *ChanType) WithType(typ TypeExprProvider) *ChanType {
	t.typ = typ
	return t
}

func (t *ChanType) TypeExpr() ast.Expr {
	// &ast.ChanType{Begin: pos, Arrow: arrow, Dir: dir, Value: value}
	var dir ast.ChanDir
	if t.send {
		dir |= ast.SEND
	}
	if t.recv {
		dir |= ast.RECV
	}
	return &ast.ChanType{Dir: dir, Value: t.typ.TypeExpr()}
}
