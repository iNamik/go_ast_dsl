package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_GotoBack = `package go_ast_dsl_test

func a() {
lbl:
	;
	goto lbl
}
`

func Test_Func_GotoBack(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_GotoBack.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
	WithFunctionBody(NewBlockStmt().
			Add(NewLabelStmt().WithName("lbl")).
	Add(NewGotoStmt().WithLabel("lbl"))))
	testAst(t, "Test_Func_GotoBack", []byte(src_Test_Func_GotoBack), file.file)
}

const src_Test_Func_GotoForward = `package go_ast_dsl_test

func a() {
	goto lbl
lbl:
}
`

func Test_Func_GotoForward(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
	NewFile("Test_Func_GotoForward.txt").
	AddFunction(NewFunction().
	WithFunctionName("a").
	WithFunctionBody(NewBlockStmt().
	Add(NewGotoStmt().WithLabel("lbl")).
	Add(NewLabelStmt().WithName("lbl"))))
	testAst(t, "Test_Func_GotoForward", []byte(src_Test_Func_GotoForward), file.file)
}

