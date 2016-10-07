package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Go = `package go_ast_dsl_test

func a() {
	go b()
}
`

func Test_Func_Go(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_Go.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").
	WithFunctionBody(NewBlockStmt().
	Add(NewGoStmt().WithCallExpr(NewCallExpr().
	WithFunction(NewTypeRef().WithName("b"))))))
	testAst(t, "Test_Func_Go", []byte(src_Test_Func_Go), file.file)
}

