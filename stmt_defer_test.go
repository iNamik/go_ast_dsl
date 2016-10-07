package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Defer = `package go_ast_dsl_test

func a() {
	defer b()
}
`

func Test_Func_Defer(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_Defer.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").
	WithFunctionBody(NewBlockStmt().
	Add(NewDeferStmt().WithCallExpr(NewCallExpr().
	WithFunction(NewTypeRef().WithName("b"))))))
	testAst(t, "Test_Func_Defer", []byte(src_Test_Func_Defer), file.file)
}

