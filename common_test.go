package go_ast_dsl

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"testing"
)

const TEST_DIR = "test"

func loadAstFromFile(filename string) (src []byte, ast *ast.File, fset *token.FileSet, err error) {
	filename = TEST_DIR + "/" + filename + ".txt"

	src, err = ioutil.ReadFile(filename)

	if err == nil {
		ast, fset, err = parseAst(filename, src)
		//fset = token.NewFileSet() // positions are relative to fset
		//
		//ast, err = parser.ParseFile(fset, filename, src, 0 /*mode*/)
	}
	return

	//	return src, ast, fset, err
}

func parseAst(filename string, src []byte) (ast *ast.File, fset *token.FileSet, err error) {
	fset = token.NewFileSet() // positions are relative to fset
	ast, err = parser.ParseFile(fset, filename, src, parser.ParseComments)
	return
}

func testAst(t *testing.T, name string, src []byte, testAst *ast.File) {
	goAst, fset, err := parseAst(name, src)
	if err != nil {
		t.Fatal(err)
	}
	var goSrc bytes.Buffer
	printer.Fprint(&goSrc, fset, goAst)
	if bytes.Compare(src, goSrc.Bytes()) != 0 {
		ast.Print(fset, goAst)
		fmt.Println("----------")
		fmt.Println(goSrc.String())
		t.Fatalf("Error testing '%s': Src from Go AST does not match input src", name)
	}
	fset = token.NewFileSet()
	var testSrc bytes.Buffer
	printer.Fprint(&testSrc, fset, testAst)
	if bytes.Compare(src, testSrc.Bytes()) != 0 {
		ast.Print(fset, testAst)
		fmt.Println("----------")
		fmt.Println(testSrc.String())
		t.Fatalf("Error testing '%s': Src from Test AST does not match input src", name)
	}
}
