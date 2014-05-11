package check

import (
	"github.com/emilsjolander/hippo/ast"
)

const (
	undefined = "_undefined"
)

type typ struct {
	name       string
	properties []ast.Property
}

type function struct {
	name string
	typ  string
	args []ast.Property
}

type scope struct {
	types []*typ
	funcs []*function
	vars  []*ast.Property
	stack []ast.Node
}

func (s *scope) push(n ast.Node) {
	s.stack = append(s.stack, n)
}

func (s *scope) pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *scope) registerType(c *checker, name string) {
	for _, t := range s.types {
		if name == t.name {
			c.errorf("Type %s allready defined once", name)
			return
		}
	}
	s.types = append(s.types, &typ{name: name})
}

func (s *scope) registerTypeProperty(c *checker, typ string, property ast.Property) {
	propertyTypeExists := false
	for _, t := range s.types {
		if property.Typ == t.name {
			propertyTypeExists = true
		}
	}
	if !propertyTypeExists {
		c.errorf("Undefined type %s", property.Typ)
		return
	}

	for _, t := range s.types {
		if typ == t.name {
			for _, p := range t.properties {
				if property.Name == p.Name {
					c.errorf("Property %s for type %s allready defined once", property.Name, typ)
					return
				}
			}
			t.properties = append(t.properties, property)
			return
		}
	}
	c.errorf("Undefined type %s", typ)
}

func (s *scope) registerFunc(c *checker, fun string, typ string, args []ast.Property) {

	// check types
	for _, t := range s.types {
		if fun == t.name {
			if len(t.properties) != len(args) {
				goto continueCheckTypes
			}
			for i, a := range t.properties {
				if args[i].Typ != a.Typ {
					goto continueCheckTypes
				}
			}
			types := make([]string, len(args))
			for i, a := range args {
				types[i] = a.Name
			}
			c.errorf("Type %s with property types %v allready defined once", fun, types)
		}
	continueCheckTypes:
	}

	for _, f := range s.funcs {
		if fun == f.name {
			if len(f.args) != len(args) {
				goto continueCheckFuncs
			}
			for i, a := range f.args {
				if args[i].Typ != a.Typ {
					goto continueCheckFuncs
				}
			}
			types := make([]string, len(args))
			for i, a := range args {
				types[i] = a.Name
			}
			c.errorf("Function %s with input types %v allready defined once", fun, types)
			return
		}
	continueCheckFuncs:
	}

	s.funcs = append(s.funcs, &function{
		name: fun,
		typ:  typ,
		args: args,
	})
}

func (s *scope) registerVar(c *checker, variable ast.Property) {
	for _, v := range s.vars {
		if variable.Name == v.Name {
			c.errorf("Variable with name %s allready defined once", variable.Name)
		}
	}
	s.vars = append(s.vars, &variable)
}

func (s *scope) unregisterVar(c *checker, variable ast.Property) {
	for i, v := range s.vars {
		if variable.Name == v.Name {
			s.vars = append(s.vars[:i], s.vars[i+1:]...)
			return
		}
	}
	c.errorf("Undefined varable %s", variable.Name)
}

func (s *scope) getExpressionType(c *checker, name string, types []string) string {

	// check types
	for _, t := range s.types {
		if name == t.name {
			if len(t.properties) != len(types) {
				goto continueCheckTypes
			}
			for i, a := range t.properties {
				if types[i] != a.Typ {
					goto continueCheckTypes
				}
			}
			return t.name
		}
	continueCheckTypes:
	}

	// check functions
	for _, f := range s.funcs {
		if name == f.name {
			if len(f.args) != len(types) {
				goto continueCheckFuncs
			}
			for i, a := range f.args {
				if types[i] != a.Typ {
					goto continueCheckFuncs
				}
			}
			return f.typ
		}
	continueCheckFuncs:
	}

	c.errorf("Undefined function or type %s for input arguments of type %v", name, types)
	return undefined
}

func (s *scope) getVariableType(c *checker, name string) string {
	for _, v := range s.vars {
		if name == v.Name {
			return v.Typ
		}
	}
	c.errorf("Undefined variable %s", name)
	return undefined
}

func (s *scope) getPropertyType(c *checker, typ string, prop string) string {
	for _, t := range s.types {
		if typ == t.name {
			for _, p := range t.properties {
				if prop == p.Name {
					return p.Typ
				}
			}
			c.errorf("Undefined property %v for type %s", prop, typ)
			return undefined
		}
	}
	c.errorf("Undefined type %v", typ)
	return undefined
}
