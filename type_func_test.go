package go_ast_dsl

import (
	"testing"
)

const src_Test_Type_FuncEmpty = `package go_ast_dsl_test

type a func()
`

func Test_Type_FuncEmpty(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_FuncEmpty.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType()))
	testAst(t, "Test_Type_FuncEmpty", []byte(src_Test_Type_FuncEmpty), file.file)
}

const src_Test_Type_Func1ParamNoReturn = `package go_ast_dsl_test

type a func(int)
`

func Test_Type_Func1ParamNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Func1ParamNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType().
				WithAnonParams(NewAnonMethodParamList().
					AddParam(NewAnonMethodParam().
						WithType(NewTypeRef().WithName("int"))))))
	testAst(t, "Test_Type_Func1ParamNoReturn", []byte(src_Test_Type_Func1ParamNoReturn), file.file)
}

const src_Test_Type_Func1Param1Result = `package go_ast_dsl_test

type a func(int) int
`

func Test_Type_Func1Param1Result(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Func1Param1Result.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType().
				WithAnonParams(NewAnonMethodParamList().
					AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("int")))).
				WithAnonResults(NewAnonMethodResultList().
					AddResult(NewAnonMethodResult().WithType(NewTypeRef().WithName("int"))))))
	testAst(t, "Test_Type_Func1Param1Result", []byte(src_Test_Type_Func1Param1Result), file.file)
}

const src_Test_Type_Func1NamedParamNoReturn = `package go_ast_dsl_test

type a func(c int)
`

func Test_Type_Func1NamedParamNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Func1NamedParamNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType().
				WithNamedParams(NewNamedMethodParamList().
					AddParam(NewNamedMethodParam().
						WithType(NewTypeRef().WithName("int")).
						WithName("c")))))
	testAst(t, "Test_Type_Func1NamedParamNoReturn", []byte(src_Test_Type_Func1NamedParamNoReturn), file.file)
}

const src_Test_Type_Func1NamedParamListNoReturn = `package go_ast_dsl_test

type a func(c, d int)
`

func Test_Type_Func1NamedParamListNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Func1NamedParamListNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType().
				WithNamedParams(NewNamedMethodParamList().
					AddParams(NewNamedMethodParams().
						WithType(NewTypeRef().WithName("int")).
						AddName("c").
						AddName("d")))))
	testAst(t, "Test_Type_Func1NamedParamListNoReturn", []byte(src_Test_Type_Func1NamedParamListNoReturn), file.file)
}

const src_Test_Type_Func2ParamsNoReturn = `package go_ast_dsl_test

type a func(int, string)
`

func Test_Type_Func2ParamsNoReturn(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Func2ParamsNoReturn.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType().
				WithAnonParams(NewAnonMethodParamList().
					AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("int"))).
					AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("string"))))))
	testAst(t, "Test_Type_Func2ParamsNoReturn", []byte(src_Test_Type_Func2ParamsNoReturn), file.file)
}

const src_Test_Type_Func1Param1NamedResult = `package go_ast_dsl_test

type a func(int) (c int)
`

func Test_Type_Func1Param1NamedResult(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_Type_Func1Param1NamedResult.txt").
		AddType(NewTypeDef().
			WithName("a").
			WithType(NewFuncType().
				WithAnonParams(NewAnonMethodParamList().
					AddParam(NewAnonMethodParam().WithType(NewTypeRef().WithName("int")))).
				WithNamedResults(NewNamedMethodResultList().
					AddResult(NewNamedMethodResult().WithName("c").
						WithType(NewTypeRef().WithName("int"))))))
	testAst(t, "Test_Type_Func1Param1NamedResult", []byte(src_Test_Type_Func1Param1NamedResult), file.file)
}
