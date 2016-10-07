package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_SimpleReturn = `package go_ast_dsl_test

func a() {
	return
}
`

func Test_Func_SimpleReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_SimpleReturn.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewReturnStmt())))
	testAst(t, "Test_Func_SimpleReturn", []byte(src_Test_Func_SimpleReturn), file.file)
}

const src_Test_Func_Return1Result = `package go_ast_dsl_test

func a() {
	return 1
}
`

func Test_Func_Return1Result(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Return1Result.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewReturnStmt().AddResult(NewIntLiteral("1")))))
	testAst(t, "Test_Func_Return1Result", []byte(src_Test_Func_Return1Result), file.file)
}

const src_Test_Func_Return2Results = `package go_ast_dsl_test

func a() {
	return 1, "two"
}
`

func Test_Func_Return2Results(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Return2Results.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewReturnStmt().
				AddResult(NewIntLiteral("1")).
				AddResult(NewStringLiteral(`"two"`)))))
	testAst(t, "Test_Func_Return2Results", []byte(src_Test_Func_Return2Results), file.file)
}
