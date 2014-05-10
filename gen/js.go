package gen

import (
	"os"
	"strings"

	"github.com/emilsjolander/hippo/ast"
)

const jsbuiltin = `

function PLUS_float_float(one,two) {
	return one+two
}

function MINUS_float_float(one,two) {
	return one+two
}

function ASTERISK_float_float(one,two) {
	return one*two
}

function SLASH_float_float(one,two) {
	return one+two
}

function PLUS_int_int(one,two) {
	return one+two
}

function MINUS_int_int(one,two) {
	return one+two
}

function ASTERISK_int_int(one,two) {
	return one*two
}

function SLASH_int_int(one,two) {
	return one+two
}

function PLUS_string_string(one,two) {
	return one+two
}

function print_float(o) {
	console.log(o)
}

function print_int(o) {
	console.log(o)
}

function print_string(o) {
	console.log(o)
}

`

type JS struct {
	basicGenerator
}

func (js *JS) Write(root ast.Node, dir string) {
	os.MkdirAll(dir, os.ModePerm)
	f, err := os.Create(dir + "/output.js")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	js.file = f

	if root, ok := root.(*ast.Root); ok {
		js.putf(jsbuiltin)
		js.writeRoot(root)
	} else {
		panic("Root node must be of type ast.Root")
	}
}

func (js *JS) writeRoot(r *ast.Root) {
	// write types
	for _, n := range r.Nodes {
		if t, ok := n.(*ast.TypeDeclaration); ok {
			js.writeType(t)
		}
	}

	// write funcs
	for _, n := range r.Nodes {
		if f, ok := n.(*ast.FuncDeclaration); ok {
			js.writeFunc(f)
		}
	}

	// write script
	for _, n := range r.Nodes {
		if e, ok := n.(*ast.Expression); ok {
			js.writeExpression(e)
		}
	}
}

func (js *JS) writeType(t *ast.TypeDeclaration) {
	js.putf("function %s(", translateName(overloadedName(t.Name, propertyTypes(t.Properties))))
	for i, a := range t.Properties {
		js.putf("%s", a.Name)
		if i != len(t.Properties)-1 {
			js.putf(",")
		}
	}
	js.putf("){return {")
	for i, a := range t.Properties {
		js.putf("%s:%s", a.Name, a.Name)
		if i != len(t.Properties)-1 {
			js.putf(",")
		}
	}
	js.putf("}}")
}

func (js *JS) writeFunc(f *ast.FuncDeclaration) {
	js.putf("function %s(", translateName(overloadedName(f.Name, propertyTypes(f.Args))))
	for i, a := range f.Args {
		js.putf("%s", a.Name)
		if i != len(f.Args)-1 {
			js.putf(",")
		}
	}
	js.putf("){return ")
	js.writeExpression(f.Body.(*ast.Expression))
	js.putf("}")
}

func (js *JS) writeExpression(e *ast.Expression) {
	js.putf("%s(", translateName(overloadedName(e.Name, e.ArgTypes)))
	for i, a := range e.Args {
		switch t := a.(type) {
		case *ast.Literal:
			js.writeLiteral(t)
		case *ast.Expression:
			js.writeExpression(t)
		case *ast.Identifier:
			js.writeIdentifier(t)
		}
		if i != len(e.Args)-1 {
			js.putf(",")
		}
	}
	js.putf(")\n")
}

func (js *JS) writeLiteral(l *ast.Literal) {
	switch l.Typ {
	case "string":
		js.putf("\"%s\"", l.Val)
	default:
		js.putf("%s", l.Val)
	}
}

func (js *JS) writeIdentifier(i *ast.Identifier) {
	js.putf("%s", strings.Join(i.Parts, "."))
}
