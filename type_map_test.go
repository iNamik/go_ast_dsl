package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_Map = `package go_ast_dsl_test

type a map[string]int
`

func Test_Type_Map(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Map.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewMapType().
				WithKeyType(NewTypeRef().WithName("string")).
				WithValueType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Type_Map", []byte(src_Test_Type_Map), file.file)
}
