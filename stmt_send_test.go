package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Send = `package go_ast_dsl_test

func a() {
	c <- 1
}
`

func Test_Func_Send(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Send.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewSendStmt().
				WithChanExpr(NewTypeRef().WithName("c")).
				WithValExpr(NewIntLiteral("1")))))
	testAst(t, "Test_Func_Send", []byte(src_Test_Func_Send), file.file)
}
