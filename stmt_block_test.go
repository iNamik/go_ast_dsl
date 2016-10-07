package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_EmptyBlock = `package go_ast_dsl_test

func a() {
}
`

func Test_Func_EmptyBlock(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_EmptyBlock.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt()))
	testAst(t, "Test_Func_EmptyBlock", []byte(src_Test_Func_EmptyBlock), file.file)
}
