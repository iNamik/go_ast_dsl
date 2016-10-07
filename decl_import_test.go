package go_ast_dsl

import (
	"testing"
)

const src_Test_SingleImport = `package go_ast_dsl_test

import "one"
`

func Test_SingleImport(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("Test_SingleImport.go").
		Import(NewImport("one"))
	testAst(t, "Test_SingleImport", []byte(src_Test_SingleImport), file.file)
}

//const src_Test_SingleImportWithLeadComment = `package go_ast_dsl_test
//
//// Lead Comment
//import "one"
//`
//
//func Test_SingleImportWithLeadComment(t *testing.T) {
//	file := NewPackage("go_ast_dsl_test").
//	NewFile("Test_SingleImportWithLeadComment.go").
//	Import(NewImport("one").WithLeadComment(NewCommentGroup().addLine(" Lead Comment")))
//	testAst(t, "Test_SingleImportWithLeadComment", []byte(src_Test_SingleImportWithLeadComment), file.Ast)
//}

//const src_Test_SingleImportWithLineComment = `package go_ast_dsl_test
//
//import "one"	// Line comment
//`
//
//func Test_SingleImportWithLineComment(t *testing.T) {
//	file := NewPackage("go_ast_dsl_test").
//	NewFile("Test_SingleImportWithLineComment.go").
//	Import(NewImport("one").WithLineComment(NewCommentGroup().addLine(" Line comment")))
//	testAst(t, "Test_SingleImportWithLineComment", []byte(src_Test_SingleImportWithLineComment), file.Ast)
//}

const src_Test_SingleImportWithAlias = `package go_ast_dsl_test

import alias "one"
`

func Test_SingleImportWithAlias(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_SingleImportWithAlias.go").
		Import(NewImport("one").WithAlias("alias"))
	testAst(t, "Test_SingleImportWithAlias", []byte(src_Test_SingleImportWithAlias), file.file)
}

const src_Test_SingleImportWithDot = `package go_ast_dsl_test

import . "one"
`

func Test_SingleImportWithDot(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_SingleImportWithDot.go").
		Import(NewImport("one").WithAlias("."))
	testAst(t, "Test_SingleImportWithDot", []byte(src_Test_SingleImportWithDot), file.file)
}

const src_Test_MultipleSingleImport = `package go_ast_dsl_test

import "one"
import "two"
`

func Test_MultipleSingleImport(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_MultipleSingleImport.go").
		Import(NewImport("one")).
		Import(NewImport("two"))
	testAst(t, "Test_MultipleSingleImport", []byte(src_Test_MultipleSingleImport), file.file)
}

const src_Test_ImportGroup = `package go_ast_dsl_test

import (
	"one"
	"two"
)
`

func Test_ImportGroup(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_ImportGroup.go").
		ImportGroup(NewImportGroup().
			Import("one").
			Import("two"))
	testAst(t, "Test_ImportGroup", []byte(src_Test_ImportGroup), file.file)
}

const src_Test_MultipleImportGroup = `package go_ast_dsl_test

import (
	"one"
	"two"
)
import (
	"three"
	"four"
)
`

func Test_MultipleImportGroup(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("src_Test_MultipleImportGroup.go").
		ImportGroup(NewImportGroup().
			Import("one").
			Import("two")).
		ImportGroup(NewImportGroup().
			Import("three").
			Import("four"))
	testAst(t, "Test_MultipleImportGroup", []byte(src_Test_MultipleImportGroup), file.file)
}
