package parse

import (
	"fmt"
	"strings"
)

func PrintAst(root Node) {
	printAst(root, 0)
}

func printAst(root Node, lvl int) {
	prefix := strings.Repeat("  ", lvl)
	switch t := root.(type) {
	case errorNode:
		fmt.Println(prefix + "errorNode: " + t.err)
	case Root:
		fmt.Println(prefix + "Root")
		for _, n := range t.Nodes {
			printAst(n, lvl+1)
		}
	case Identifier:
		fmt.Println(prefix + "Identifier: " + strings.Join(t.Parts, "."))
	case Expression:
		fmt.Println(prefix + "Expression: " + t.Name)
		for _, n := range t.Args {
			printAst(n, lvl+1)
		}
	case TypeDeclaration:
		fmt.Println(prefix + "TypeDeclaration: " + t.Name)
		for _, p := range t.Properties {
			prefix = strings.Repeat("  ", lvl+1)
			fmt.Println(prefix + p.Name + ":" + p.Typ)
		}
	case FuncDeclaration:
		props := t.Name + ":" + t.Typ
		for _, p := range t.Args {
			props += " " + p.Name + ":" + p.Typ
		}
		fmt.Println(prefix + "FuncDeclaration: " + props)
		printAst(t.Body, lvl+1)
	case Literal:
		fmt.Println(prefix + "Literal: " + t.Val)
	}
}
