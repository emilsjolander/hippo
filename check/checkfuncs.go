package check

import (
	"github.com/emilsjolander/hippo/ast"
)

func checkRoot(c *checker, r *ast.Root) {
	c.scope.push(r)
	defer c.scope.pop()

	// gather type names
	for _, n := range r.Nodes {
		if t, ok := n.(*ast.TypeDeclaration); ok {
			c.scope.registerType(c, t.Name)
		}
	}

	// gather type properties now that all types are known
	for _, n := range r.Nodes {
		if t, ok := n.(*ast.TypeDeclaration); ok {
			for _, p := range t.Properties {
				c.scope.registerTypeProperty(c, t.Name, p)
			}
		}
	}

	// gather all function definitions now that all types are known
	for _, n := range r.Nodes {
		if f, ok := n.(*ast.FuncDeclaration); ok {
			c.scope.registerFunc(c, f.Name, f.Typ, f.Args)
		}
	}

	// check all function bodies and script body now that all functions and types are known
	for _, n := range r.Nodes {
		if f, ok := n.(*ast.FuncDeclaration); ok {
			for _, a := range f.Args {
				c.scope.registerVar(c, a)
			}
			checkExpression(c, f.Body.(*ast.Expression))
			for _, a := range f.Args {
				c.scope.unregisterVar(c, a)
			}
		}
		if e, ok := n.(*ast.Expression); ok {
			checkExpression(c, e)
		}
	}
}

func checkExpression(c *checker, e *ast.Expression) string {
	c.scope.push(e)
	defer c.scope.pop()

	var types []string
	for _, a := range e.Args {
		switch n := a.(type) {
		case *ast.Identifier:
			types = append(types, checkIdentifier(c, n))
		case *ast.Expression:
			types = append(types, checkExpression(c, n))
		case *ast.Literal:
			types = append(types, n.Typ)
		}
	}
	e.ArgTypes = types

	return c.scope.getExpressionType(c, e.Name, types)
}

func checkIdentifier(c *checker, i *ast.Identifier) string {
	c.scope.push(i)
	defer c.scope.pop()

	typ := c.scope.getVariableType(c, i.Parts[0])
	for _, p := range i.Parts[1:] {
		typ = c.scope.getPropertyType(c, typ, p)
	}
	return typ
}
