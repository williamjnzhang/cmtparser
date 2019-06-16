package main

import (
	"fmt"
	"cmt/ast"
	"cmt/parser"
	"go/token"
	"strings"
)

var (
	functypesrc = "./commentFuncType.src"
	strutypesrc = "./commentStructType.src"
)

var a struct {
	b string
}

func main() {

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, strutypesrc, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	var v visitor
	ast.Walk(v, f)
}

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	switch t := n.(type) {
	// case *ast.CommentGroup:
	// 	if t.List != nil {
	// 		for _, c := range t.List {
	// 			fmt.Printf("cmt: %s", c.Text)
	// 		}
	// 	}
	case *ast.Ident:
		if t.Comment != nil {
			for _, c := range t.Comment.List {
				fmt.Printf("ident cmt: %s", c.Text)
			}
		}
	}
	return v + 1
}
