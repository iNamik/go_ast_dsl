package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_InterfaceEmpty = `package go_ast_dsl_test

type a interface {
}
`

func Test_Type_InterfaceEmpty(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceEmpty.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType()))
	testAst(t, "Test_Type_InterfaceEmpty", []byte(src_Test_Type_InterfaceEmpty), file.file)
}

const src_Test_Type_InterfaceEmbedded = `package go_ast_dsl_test

type a interface {
	myType
}
`

func Test_Type_InterfaceEmbedded(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceEmbedded.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddEmbedded(NewInterfaceEmbedField().
					WithType(NewTypeRef().WithName("myType")))))
	testAst(t, "Test_Type_InterfaceEmbedded", []byte(src_Test_Type_InterfaceEmbedded), file.file)
}

const src_Test_Type_InterfaceMethod1ParamNoReturn = `package go_ast_dsl_test

type a interface {
	b(int)
}
`

func Test_Type_InterfaceMethod1ParamNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceMethod1ParamNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddMethod(NewInterfaceMethodField().
					WithMethodName("b").
					WithAnonParams(NewAnonMethodParamList().
						AddParam(NewAnonMethodParam().
							WithType(NewTypeRef().WithName("int")))))))
	testAst(t, "Test_Type_InterfaceMethod1ParamNoReturn", []byte(src_Test_Type_InterfaceMethod1ParamNoReturn), file.file)
}

const src_Test_Type_InterfaceMethod1Param1Result = `package go_ast_dsl_test

type a interface {
	b(int) int
}
`

func Test_Type_InterfaceMethod1Param1Result(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceMethod1Param1Result.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddMethod(NewInterfaceMethodField().
					WithMethodName("b").
					WithAnonParams(NewAnonMethodParamList().
						AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("int")))).
					WithAnonResults(NewAnonMethodResultList().
						AddResult(NewAnonMethodResult().WithType(NewTypeRef().WithName("int")))))))
	testAst(t, "Test_Type_InterfaceMethod1Param1Result", []byte(src_Test_Type_InterfaceMethod1Param1Result), file.file)
}

const src_Test_Type_InterfaceMethod1NamedParamNoReturn = `package go_ast_dsl_test

type a interface {
	b(c int)
}
`

func Test_Type_InterfaceMethod1NamedParamNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceMethod1NamedParamNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddMethod(NewInterfaceMethodField().
					WithMethodName("b").
					WithNamedParams(NewNamedMethodParamList().
						AddParam(NewNamedMethodParam().
							WithType(NewTypeRef().WithName("int")).
							WithName("c"))))))
	testAst(t, "Test_Type_InterfaceMethod1NamedParamNoReturn", []byte(src_Test_Type_InterfaceMethod1NamedParamNoReturn), file.file)
}

const src_Test_Type_InterfaceMethod1NamedParamListNoReturn = `package go_ast_dsl_test

type a interface {
	b(c, d int)
}
`

func Test_Type_InterfaceMethod1NamedParamListNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceMethod1NamedParamListNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddMethod(NewInterfaceMethodField().
					WithMethodName("b").
					WithNamedParams(NewNamedMethodParamList().
						AddParams(NewNamedMethodParams().
							WithType(NewTypeRef().WithName("int")).
							AddName("c").
							AddName("d"))))))
	testAst(t, "Test_Type_InterfaceMethod1NamedParamListNoReturn", []byte(src_Test_Type_InterfaceMethod1NamedParamListNoReturn), file.file)
}

const src_Test_Type_InterfaceMethod2ParamsNoReturn = `package go_ast_dsl_test

type a interface {
	b(int, string)
}
`

func Test_Type_InterfaceMethod2ParamsNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceMethod2ParamsNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddMethod(NewInterfaceMethodField().
					WithMethodName("b").
					WithAnonParams(NewAnonMethodParamList().
						AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("int"))).
						AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("string")))))))
	testAst(t, "Test_Type_InterfaceMethod2ParamsNoReturn", []byte(src_Test_Type_InterfaceMethod2ParamsNoReturn), file.file)
}

const src_Test_Type_InterfaceMethod1Param1NamedResult = `package go_ast_dsl_test

type a interface {
	b(int) (c int)
}
`

func Test_Type_InterfaceMethod1Param1NamedResult(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_InterfaceMethod1Param1NamedResult.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewInterfaceType().
				AddMethod(NewInterfaceMethodField().
					WithMethodName("b").
					WithAnonParams(NewAnonMethodParamList().
						AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("int")))).
					WithNamedResults(NewNamedMethodResultList().
						AddResult(NewNamedMethodResult().WithName("c").
							WithType(NewTypeRef().WithName("int")))))))
	testAst(t, "Test_Type_InterfaceMethod1Param1NamedResult", []byte(src_Test_Type_InterfaceMethod1Param1NamedResult), file.file)
}
