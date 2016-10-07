package go_ast_dsl

import "go/ast"

// InterfaceEmbedField
//
type InterfaceEmbedField struct {
	typ *TypeRef
}

func NewInterfaceEmbedField() *InterfaceEmbedField {
	return &InterfaceEmbedField{}
}

func (f *InterfaceEmbedField) WithType(typ *TypeRef) *InterfaceEmbedField {
	f.typ = typ
	return f
}

func (f *InterfaceEmbedField) FieldDecl() *ast.Field {
	// &ast.Field{Doc: doc, Names: idents, Type: typ, Comment: p.lineComment}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	return field
}

// NamedMethodParamList
//
type NamedMethodParamList struct {
	list []FieldDeclProvider
}

func NewNamedMethodParamList() *NamedMethodParamList {
	return &NamedMethodParamList{}
}

func (l *NamedMethodParamList) AddParam(param *NamedMethodParam) *NamedMethodParamList {
	l.list = append(l.list, param)
	return l
}

func (l *NamedMethodParamList) AddParams(params *NamedMethodParams) *NamedMethodParamList {
	l.list = append(l.list, params)
	return l
}

func (l *NamedMethodParamList) FieldList() *ast.FieldList {
	params := &ast.FieldList{}
	for _, field := range l.list {
		params.List = append(params.List, field.FieldDecl())
	}
	return params
}

// NamedMethodParam
//
type NamedMethodParam struct {
	name string
	typ  TypeExprProvider
}

func NewNamedMethodParam() *NamedMethodParam {
	return &NamedMethodParam{}
}

func (p *NamedMethodParam) WithName(name string) *NamedMethodParam {
	p.name = name
	return p
}

func (p *NamedMethodParam) WithType(typ TypeExprProvider) *NamedMethodParam {
	p.typ = typ
	return p
}

func (f *NamedMethodParam) FieldDecl() *ast.Field {
	// &ast.Field{Names: idents, Type: typ}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	ident := &ast.Ident{Name: f.name}
	field.Names = append(field.Names, ident)
	return field
}

// NamedMethodParams
//
type NamedMethodParams struct {
	names []string
	typ   TypeExprProvider
}

func NewNamedMethodParams() *NamedMethodParams {
	return &NamedMethodParams{}
}

func (p *NamedMethodParams) AddName(name string) *NamedMethodParams {
	p.names = append(p.names, name)
	return p
}

func (p *NamedMethodParams) WithType(typ TypeExprProvider) *NamedMethodParams {
	p.typ = typ
	return p
}

func (f *NamedMethodParams) FieldDecl() *ast.Field {
	// &ast.Field{Names: idents, Type: typ}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	for _, name := range f.names {
		ident := &ast.Ident{Name: name}
		field.Names = append(field.Names, ident)
	}
	return field
}

// AnonMethodParamList
//
type AnonMethodParamList struct {
	list []FieldDeclProvider
}

func NewAnonMethodParamList() *AnonMethodParamList {
	return &AnonMethodParamList{}
}

func (l *AnonMethodParamList) AddParam(param *AnonMethodParam) *AnonMethodParamList {
	l.list = append(l.list, param)
	return l
}

func (l *AnonMethodParamList) FieldList() *ast.FieldList {
	params := &ast.FieldList{}
	for _, field := range l.list {
		params.List = append(params.List, field.FieldDecl())
	}
	return params
}

// AnonMethodParam
//
type AnonMethodParam struct {
	typ TypeExprProvider
}

func NewAnonMethodParam() *AnonMethodParam {
	return &AnonMethodParam{}
}

func (p *AnonMethodParam) WithType(typ TypeExprProvider) *AnonMethodParam {
	p.typ = typ
	return p
}

func (f *AnonMethodParam) FieldDecl() *ast.Field {
	// &ast.Field{Names: idents, Type: typ}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	return field
}

// NamedMethodResultList
//
type NamedMethodResultList struct {
	list []FieldDeclProvider
}

func NewNamedMethodResultList() *NamedMethodResultList {
	return &NamedMethodResultList{}
}

func (l *NamedMethodResultList) AddResult(result *NamedMethodResult) *NamedMethodResultList {
	l.list = append(l.list, result)
	return l
}

func (l *NamedMethodResultList) AddResults(results *NamedMethodResults) *NamedMethodResultList {
	l.list = append(l.list, results)
	return l
}

func (l *NamedMethodResultList) FieldList() *ast.FieldList {
	results := &ast.FieldList{}
	for _, field := range l.list {
		results.List = append(results.List, field.FieldDecl())
	}
	return results
}

// NamedMethodResult
//
type NamedMethodResult struct {
	name string
	typ  TypeExprProvider
}

func NewNamedMethodResult() *NamedMethodResult {
	return &NamedMethodResult{}
}

func (p *NamedMethodResult) WithName(name string) *NamedMethodResult {
	p.name = name
	return p
}

func (p *NamedMethodResult) WithType(typ TypeExprProvider) *NamedMethodResult {
	p.typ = typ
	return p
}

func (f *NamedMethodResult) FieldDecl() *ast.Field {
	//
	field := &ast.Field{Type: f.typ.TypeExpr()}
	ident := &ast.Ident{Name: f.name}
	field.Names = append(field.Names, ident)
	return field
}

// NamedMethodResults
//
type NamedMethodResults struct {
	names []string
	typ   TypeExprProvider
}

func NewNamedMethodResults() *NamedMethodResults {
	return &NamedMethodResults{}
}

func (p *NamedMethodResults) AddName(name string) *NamedMethodResults {
	p.names = append(p.names, name)
	return p
}

func (p *NamedMethodResults) WithType(typ TypeExprProvider) *NamedMethodResults {
	p.typ = typ
	return p
}

func (f *NamedMethodResults) FieldDecl() *ast.Field {
	//
	field := &ast.Field{Type: f.typ.TypeExpr()}
	for _, name := range f.names {
		ident := &ast.Ident{Name: name}
		field.Names = append(field.Names, ident)
	}
	return field
}

// AnonMethodResultList
//
type AnonMethodResultList struct {
	list []FieldDeclProvider
}

func NewAnonMethodResultList() *AnonMethodResultList {
	return &AnonMethodResultList{}
}

func (l *AnonMethodResultList) AddResult(result *AnonMethodResult) *AnonMethodResultList {
	l.list = append(l.list, result)
	return l
}

func (l *AnonMethodResultList) FieldList() *ast.FieldList {
	results := &ast.FieldList{}
	for _, field := range l.list {
		results.List = append(results.List, field.FieldDecl())
	}
	return results
}

// AnonMethodResult
//
type AnonMethodResult struct {
	typ TypeExprProvider
}

func NewAnonMethodResult() *AnonMethodResult {
	return &AnonMethodResult{}
}

func (p *AnonMethodResult) WithType(typ TypeExprProvider) *AnonMethodResult {
	p.typ = typ
	return p
}

func (f *AnonMethodResult) FieldDecl() *ast.Field {
	//
	field := &ast.Field{Type: f.typ.TypeExpr()}
	return field
}

// InterfaceMethodField
//
type InterfaceMethodField struct {
	name    string
	params  FieldListProvider
	results FieldListProvider
}

func NewInterfaceMethodField() *InterfaceMethodField {
	return &InterfaceMethodField{}
}

func (f *InterfaceMethodField) WithMethodName(name string) *InterfaceMethodField {
	f.name = name
	return f
}

func (f *InterfaceMethodField) WithNamedParams(params *NamedMethodParamList) *InterfaceMethodField {
	f.params = params
	return f
}

func (f *InterfaceMethodField) WithAnonParams(params *AnonMethodParamList) *InterfaceMethodField {
	f.params = params
	return f
}

func (f *InterfaceMethodField) WithNamedResults(results *NamedMethodResultList) *InterfaceMethodField {
	f.results = results
	return f
}

func (f *InterfaceMethodField) WithAnonResults(results *AnonMethodResultList) *InterfaceMethodField {
	f.results = results
	return f
}

func (f *InterfaceMethodField) FieldDecl() *ast.Field {
	// &ast.FuncType{Func: token.NoPos, Params: params, Results: results}
	// &ast.Field{Doc: doc, Names: idents, Type: typ, Comment: p.lineComment}
	//
	typ := &ast.FuncType{}
	if nil != f.params {
		typ.Params = f.params.FieldList()
	}
	if nil != f.results {
		typ.Results = f.results.FieldList()
	}
	field := &ast.Field{Type: typ}
	ident := &ast.Ident{Name: f.name}
	field.Names = append(field.Names, ident)
	return field
}

// InterfaceType
//
type InterfaceType struct {
	fields []FieldDeclProvider
}

func NewInterfaceType() *InterfaceType {
	return &InterfaceType{}
}

func (t *InterfaceType) AddMethod(field *InterfaceMethodField) *InterfaceType {
	t.fields = append(t.fields, field)
	return t
}

func (t *InterfaceType) AddEmbedded(field *InterfaceEmbedField) *InterfaceType {
	t.fields = append(t.fields, field)
	return t
}

func (t *InterfaceType) TypeExpr() ast.Expr {
	fields := &ast.FieldList{}
	for _, field := range t.fields {
		fields.List = append(fields.List, field.FieldDecl())
	}
	return &ast.InterfaceType{Methods: fields}
}
