package go_ast_dsl

import (
	"testing"
)

const src_Test_Package = `package go_ast_dsl_test
`

func Test_Package(t *testing.T) {
	file := NewPackage("go_ast_dsl_test").
		NewFile("package.txt")
	testAst(t, "Test_Package", []byte(src_Test_Package), file.file)
}
