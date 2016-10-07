package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Recv = `package go_ast_dsl_test

func a() {
	b := <-c
}
`

func Test_Func_Recv(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Recv.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewAssignStmt().
				WithLhsExpr(NewTypeRef().WithName("b")).
				WithOperator(ASSIGN_DEFINE).
				WithRhsExpr(NewUnaryExpr().
					WithExpr(NewTypeRef().WithName("c")).
					WithOperator(UNARY_RECV)))))
	testAst(t, "Test_Func_Recv", []byte(src_Test_Func_Recv), file.file)
}

const src_Test_Func_Add = `package go_ast_dsl_test

func a() {
	b := c + 1
}
`

func Test_Func_Add(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_Add.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").WithFunctionBody(NewBlockStmt().
	Add(NewAssignStmt().
	WithLhsExpr(NewTypeRef().WithName("b")).
	WithOperator(ASSIGN_DEFINE).
	WithRhsExpr(NewBinaryExpr().
	WithLeftExpr(NewTypeRef().WithName("c")).
	WithRightExpr(NewIntLiteral("1")).
	WithOperator(BINARY_ADD)))))
	testAst(t, "Test_Func_Add", []byte(src_Test_Func_Add), file.file)
}

const src_Test_Func_CallNoParams = `package go_ast_dsl_test

func a() {
	c()
}
`

func Test_Func_CallNoParams(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_CallNoParams.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").
	WithFunctionBody(NewBlockStmt().
	Add(NewExprStmt().
	WithExpr(NewCallExpr().
	WithFunction(NewTypeRef().WithName("c"))))))
	testAst(t, "Test_Func_CallNoParams", []byte(src_Test_Func_CallNoParams), file.file)
}

const src_Test_Func_Call1Param = `package go_ast_dsl_test

func a() {
	c(1)
}
`

func Test_Func_Call1Param(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_Call1Param.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").
	WithFunctionBody(NewBlockStmt().
	Add(NewExprStmt().
	WithExpr(NewCallExpr().
	WithFunction(NewTypeRef().WithName("c")).
	AddArg(NewIntLiteral("1"))))))
	testAst(t, "Test_Func_Call1Param", []byte(src_Test_Func_Call1Param), file.file)
}

const src_Test_Func_CallEllipses = `package go_ast_dsl_test

func a() {
	c(d...)
}
`

func Test_Func_CallEllipses(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_CallEllipses.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").WithFunctionBody(NewBlockStmt().
	Add(NewExprStmt().
	WithExpr(NewCallExpr().
	WithFunction(NewTypeRef().WithName("c")).
	AddArgWithEllipses(NewTypeRef().WithName("d"))))))
	testAst(t, "Test_Func_CallEllipses", []byte(src_Test_Func_CallEllipses), file.file)
}
