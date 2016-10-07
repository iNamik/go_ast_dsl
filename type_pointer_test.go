package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_Pointer = `package go_ast_dsl_test

type a *int
`

func Test_Type_Pointer(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Pointer.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewPointerType().
				WithType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Type_Pointer", []byte(src_Test_Type_Pointer), file.file)
}

const src_Test_Type_PointerPointer = `package go_ast_dsl_test

type a **int
`

func Test_Type_PointerPointer(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_PointerPointer.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewPointerType().
				WithType(NewPointerType().
					WithType(NewTypeRef().WithName("int")))))
	testAst(t, "Test_Type_PointerPointer", []byte(src_Test_Type_PointerPointer), file.file)
}

const src_Test_Type_ArrayPointer = `package go_ast_dsl_test

type a []*int
`

func Test_Type_ArrayPointer(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_ArrayPointer.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewArrayType().
				WithType(NewPointerType().
					WithType(NewTypeRef().WithName("int")))))
	testAst(t, "Test_Type_ArrayPointer", []byte(src_Test_Type_ArrayPointer), file.file)
}
