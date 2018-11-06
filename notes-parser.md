# Parser Notes

##

    package
    imports
    decls
        const
            valueSpec
        var
            valueSpec
        tye
            typeSpec
        func

#### File Parser (parser/parser.go:2452)

    func (p *parser) parseFile() *ast.File {
        if p.trace {
            defer un(trace(p, "File"))
        }
    
        // Don't bother parsing the rest if we had errors scanning the first token.
        // Likely not a Go source file at all.
        if p.errors.Len() != 0 {
            return nil
        }
    
        // package clause
        doc := p.leadComment
        pos := p.expect(token.PACKAGE)
        // Go spec: The package clause is not a declaration;
        // the package name does not appear in any scope.
        ident := p.parseIdent()
        if ident.Name == "_" && p.mode&DeclarationErrors != 0 {
            p.error(p.pos, "invalid package name _")
        }
        p.expectSemi()
    
        // Don't bother parsing the rest if we had errors parsing the package clause.
        // Likely not a Go source file at all.
        if p.errors.Len() != 0 {
            return nil
        }
    
        p.openScope()
        p.pkgScope = p.topScope
        var decls []ast.Decl
        if p.mode&PackageClauseOnly == 0 {
            // import decls
            for p.tok == token.IMPORT {
                decls = append(decls, p.parseGenDecl(token.IMPORT, p.parseImportSpec))
            }
    
            if p.mode&ImportsOnly == 0 {
                // rest of package body
                for p.tok != token.EOF {
                    decls = append(decls, p.parseDecl(syncDecl))
                }
            }
        }
        p.closeScope()
        assert(p.topScope == nil, "unbalanced scopes")
        assert(p.labelScope == nil, "unbalanced label scopes")
    
        // resolve global identifiers within the same file
        i := 0
        for _, ident := range p.unresolved {
            // i <= index for current ident
            assert(ident.Obj == unresolved, "object already resolved")
            ident.Obj = p.pkgScope.Lookup(ident.Name) // also removes unresolved sentinel
            if ident.Obj == nil {
                p.unresolved[i] = ident
                i++
            }
        }
    
        return &ast.File{
            Doc:        doc,
            Package:    pos,
            Name:       ident,
            Decls:      decls,
            Scope:      p.pkgScope,
            Imports:    p.imports,
            Unresolved: p.unresolved[0:i],
            Comments:   p.comments,
        }
    }


### parser struct (parser/parser.go:30)

    // The parser structure holds the parser's internal state.
    type parser struct {
        file    *token.File
        errors  scanner.ErrorList
        scanner scanner.Scanner
    
        // Tracing/debugging
        mode   Mode // parsing mode
        trace  bool // == (mode & Trace != 0)
        indent int  // indentation used for tracing output
    
        // Comments
        comments    []*ast.CommentGroup
        leadComment *ast.CommentGroup // last lead comment
        lineComment *ast.CommentGroup // last line comment
    
        // Next token
        pos token.Pos   // token position
        tok token.Token // one token look-ahead
        lit string      // token literal
    
        // Error recovery
        // (used to limit the number of calls to syncXXX functions
        // w/o making scanning progress - avoids potential endless
        // loops across multiple parser functions during error recovery)
        syncPos token.Pos // last synchronization position
        syncCnt int       // number of calls to syncXXX without progress
    
        // Non-syntactic parser control
        exprLev int  // < 0: in control clause, >= 0: in expression
        inRhs   bool // if set, the parser is parsing a rhs expression
    
        // Ordinary identifier scopes
        pkgScope   *ast.Scope        // pkgScope.Outer == nil
        topScope   *ast.Scope        // top-most scope; may be pkgScope
        unresolved []*ast.Ident      // unresolved identifiers
        imports    []*ast.ImportSpec // list of imports
    
        // Label scopes
        // (maintained by open/close LabelScope)
        labelScope  *ast.Scope     // label scope for current function
        targetStack [][]*ast.Ident // stack of unresolved labels
    }


# parsStmt

		// tokens that may start an expression
		token.IDENT, 
		token.INT, 
		token.FLOAT, 
		token.IMAG, 
		token.CHAR, 
		token.STRING, 
		token.FUNC, 
		token.LPAREN, // operands
		
		token.LBRACK, 
		token.STRUCT, 
		token.MAP, 
		token.CHAN, 
		token.INTERFACE, // composite types
		
		token.ADD, 
		token.SUB, 
		token.MUL, 
		token.AND, 
		token.XOR, 
		token.ARROW, 
		token.NOT: // unary operators
		
		s, _ = p.parseSimpleStmt(labelOk)
		// because of the required look-ahead, labeled statements are

## parseSimpleStmt

    -> parseLHSList
    
## praseLhsList

    -> parseExprList(lhs:true)

## parseExprList

    -> parseExpr(lsh)


## parseExpr(lhs)

    -> parseBinaryExpr(lhs, token.LowestPrec+1)

## parseBinaryExpr(lhs, prec)

    -> parseUnaryExpr(lhs)
    
## parseUnaryExpr(lhs)

