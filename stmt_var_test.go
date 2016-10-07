package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Var = `package go_ast_dsl_test

func a() {
	var b int = 1
}
`

func Test_Func_Var(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Var.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewVarStmt().
					WithVar(NewVar().
						WithName("b").
						WithType(NewTypeRef().WithName("int")).
						WithValue(NewIntLiteral("1"))))))
	testAst(t, "Test_Func_Var", []byte(src_Test_Func_Var), file.file)
}

const src_Test_Func_VarList = `package go_ast_dsl_test

func a() {
	var b, c int = 1, 2
}
`

func Test_Func_VarList(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_VarList.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewVarStmt().
					WithVarList(NewVarList().
						WithType(NewTypeRef().WithName("int")).
						Add(NewVar().
							WithName("b").
							WithValue(NewIntLiteral("1"))).
						Add(NewVar().
							WithName("c").
							WithValue(NewIntLiteral("2")))))))
	testAst(t, "Test_Func_VarList", []byte(src_Test_Func_VarList), file.file)
}

const src_Test_Func_VarGroup = `package go_ast_dsl_test

func a() {
	var (
		c	string	= "1"
		d, e	int	= 2, 3
	)
}
`

func Test_Func_VarGroup(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_VarGroup.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewVarStmt().
					WithVarGroup(NewVarGroup().
						AddVar(NewVar().
							WithName("c").
							WithType(NewTypeRef().WithName("string")).
							WithValue(NewIntLiteral("\"1\""))).
						AddList(NewVarList().
							WithType(NewTypeRef().WithName("int")).
							Add(NewVar().
								WithName("d").
								WithValue(NewIntLiteral("2"))).
							Add(NewVar().
								WithName("e").
								WithValue(NewIntLiteral("3"))))))))
	testAst(t, "Test_Func_VarGroup", []byte(src_Test_Func_VarGroup), file.file)
}
