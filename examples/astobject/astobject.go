package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example.go", src, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Name == "x" {
			var decl interface{}
			if ident != nil && ident.Obj != nil {
				decl = ident.Obj.Decl
			}
			fmt.Printf("%-17sobj=%-12p  decl=%T\n", fset.Position(ident.Pos()), ident.Obj, decl)
		}
		return true
	})
}

var src = `package p

import x "pkg"

func f() {
    if x := x.f(); x != nil {
        x(func(x int) int { return x + 1 })
    }
}
`
