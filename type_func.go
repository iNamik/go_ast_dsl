package go_ast_dsl

import "go/ast"

type FuncType struct {
	params  FieldListProvider
	results FieldListProvider
}

func NewFuncType() *FuncType {
	return &FuncType{}
}

func (t *FuncType) WithNamedParams(params *NamedMethodParamList) *FuncType {
	t.params = params
	return t
}

func (t *FuncType) WithAnonParams(params *AnonMethodParamList) *FuncType {
	t.params = params
	return t
}

func (t *FuncType) WithNamedResults(results *NamedMethodResultList) *FuncType {
	t.results = results
	return t
}

func (t *FuncType) WithAnonResults(results *AnonMethodResultList) *FuncType {
	t.results = results
	return t
}

func (t *FuncType) TypeExpr() ast.Expr {
	// &ast.FuncType{Func: token.NoPos, Params: params, Results: results}
	// &ast.Field{Doc: doc, Names: idents, Type: typ, Comment: p.lineComment}
	//
	typ := &ast.FuncType{}
	if nil != t.params {
		typ.Params = t.params.FieldList()
	}
	if nil != t.results {
		typ.Results = t.results.FieldList()
	}
	return typ
}
