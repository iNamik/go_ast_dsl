package go_ast_dsl

import "go/ast"

func NewPackage(name string) *Package {
	pkg := &ast.Package{Name: name, Scope: ast.NewScope(nil), Files: make(map[string]*ast.File)}
	ident := &ast.Ident{Name: name}
	return &Package{Package: pkg, Ident: ident}
}

type Package struct {
	Package *ast.Package
	Ident   *ast.Ident
	Files   []*File
}

func (p *Package) NewFile(name string) *File {
	file := newFile(p.Ident /*p.Package.Scope*/, nil)
	p.Package.Files[name] = file.file
	p.Files = append(p.Files, file)
	return file
}

func (p *Package) Ast() *ast.Package {
	return p.Package
}
