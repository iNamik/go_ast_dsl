package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_Paren = `package go_ast_dsl_test

type a (*int)
`

func Test_Type_Paren(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Paren.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewParenType().
				WithType(NewPointerType().
					WithType(NewTypeRef().WithName("int")))))
	testAst(t, "Test_Type_Paren", []byte(src_Test_Type_Paren), file.file)
}
