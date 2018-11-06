# AST Notes

## Object

### ObjKind (ast/scope:139)

    // ObjKind describes what an object represents.
    type ObjKind int
    
    // The list of possible Object kinds.
    const (
        Bad ObjKind = iota // for error handling
        Pkg                // package
        Con                // constant
        Typ                // type
        Var                // variable
        Fun                // function or method
        Lbl                // label
    )
    
    var objKindStrings = [...]string{
        Bad: "bad",
        Pkg: "package",
        Con: "const",
        Typ: "type",
        Var: "var",
        Fun: "func",
        Lbl: "label",
    }


### Object (ast/scope.go:78)

    // ----------------------------------------------------------------------------
    // Objects
    
    // An Object describes a named language entity such as a package,
    // constant, type, variable, function (incl. methods), or label.
    //
    // The Data fields contains object-specific data:
    //
    //	Kind    Data type         Data value
    //	Pkg	*types.Package    package scope
    //	Con     int               iota for the respective declaration
    //	Con     != nil            constant value
    //	Typ     *Scope            (used as method scope during type checking - transient)
    //
    type Object struct {
        Kind ObjKind
        Name string      // declared name
        Decl interface{} // corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
        Data interface{} // object-specific data; or nil
        Type interface{} // placeholder for type information; may be nil
    }

### NewObj() (ast/scope.go:87)

    // NewObj creates a new object of a given kind and name.
    func NewObj(kind ObjKind, name string) *Object {
        return &Object{Kind: kind, Name: name}
    }

### Object.Decl Types (ast/scope.go:94)

	switch d := obj.Decl.(type) {
	case *Field:
	case *ImportSpec:
	case *ValueSpec:
	case *TypeSpec:
	case *FuncDecl:
	case *LabeledStmt:
	case *AssignStmt:
	case *Scope:

## Scope

### Scope (ast/scope.go:19)

    // A Scope maintains the set of named language entities declared
    // in the scope and a link to the immediately surrounding (outer)
    // scope.
    //
    type Scope struct {
        Outer   *Scope
        Objects map[string]*Object
    }

### NewScope() (ast/scope.go:25)

    // NewScope creates a new scope nested in the outer scope.
    func NewScope(outer *Scope) *Scope {
        const n = 4 // initial scope capacity
        return &Scope{outer, make(map[string]*Object, n)}
    }

### Lookup() (ast/scope.go:34)

    // Lookup returns the object with the given name if it is
    // found in scope s, otherwise it returns nil. Outer scopes
    // are ignored.
    //
    func (s *Scope) Lookup(name string) *Object {
        return s.Objects[name]
    }

### Insert() (ast/scope.go:43)

    // Insert attempts to insert a named object obj into the scope s.
    // If the scope already contains an object alt with the same name,
    // Insert leaves the scope unchanged and returns alt. Otherwise
    // it inserts obj and returns nil.
    //
    func (s *Scope) Insert(obj *Object) (alt *Object) {
        if alt = s.Objects[obj.Name]; alt == nil {
            s.Objects[obj.Name] = obj
        }
        return
    }

## Node

### Node (ast/ast.go:35)

    // All node types implement the Node interface.
    type Node interface {
        Pos() token.Pos // position of first character belonging to the node
        End() token.Pos // position of first character immediately after the node
    }

### Expr (ast/ast.go:41)

    // All expression nodes implement the Expr interface.
    type Expr interface {
        Node
        exprNode()
    }

### Stmt (ast/ast.go:47)

    // All statement nodes implement the Stmt interface.
    type Stmt interface {
        Node
        stmtNode()
    }

### Decl (ast/ast.go:53)

    // All declaration nodes implement the Decl interface.
    type Decl interface {
        Node
        declNode()
    }


## Coment

### Comment (ast/ast.go:62)

    // A Comment node represents a single //-style or /*-style comment.
    type Comment struct {
        Slash token.Pos // position of "/" starting the comment
        Text  string    // comment text (excluding '\n' for //-style comments)
    }


### ComentGroup (ast/ast.go:73)

    // A CommentGroup represents a sequence of comments
    // with no other tokens and no empty lines between.
    //
    type CommentGroup struct {
        List []*Comment // len(List) > 0
    }

### Text() (ast/ast.go:96)

    // Text returns the text of the comment.
    // Comment markers (//, /*, and */), the first space of a line comment, and
    // leading and trailing empty lines are removed. Multiple empty lines are
    // reduced to one, and trailing space on lines is trimmed. Unless the result
    // is empty, it is newline-terminated.
    //
    func (g *CommentGroup) Text() string {

## Field

### Field (ast/ast.go:157)

    // A Field represents a Field declaration list in a struct type,
    // a method list in an interface type, or a parameter/result declaration
    // in a signature.
    //
    type Field struct {
        Doc     *CommentGroup // associated documentation; or nil
        Names   []*Ident      // field/method/parameter names; or nil if anonymous field
        Type    Expr          // field/method/parameter type
        Tag     *BasicLit     // field tag; or nil
        Comment *CommentGroup // line comments; or nil
    }

### FieldList (ast/ast.go:180)

    // A FieldList represents a list of Fields, enclosed by parentheses or braces.
    type FieldList struct {
        Opening token.Pos // position of opening parenthesis/brace, if any
        List    []*Field  // field list; or nil
        Closing token.Pos // position of closing parenthesis/brace, if any
    }


### NumFields() (ast/ast.go:211)

    // NumFields returns the number of (named and anonymous fields) in a FieldList.
    func (f *FieldList) NumFields() int {


## Ident

### Ident (ast/ast.go:238)

	// An Ident node represents an identifier.
	Ident struct {
		NamePos token.Pos // identifier position
		Name    string    // identifier name
		Obj     *Object   // denoted object; or nil
	}

### NewIdent() (ast/ast.go:521)

    // NewIdent creates a new Ident without position.
    // Useful for ASTs generated by code other than the Go parser.
    //
    func NewIdent(name string) *Ident { return &Ident{token.NoPos, name, nil} }

## Expr

### exprNode list (ast/ast.go:491)

    // exprNode() ensures that only expression/type nodes can be
    // assigned to an Expr.
    //
    func (* BadExpr        ) exprNode() {}

    func (* Ident          ) exprNode() {}

    func (* Ellipsis       ) exprNode() {}

    func (* BasicLit       ) exprNode() {}
    func (* FuncLit        ) exprNode() {}
    func (* CompositeLit   ) exprNode() {}

    func (* ParenExpr      ) exprNode() {}
    func (* SelectorExpr   ) exprNode() {}
    func (* IndexExpr      ) exprNode() {}
    func (* SliceExpr      ) exprNode() {}
    func (* TypeAssertExpr ) exprNode() {}
    func (* CallExpr       ) exprNode() {}
    func (* StarExpr       ) exprNode() {}
    func (* UnaryExpr      ) exprNode() {}
    func (* BinaryExpr     ) exprNode() {}
    func (* KeyValueExpr   ) exprNode() {}
    
    func (* ArrayType      ) exprNode() {}
    func (* StructType     ) exprNode() {}
    func (* FuncType       ) exprNode() {}
    func (* InterfaceType  ) exprNode() {}
    func (* MapType        ) exprNode() {}
    func (* ChanType       ) exprNode() {}

## Stmt

### stmtNode List ast/ast.go:792

    // stmtNode() ensures that only statement nodes can be
    // assigned to a Stmt.
    //
    func (* BadStmt        ) stmtNode() {}

    func (* DeclStmt       ) stmtNode() {}
    func (* EmptyStmt      ) stmtNode() {}
    func (* LabeledStmt    ) stmtNode() {}
    func (* ExprStmt       ) stmtNode() {}
    func (* SendStmt       ) stmtNode() {}
    func (* IncDecStmt     ) stmtNode() {}
    func (* AssignStmt     ) stmtNode() {}
    func (* GoStmt         ) stmtNode() {}
    func (* DeferStmt      ) stmtNode() {}
    func (* ReturnStmt     ) stmtNode() {}
    func (* BranchStmt     ) stmtNode() {}
    func (* BlockStmt      ) stmtNode() {}
    func (* IfStmt         ) stmtNode() {}
    func (* CaseClause     ) stmtNode() {}
    func (* SwitchStmt     ) stmtNode() {}
    func (* TypeSwitchStmt ) stmtNode() {}
    func (* CommClause     ) stmtNode() {}
    func (* SelectStmt     ) stmtNode() {}
    func (* ForStmt        ) stmtNode() {}
    func (* RangeStmt      ) stmtNode() {}

## Spec

### SpeckNode List (ast/ast.go:888)

    // specNode() ensures that only spec nodes can be
    // assigned to a Spec.
    //
    func (* ImportSpec ) specNode() {}
    func (* ValueSpec  ) specNode() {}
    func (* TypeSpec   ) specNode() {}


## Decl

### DeclNode List (ast/ast.go:956)

    // declNode() ensures that only declaration nodes can be
    // assigned to a Decl.
    //
    func (* BadDecl  ) declNode() {}

    func (* GenDecl  ) declNode() {}
    func (* FuncDecl ) declNode() {}

## File

### File (aat/ast.go:969)

    // A File node represents a Go source file.
    //
    // The Comments list contains all comments in the source file in order of
    // appearance, including the comments that are pointed to from other nodes
    // via Doc and Comment fields.
    //
    type File struct {
        Doc        *CommentGroup   // associated documentation; or nil
        Package    token.Pos       // position of "package" keyword
        Name       *Ident          // package name
        Decls      []Decl          // top-level declarations; or nil
        Scope      *Scope          // package scope (this file only)
        Imports    []*ImportSpec   // imports in this file
        Unresolved []*Ident        // unresolved identifiers in this file
        Comments   []*CommentGroup // list of all comments in the source file
    }

## Package

### Package (ast/ast.go:991)

    // A Package node represents a set of source files
    // collectively building a Go package.
    //
    type Package struct {
        Name    string             // package name
        Scope   *Scope             // package scope across all files
        Imports map[string]*Object // map of package id -> package object
        Files   map[string]*File   // Go source files by filename
    }
