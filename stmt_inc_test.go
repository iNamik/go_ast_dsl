package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Inc = `package go_ast_dsl_test

func a() {
	b++
}
`

func Test_Func_Inc(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Inc.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewIncStmt().
				WithExpr(NewTypeRef().WithName("b")))))
	testAst(t, "Test_Func_Inc", []byte(src_Test_Func_Inc), file.file)
}
