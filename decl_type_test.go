package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_Alias = `package go_ast_dsl_test

type a int
`

func Test_Type_Alias(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Alias.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewTypeRef().WithName("int")))
	testAst(t, "Test_Type_Alias", []byte(src_Test_Type_Alias), file.file)
}

const src_Test_Type_AliasWithPkg = `package go_ast_dsl_test

type a pkg.Type
`

func Test_Type_AliasWithPkg(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_AliasWithPkg.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewTypeRef().WithPackage("pkg").WithName("Type")))
	testAst(t, "Test_Type_AliasWithPkg", []byte(src_Test_Type_AliasWithPkg), file.file)
}

const src_Test_Type_GroupAlias = `package go_ast_dsl_test

type (
	a	int
	b	string
)
`

func Test_Type_GroupAlias(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_GroupAlias.txt")
	group := NewTypeDefGroup().
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewTypeRef().WithName("int"))).
		AddType(NewTypeDef().
			WithName("b").
			WithType(NewTypeRef().WithName("string")))
	file.AddTypeGroup(group)
	testAst(t, "Test_Type_GroupAlias", []byte(src_Test_Type_GroupAlias), file.file)
}
