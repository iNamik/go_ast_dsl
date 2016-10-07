package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_Array = `package go_ast_dsl_test

type a []int
`

func Test_Type_Array(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Array.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewArrayType().
				WithType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Type_Array", []byte(src_Test_Type_Array), file.file)
}

const src_Test_Type_ArrayIntLen = `package go_ast_dsl_test

type a [1]int
`

func Test_Type_ArrayIntLen(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_ArrayIntLen.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewArrayType().
				WithLen(NewIntLiteral("1")).
				WithType(NewTypeRef().
					WithName("int"))))
	testAst(t, "Test_Type_ArrayIntLen", []byte(src_Test_Type_ArrayIntLen), file.file)
}
