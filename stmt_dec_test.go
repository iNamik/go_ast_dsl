package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Dec = `package go_ast_dsl_test

func a() {
	b--
}
`

func Test_Func_Dec(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Dec.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewDecStmt().
				WithExpr(NewTypeRef().WithName("b")))))
	testAst(t, "Test_Func_Dec", []byte(src_Test_Func_Dec), file.file)
}
