package go_ast_dsl

import (
	"testing"
)

const src_Test_Const = `package go_ast_dsl_test

const c = 1
`

func Test_Const(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Const.txt")
	con := NewConstDecl().
		WithName("c").
		WithValue(NewIntLiteral("1")).
		WithIota(0)
	file.AddConst(con)
	testAst(t, "Test_Const", []byte(src_Test_Const), file.file)
}

//const src_Test_ConstExpr = `package go_ast_dsl_test
//
//const c = 1 * (2 + 3)
//
//`
//
//func Test_ConstExpr(t *testing.T) {
//	file := NewPackage("go_ast_dsl_test").
//	NewFile("Test_ConstExpr.txt")
//	//con := NewConst().
//	//WithName("c").
//	//WithIntValue("1").
//	//WithIota(0)
//	//file.AddConst(con)
//	testAst(t, "Test_ConstExpr", []byte(src_Test_ConstExpr), file.file)
//}

const src_Test_Const_WithType = `package go_ast_dsl_test

const c int = 1
`

func Test_Const_WithType(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Const_WithType.txt")
	con := NewConstDecl().
		WithName("c").
		WithType(NewTypeRef().WithName("int")).
		WithValue(NewIntLiteral("1")).
		WithIota(0)
	file.AddConst(con)
	testAst(t, "Test_ConstTyped", []byte(src_Test_Const_WithType), file.file)
}

const src_Test_Const_List = `package go_ast_dsl_test

const c, d = 1, 2
`

func Test_Const_list(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Const_list.txt")
	list := NewConstDeclList().
		Add(NewConstDecl().
			WithName("c").
			WithValue(NewIntLiteral("1")).
			WithIota(0)).
		Add(NewConstDecl().
			WithName("d").
			WithValue(NewIntLiteral("2")).
			WithIota(1))
	file.AddConstList(list)
	testAst(t, "Test_Const_list", []byte(src_Test_Const_List), file.file)
}

const src_Test_Const_ListWithType = `package go_ast_dsl_test

const c, d int = 1, 2
`

func Test_Const_ListWithType(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Const_ListWithType.txt")
	list := NewConstDeclList().
		WithType(NewTypeRef().WithName("int")).
		Add(NewConstDecl().
			WithName("c").
			WithValue(NewIntLiteral("1")).
			WithIota(0)).
		Add(NewConstDecl().
			WithName("d").
			WithValue(NewIntLiteral("2")).
			WithIota(1))
	file.AddConstList(list)
	testAst(t, "Test_Const_ListWithType", []byte(src_Test_Const_ListWithType), file.file)
}

const src_Test_Const_Group = `package go_ast_dsl_test

const (
	c	= 1
	d	= 2
)
`

func Test_Const_Group(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_Const_Group.txt")
	group := NewConstDeclGroup().
		AddConst(NewConstDecl().
			WithName("c").
			WithValue(NewIntLiteral("1")).
			WithIota(0)).
		AddConst(NewConstDecl().
			WithName("d").
			WithValue(NewIntLiteral("2")).
			WithIota(1))
	file.AddConstGroup(group)
	testAst(t, "Test_Const_Group", []byte(src_Test_Const_Group), file.file)
}

const src_Test_Const_Group2 = `package go_ast_dsl_test

const (
	c	= 1
	d, e	= 2, 3
)
`

func Test_Const_Group2(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_Const_Group2.txt")
	group := NewConstDeclGroup().
		AddConst(NewConstDecl().
			WithName("c").
			WithValue(NewIntLiteral("1")).
			WithIota(0)).
		AddList(NewConstDeclList().
			Add(NewConstDecl().
				WithName("d").
				WithValue(NewIntLiteral("2")).
				WithIota(0)).
			Add(NewConstDecl().
				WithName("e").
				WithValue(NewIntLiteral("3")).
				WithIota(1)))
	file.AddConstGroup(group)
	testAst(t, "Test_Const_Group2", []byte(src_Test_Const_Group2), file.file)
}

const src_Test_Const_GroupWithType = `package go_ast_dsl_test

const (
	c	int	= 1
	d	string	= "2"
)
`

func Test_Const_GroupWithType(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_Const_GroupWithType.txt")
	group := NewConstDeclGroup().
		AddConst(NewConstDecl().
			WithName("c").
			WithType(NewTypeRef().WithName("int")).
			WithValue(NewIntLiteral("1")).
			WithIota(0)).
		AddConst(NewConstDecl().
			WithName("d").
			WithType(NewTypeRef().WithName("string")).
			WithValue(NewIntLiteral("\"2\"")).
			WithIota(1))
	file.AddConstGroup(group)
	testAst(t, "Test_Const_GroupWithType", []byte(src_Test_Const_GroupWithType), file.file)
}
