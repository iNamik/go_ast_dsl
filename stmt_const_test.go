package go_ast_dsl

import (
	"testing"
)

const src_Test_Func_Const = `package go_ast_dsl_test

func a() {
	const A = 1
}
`

func Test_Func_Const(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_Const.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewConstStmt().
					WithConst(NewConstDecl().WithName("A").WithValue(NewIntLiteral("1"))))))
	testAst(t, "Test_Func_Const", []byte(src_Test_Func_Const), file.file)
}

const src_Test_Func_ConstList = `package go_ast_dsl_test

func a() {
	const A, B = 1, 2
}
`

func Test_Func_ConstList(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_ConstList.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewConstStmt().
					WithConstList(NewConstDeclList().
						Add(NewConstDecl().
							WithName("A").
							WithValue(NewIntLiteral("1")).
							WithIota(0)).
						Add(NewConstDecl().
							WithName("B").
							WithValue(NewIntLiteral("2")).
							WithIota(1))))))
	testAst(t, "Test_Func_ConstList", []byte(src_Test_Func_ConstList), file.file)
}

const src_Test_Func_ConstGroup = `package go_ast_dsl_test

func a() {
	const (
		A	= 1
		B	= 2
	)
}
`

func Test_Func_ConstGroup(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Func_ConstGroup.txt").
		AddFunction(NewFunction().
			WithFunctionName("a").
			WithFunctionBody(NewBlockStmt().
				Add(NewConstStmt().
					WithConstGroup(NewConstDeclGroup().
						AddConst(NewConstDecl().
							WithName("A").
							WithValue(NewIntLiteral("1")).
							WithIota(0)).
						AddConst(NewConstDecl().
							WithName("B").
							WithValue(NewIntLiteral("2")).
							WithIota(1))))))
	testAst(t, "Test_Func_ConstGroup", []byte(src_Test_Func_ConstGroup), file.file)
}
