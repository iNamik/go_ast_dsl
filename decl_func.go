package go_ast_dsl

import "go/ast"

// Function
//
type Function struct {
	name    string
	params  FieldListProvider
	results FieldListProvider
	body    *BlockStmt
}

func NewFunction() *Function {
	return &Function{}
}

func (f *Function) WithFunctionName(name string) *Function {
	f.name = name
	return f
}

func (f *Function) WithNamedParams(params *NamedMethodParamList) *Function {
	f.params = params
	return f
}

func (f *Function) WithAnonParams(params *AnonMethodParamList) *Function {
	f.params = params
	return f
}

func (f *Function) WithNamedResults(results *NamedMethodResultList) *Function {
	f.results = results
	return f
}

func (f *Function) WithAnonResults(results *AnonMethodResultList) *Function {
	f.results = results
	return f
}

func (f *Function) WithFunctionBody(body *BlockStmt) *Function {
	f.body = body
	return f
}

func (f *Function) Decl() ast.Decl {
	decl := &ast.FuncDecl{}
	decl.Name = &ast.Ident{Name: f.name}
	typ := &ast.FuncType{}
	if nil != f.params {
		typ.Params = f.params.FieldList()
	}
	if nil != f.results {
		typ.Results = f.results.FieldList()
	}
	decl.Type = typ
	if nil != f.body {
		decl.Body = f.body.BlockStmt()
	} else {
		decl.Body = &ast.BlockStmt{}
	}
	return decl
}

//  NamedMethodReceiver
//
type NamedMethodReceiver struct {
	name string
	typ  TypeExprProvider
}

func NewNamedMethodReceiver() *NamedMethodReceiver {
	return &NamedMethodReceiver{}
}

func (r *NamedMethodReceiver) WithName(name string) *NamedMethodReceiver {
	r.name = name
	return r
}

func (r *NamedMethodReceiver) WithType(typ TypeExprProvider) *NamedMethodReceiver {
	r.typ = typ
	return r
}

func (r *NamedMethodReceiver) FieldList() *ast.FieldList {
	params := &ast.FieldList{}
	params.List = append(params.List, r.FieldDecl())
	return params

}

func (r *NamedMethodReceiver) FieldDecl() *ast.Field {
	// &ast.Field{Names: idents, Type: typ}
	field := &ast.Field{Type: r.typ.TypeExpr()}
	ident := &ast.Ident{Name: r.name}
	field.Names = append(field.Names, ident)
	return field
}

//  AnonMethodReceiver
//
type AnonMethodReceiver struct {
	typ TypeExprProvider
}

func NewAnonMethodReceiver() *AnonMethodReceiver {
	return &AnonMethodReceiver{}
}

func (r *AnonMethodReceiver) WithType(typ TypeExprProvider) *AnonMethodReceiver {
	r.typ = typ
	return r
}

func (r *AnonMethodReceiver) FieldList() *ast.FieldList {
	params := &ast.FieldList{}
	params.List = append(params.List, r.FieldDecl())
	return params

}

func (r *AnonMethodReceiver) FieldDecl() *ast.Field {
	// &ast.Field{Names: idents, Type: typ}
	field := &ast.Field{Type: r.typ.TypeExpr()}
	return field
}

// Method
//
type Method struct {
	recv    FieldListProvider
	name    string
	params  FieldListProvider
	results FieldListProvider
	body    BlockStmtProvider
}

func NewMethod() *Method {
	return &Method{}
}

func (m *Method) WithNamedReceiver(recv *NamedMethodReceiver) *Method {
	m.recv = recv
	return m
}

func (m *Method) WithAnonReceiver(recv *AnonMethodReceiver) *Method {
	m.recv = recv
	return m
}

func (m *Method) WithName(name string) *Method {
	m.name = name
	return m
}

func (m *Method) WithBody(body BlockStmtProvider) *Method {
	m.body = body
	return m
}

func (m *Method) Decl() ast.Decl {
	decl := &ast.FuncDecl{}
	decl.Recv = m.recv.FieldList()
	decl.Name = &ast.Ident{Name: m.name}
	typ := &ast.FuncType{}
	if nil != m.params {
		typ.Params = m.params.FieldList()
	}
	if nil != m.results {
		typ.Results = m.results.FieldList()
	}
	decl.Type = typ
	if nil != m.body {
		decl.Body = m.body.BlockStmt()
	} else {
		decl.Body = &ast.BlockStmt{}
	}
	return decl
}
