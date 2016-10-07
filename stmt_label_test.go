package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Label = `package go_ast_dsl_test

func a() {
lbl:
}
`

func Test_Func_Label(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Label.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewLabelStmt().WithName("lbl"))))
	testAst(t, "Test_Func_Label", []byte(src_Test_Func_Label), file.file)
}

// TODO Go's code generator currently forces a new line after labels
const src_Test_Func_LabelWithReturn = `package go_ast_dsl_test

func a() {
lbl:
	return
}
`

func Test_Func_LabelWithReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_LabelWithReturn.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").WithFunctionBody(NewBlockStmt().
			Add(NewLabelStmt().
				WithName("lbl").
				WithStmt(NewReturnStmt()))))
	testAst(t, "Test_Func_LabelWithReturn", []byte(src_Test_Func_LabelWithReturn), file.file)
}
