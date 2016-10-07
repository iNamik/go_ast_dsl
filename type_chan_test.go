package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_Chan = `package go_ast_dsl_test

type a chan int
`

func Test_Type_Chan(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Chan.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewChanType().
				WithSendAndReceive().
				WithType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Type_Chan", []byte(src_Test_Type_Chan), file.file)
}

const src_Test_Type_ChanSend = `package go_ast_dsl_test

type a chan<- int
`

func Test_Type_ChanSend(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_ChanSend.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewChanType().
				WithSend().
				WithType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Type_ChanSend", []byte(src_Test_Type_ChanSend), file.file)
}

const src_Test_Type_ChanRecv = `package go_ast_dsl_test

type a <-chan int
`

func Test_Type_ChanRecv(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_ChanRecv.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewChanType().
				WithRecv().
				WithType(NewTypeRef().WithName("int"))))
	testAst(t, "Test_Type_ChanRecv", []byte(src_Test_Type_ChanRecv), file.file)
}
