package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Type = `package go_ast_dsl_test

func a() {
	type b int
}
`

func Test_Func_Type(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Type.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewTypeStmt().
					WithType(NewTypeDef().
						WithName("b").
						WithType(NewTypeRef().WithName("int"))))))
	testAst(t, "Test_Func_Type", []byte(src_Test_Func_Type), file.file)
}

const src_Test_FuncTypeGroup = `package go_ast_dsl_test

func a() {
	type (
		b	int
		c	string
	)
}
`

func Test_FuncTypeGroup(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_FuncTypeGroup.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewTypeStmt().
					WithTypeGroup(NewTypeDefGroup().
						AddType(NewTypeDef().
							WithName("b").
							WithType(NewTypeRef().WithName("int"))).
						AddType(NewTypeDef().
							WithName("c").
							WithType(NewTypeRef().WithName("string")))))))
	testAst(t, "Test_FuncTypeGroup", []byte(src_Test_FuncTypeGroup), file.file)
}
