package go_ast_dsl

import (
	"go/ast"
)

func newFile(ident *ast.Ident, scope *ast.Scope) *File {
	file := &ast.File{
		Name:  ident,
		Scope: ast.NewScope(scope),
	}
	return &File{file: file}
}

type File struct {
	file *ast.File
}

func (f *File) Import(i *Import) *File {
	f.importDecl(i.Ast)
	return f
}

func (f *File) ImportGroup(ig *ImportGroup) *File {
	f.importDecl(ig.decl)
	return f
}

// See:
//    parser/declare
func (f *File) AddConst(c *ConstDecl) *File {
	spec := c.ValueSpec()
	f.declareConst(spec, c.iota)
	decl := c.GenDecl()
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddConstList(l *ConstDeclList) *File {
	spec := l.ValueSpec()
	f.declareConst(spec, 0) // TODO iota
	decl := l.GenDecl()
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddConstGroup(g *ConstDeclGroup) *File {
	decl := g.GenDecl()
	for _, spec := range decl.Specs {
		f.declareConst(spec.(*ast.ValueSpec), 0) // TODO iota
	}
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddVar(v *Var) *File {
	spec := v.ValueSpec()
	f.declareVar(spec, 0) // TODO iota
	decl := v.GenDecl()
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddVarList(l *VarList) *File {
	spec := l.ValueSpec()
	f.declareVar(spec, 0) // TODO iota
	decl := l.GenDecl()
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddVarGroup(g *VarGroup) *File {
	decl := g.GenDecl()
	for _, spec := range decl.Specs {
		f.declareVar(spec.(*ast.ValueSpec), 0) // TODO iota
	}
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddType(t *TypeDef) *File {
	decl := t.GenDecl()
	for _, spec := range decl.Specs {
		f.declareType(spec.(*ast.TypeSpec))
	}
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddTypeGroup(g *TypeDefGroup) *File {
	decl := g.GenDecl()
	for _, spec := range decl.Specs {
		f.declareType(spec.(*ast.TypeSpec))
	}
	f.file.Decls = append(f.file.Decls, decl)
	return f
}

func (f *File) AddFunction(function *Function) *File {
	f.file.Decls = append(f.file.Decls, function.Decl())
	return f
}

/*********************************************************************
 * Private
 *********************************************************************/

func (f *File) importDecl(decl *ast.GenDecl) {
	f.file.Decls = append(f.file.Decls, decl)
	f.addImports(decl.Specs)
	//	f.addComment(decl.Doc)
}

func (f *File) addImports(specs []ast.Spec) {
	for _, spec := range specs {
		ispec := spec.(*ast.ImportSpec)
		f.file.Imports = append(f.file.Imports, ispec)
	}
}

func (f *File) addComment(comment *ast.CommentGroup) {
	if nil != comment && len(comment.List) > 0 {
		f.file.Comments = append(f.file.Comments, comment)
	}
}

// see parser/parseValueSpec
// 	p.declare(*ValueSpec, iota, p.topScope, ast.Con, idents...)
// see parser/declare
//	declare(decl, data interface{}, scope *ast.Scope, kind ast.ObjKind, idents ...*ast.Ident)
//
func (f *File) declareConst(spec *ast.ValueSpec, iota int) {
	for _, ident := range spec.Names {
		// TODO Check if object already assigned
		//
		obj := ast.NewObj(ast.Con, ident.Name)
		obj.Decl = spec
		obj.Data = iota
		ident.Obj = obj

		// Skip '_'
		if "_" != ident.Name {
			// TODO Check return value for existing object
			//
			_ = f.file.Scope.Insert(obj)
		}
	}
}

func (f *File) declareVar(spec *ast.ValueSpec, iota int) {
	for _, ident := range spec.Names {
		// TODO Check if object already assigned
		//
		obj := ast.NewObj(ast.Var, ident.Name)
		obj.Decl = spec
		obj.Data = iota
		ident.Obj = obj

		// Skip '_'
		if "_" != ident.Name {
			// TODO Check return value for existing object
			//
			_ = f.file.Scope.Insert(obj)
		}
	}
}

func (f *File) declareType(spec *ast.TypeSpec) {
	// 	p.declare(spec, nil, p.topScope, ast.Typ, ident)
	ident := spec.Name
	// TODO Check if object already assigned
	//
	obj := ast.NewObj(ast.Var, ident.Name)
	obj.Decl = spec
	obj.Data = nil
	ident.Obj = obj

	// Skip '_'
	if "_" != ident.Name {
		// TODO Check return value for existing object
		//
		_ = f.file.Scope.Insert(obj)
	}
}
