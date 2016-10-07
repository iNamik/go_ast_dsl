package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_StructEmpty = `package go_ast_dsl_test

type a struct {
}
`

func Test_Type_StructEmpty(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructEmpty.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType()))
	testAst(t, "Test_Type_StructEmpty", []byte(src_Test_Type_StructEmpty), file.file)
}

const src_Test_Type_StructNamedField = `package go_ast_dsl_test

type a struct {
	b int
}
`

func Test_Type_StructNamedField(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructNamedField.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddField(NewStructField().
					WithName("b").
					WithType(NewTypeRef().
						WithName("int")))))
	testAst(t, "Test_Type_StructNamedField", []byte(src_Test_Type_StructNamedField), file.file)
}

const src_Test_Type_StructNamedFields = `package go_ast_dsl_test

type a struct {
	b	int
	c	string
}
`

func Test_Type_StructNamedFields(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructNamedFields.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddField(NewStructField().
					WithName("b").
					WithType(NewTypeRef().
						WithName("int"))).
				AddField(NewStructField().
					WithName("c").
					WithType(NewTypeRef().
						WithName("string")))))
	testAst(t, "Test_Type_StructNamedFields", []byte(src_Test_Type_StructNamedFields), file.file)
}

const src_Test_Type_StructNamedFieldList = `package go_ast_dsl_test

type a struct {
	b, c int
}
`

func Test_Type_StructNamedFieldList(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructNamedFieldList.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddList(NewStructFieldList().
					AddName("b").
					AddName("c").
					WithType(NewTypeRef().
						WithName("int")))))
	testAst(t, "Test_Type_StructNamedFieldList", []byte(src_Test_Type_StructNamedFieldList), file.file)
}

const src_Test_Type_StructNamedFieldWithTag = `package go_ast_dsl_test

type a struct {
	b int "tag"
}
`

func Test_Type_StructNamedFieldWithTag(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructNamedFieldWithTag.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddField(NewStructField().
					WithName("b").
					WithType(NewTypeRef().
						WithName("int")).
					WithTag(`"tag"`))))
	testAst(t, "Test_Type_StructNamedFieldWithTag", []byte(src_Test_Type_StructNamedFieldWithTag), file.file)
}

const src_Test_Type_StructNamedFieldWithTag2 = `package go_ast_dsl_test

type a struct {
	b int ` + "`tag`" + `
}
`

func Test_Type_StructNamedFieldWithTag2(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructNamedFieldWithTag2.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddField(NewStructField().
					WithName("b").
					WithType(NewTypeRef().
						WithName("int")).
					WithTag("`tag`"))))
	testAst(t, "Test_Type_StructNamedFieldWithTag2", []byte(src_Test_Type_StructNamedFieldWithTag2), file.file)
}

const src_Test_Type_StructAnonField = `package go_ast_dsl_test

type a struct {
	MyType
}
`

func Test_Type_StructAnonField(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructAnonField.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddAnonField(NewStructAnonField().
					WithType(NewTypeRef().
						WithName("MyType")))))
	testAst(t, "Test_Type_StructAnonField", []byte(src_Test_Type_StructAnonField), file.file)
}

const src_Test_Type_StructAnonFieldWithTag = `package go_ast_dsl_test

type a struct {
	MyType "tag"
}
`

func Test_Type_StructAnonFieldWithTag(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructAnonFieldWithTag.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddAnonField(NewStructAnonField().
					WithType(NewTypeRef().
						WithName("MyType")).
					WithTag(`"tag"`))))
	testAst(t, "Test_Type_StructAnonFieldWithTag", []byte(src_Test_Type_StructAnonFieldWithTag), file.file)
}

const src_Test_Type_StructAnonFields = `package go_ast_dsl_test

type a struct {
	MyType1
	MyType2
}
`

func Test_Type_StructAnonFields(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructAnonFields.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddAnonField(NewStructAnonField().
					WithType(NewTypeRef().
						WithName("MyType1"))).
				AddAnonField(NewStructAnonField().
					WithType(NewTypeRef().
						WithName("MyType2")))))
	testAst(t, "Test_Type_StructAnonFields", []byte(src_Test_Type_StructAnonFields), file.file)
}

const src_Test_Type_StructAnonFields2 = `package go_ast_dsl_test

type a struct {
	MyType1
	b	int
	MyType2
}
`

func Test_Type_StructAnonFields2(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_StructAnonFields2.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewStructType().
				AddAnonField(NewStructAnonField().
					WithType(NewTypeRef().
						WithName("MyType1"))).
				AddField(NewStructField().
					WithName("b").
					WithType(NewTypeRef().
						WithName("int"))).
				AddAnonField(NewStructAnonField().
					WithType(NewTypeRef().
						WithName("MyType2")))))
	testAst(t, "Test_Type_StructAnonFields2", []byte(src_Test_Type_StructAnonFields2), file.file)
}
