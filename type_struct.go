package go_ast_dsl

import "go/ast"

// StructField
//
type StructField struct {
	name string
	typ  TypeExprProvider
	tag  string
}

func NewStructField() *StructField {
	return &StructField{}
}

func (f *StructField) WithType(typ TypeExprProvider) *StructField {
	f.typ = typ
	return f
}

func (f *StructField) WithName(name string) *StructField {
	f.name = name
	return f
}

func (f *StructField) WithTag(tag string) *StructField {
	f.tag = tag
	return f
}

func (f *StructField) FieldDecl() *ast.Field {
	// &ast.Field{Doc: doc, Names: idents, Type: typ, Tag: tag, Comment: p.lineComment}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	ident := &ast.Ident{Name: f.name}
	field.Names = append(field.Names, ident)
	if len(f.tag) > 0 {
		field.Tag = NewStringLiteral(f.tag).BasicLit()
	}
	return field
}

// StructAnonField
//
type StructAnonField struct {
	typ TypeExprProvider
	tag string
}

func NewStructAnonField() *StructAnonField {
	return &StructAnonField{}
}

func (f *StructAnonField) WithType(typ TypeExprProvider) *StructAnonField {
	f.typ = typ
	return f
}

func (f *StructAnonField) WithTag(tag string) *StructAnonField {
	f.tag = tag
	return f
}

func (f *StructAnonField) FieldDecl() *ast.Field {
	// &ast.Field{Doc: doc, Names: idents, Type: typ, Tag: tag, Comment: p.lineComment}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	if len(f.tag) > 0 {
		field.Tag = NewStringLiteral(f.tag).BasicLit()
	}
	return field
}

// StructFieldList
//
type StructFieldList struct {
	names []string
	typ   TypeExprProvider
	tag   string
}

func NewStructFieldList() *StructFieldList {
	return &StructFieldList{}
}

func (f *StructFieldList) WithType(typ TypeExprProvider) *StructFieldList {
	f.typ = typ
	return f
}

func (f *StructFieldList) AddName(name string) *StructFieldList {
	f.names = append(f.names, name)
	return f
}

func (f *StructFieldList) WithTag(tag string) *StructFieldList {
	f.tag = tag
	return f
}

func (f *StructFieldList) FieldDecl() *ast.Field {
	// &ast.Field{Doc: doc, Names: idents, Type: typ, Tag: tag, Comment: p.lineComment}
	field := &ast.Field{Type: f.typ.TypeExpr()}
	for _, name := range f.names {
		ident := &ast.Ident{Name: name}
		field.Names = append(field.Names, ident)
	}
	if len(f.tag) > 0 {
		field.Tag = NewStringLiteral(f.tag).BasicLit()
	}
	return field
}

// StructType
//
type StructType struct {
	fields []FieldDeclProvider
}

func NewStructType() *StructType {
	return &StructType{}
}

func (t *StructType) AddField(field *StructField) *StructType {
	t.fields = append(t.fields, field)
	return t
}

func (t *StructType) AddAnonField(field *StructAnonField) *StructType {
	t.fields = append(t.fields, field)
	return t
}

func (t *StructType) AddList(field *StructFieldList) *StructType {
	t.fields = append(t.fields, field)
	return t
}

func (t *StructType) TypeExpr() ast.Expr {
	fields := &ast.FieldList{}
	for _, field := range t.fields {
		fields.List = append(fields.List, field.FieldDecl())
	}
	return &ast.StructType{Fields: fields}
}
