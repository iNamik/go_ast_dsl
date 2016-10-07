package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Assign = `package go_ast_dsl_test

func a() {
	b = 1
}
`

func Test_Func_Assign(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Assign.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewAssignStmt().
				WithLhsExpr(NewTypeRef().WithName("b")).
				WithRhsExpr(NewIntLiteral("1")))))
	testAst(t, "Test_Func_Assign", []byte(src_Test_Func_Assign), file.file)
}

const src_Test_Func_Define = `package go_ast_dsl_test

func a() {
	b := 1
}
`

func Test_Func_Define(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Define.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewAssignStmt().
				WithLhsExpr(NewTypeRef().WithName("b")).
				WithOperator(ASSIGN_DEFINE).
				WithRhsExpr(NewIntLiteral("1")))))
	testAst(t, "Test_Func_Define", []byte(src_Test_Func_Define), file.file)
}

const src_Test_Func_PlusEquals = `package go_ast_dsl_test

func a() {
	b += 1
}
`

func Test_Func_PlusEquals(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_PlusEquals.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewAssignStmt().
				WithLhsExpr(NewTypeRef().WithName("b")).
				WithRhsExpr(NewIntLiteral("1")).
				WithOperator(ASSIGN_ADD))))
	testAst(t, "Test_Func_PlusEquals", []byte(src_Test_Func_PlusEquals), file.file)
}

const src_Test_Func_AssignList = `package go_ast_dsl_test

func a() {
	b, c = 1, 2
}
`

func Test_Func_AssignList(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_AssignList.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewAssignListStmt().
				AddLhsExpr(NewTypeRef().WithName("b")).
				AddLhsExpr(NewTypeRef().WithName("c")).
				AddRhsExpr(NewIntLiteral("1")).
				AddRhsExpr(NewIntLiteral("2")))))
	testAst(t, "Test_Func_AssignList", []byte(src_Test_Func_AssignList), file.file)
}

const src_Test_Func_DefineList = `package go_ast_dsl_test

func a() {
	b, c := 1, 2
}
`

func Test_Func_DefineList(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_DefineList.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewAssignListStmt().
				AddLhsExpr(NewTypeRef().WithName("b")).
				AddLhsExpr(NewTypeRef().WithName("c")).
				WithOperator(ASSIGN_DEFINE).
				AddRhsExpr(NewIntLiteral("1")).
				AddRhsExpr(NewIntLiteral("2")))))
	testAst(t, "Test_Func_DefineList", []byte(src_Test_Func_DefineList), file.file)
}
