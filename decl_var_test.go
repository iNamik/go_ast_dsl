package go_ast_dsl

import (
	"testing"
)

const src_Test_Var_IntUnassigned = `package go_ast_dsl_test

var a int
`

func Test_Var_IntUnassigned(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Var_IntUnassigned.txt").
		AddVar(NewVar().
			WithName("a").
			WithType(NewTypeRef().WithName("int")))
	testAst(t, "Test_Var_IntUnassigned", []byte(src_Test_Var_IntUnassigned), file.file)
}

const src_Test_Var_IntArrayUnassigned = `package go_ast_dsl_test

var a []int
`

func Test_Var_IntArrayUnassigned(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Var_IntArrayUnassigned.txt").
		AddVar(NewVar().
			WithName("a").
			WithType(NewArrayType().WithType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Var_IntArrayUnassigned", []byte(src_Test_Var_IntArrayUnassigned), file.file)
}

const src_Test_Var_IntAssigned = `package go_ast_dsl_test

var a int = 1
`

func Test_Var_IntAssigned(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Var_IntAssigned.txt").
		AddVar(NewVar().
			WithName("a").
			WithType(NewTypeRef().WithName("int")).
			WithValue(NewIntLiteral("1")))
	testAst(t, "Test_Var_IntAssigned", []byte(src_Test_Var_IntAssigned), file.file)
}

const src_Test_Var_ListIntUnassigned = `package go_ast_dsl_test

var a, b int
`

func Test_Var_ListIntUnassigned(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Var_ListIntUnassigned.txt")
	list := NewVarList().
		WithType(NewTypeRef().WithName("int")).
		Add(NewVar().
			WithName("a").
			WithType(NewTypeRef().WithName("int"))).
		Add(NewVar().
			WithName("b").
			WithType(NewTypeRef().WithName("int")))
	file.AddVarList(list)
	testAst(t, "Test_Var_ListIntUnassigned", []byte(src_Test_Var_ListIntUnassigned), file.file)
}

const src_Test_Var_ListIntAssigned = `package go_ast_dsl_test

var a, b int = 1, 2
`

func Test_Var_ListIntAssigned(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Var_ListIntAssigned.txt")
	list := NewVarList().
		WithType(NewTypeRef().WithName("int")).
		Add(NewVar().
			WithName("a").
			WithValue(NewIntLiteral("1"))).
		Add(NewVar().
			WithName("b").
			WithValue(NewIntLiteral("2")))
	file.AddVarList(list)
	testAst(t, "Test_Var_ListIntAssigned", []byte(src_Test_Var_ListIntAssigned), file.file)
}

const src_Test_Var_Group = `package go_ast_dsl_test

var (
	c	= 1
	d	= 2
)
`

func Test_Var_Group(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_Var_Group.txt")
	group := NewVarGroup().
		AddVar(NewVar().
			WithName("c").
			WithValue(NewIntLiteral("1"))).
		AddVar(NewVar().
			WithName("d").
			WithValue(NewIntLiteral("2")))
	file.AddVarGroup(group)
	testAst(t, "Test_Var_Group", []byte(src_Test_Var_Group), file.file)
}

const src_Test_Var_Group2 = `package go_ast_dsl_test

var (
	c	string	= "1"
	d, e	int	= 2, 3
)
`

func Test_Var_Group2(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_Var_Group2.txt")
	group := NewVarGroup().
		AddVar(NewVar().
			WithName("c").
			WithType(NewTypeRef().WithName("string")).
			WithValue(NewIntLiteral("\"1\""))).
		AddList(NewVarList().
			WithType(NewTypeRef().WithName("int")).
			Add(NewVar().
				WithName("d").
				WithValue(NewIntLiteral("2"))).
			Add(NewVar().
				WithName("e").
				WithValue(NewIntLiteral("3"))))
	file.AddVarGroup(group)
	testAst(t, "Test_Var_Group2", []byte(src_Test_Var_Group2), file.file)
}
