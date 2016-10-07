package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Empty = `package go_ast_dsl_test

func a() {
}
`

func Test_Func_Empty(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Empty.txt").
		AddFunction(NewFunction().
			WithFunctionName("a"))
	testAst(t, "Test_Func_Empty", []byte(src_Test_Func_Empty), file.file)
}
