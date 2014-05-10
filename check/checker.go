package check

import (
	"fmt"

	"github.com/emilsjolander/hippo/ast"
)

type checker struct {
	errors []error
	scope  scope
}

func (c *checker) errorf(cause string, args ...interface{}) {
	cause = fmt.Sprintf(cause, args...)
	node := c.scope.stack[len(c.scope.stack)-1]
	c.errors = append(
		c.errors,
		fmt.Errorf("%v - "+cause, node.Start()),
	)
}

func Check(root ast.Node) []error {
	c := &checker{
		scope: scope{
			types: builtinTypes,
			funcs: builtinFuncs,
		},
	}
	switch n := root.(type) {
	case *ast.Root:
		checkRoot(c, n)
	default:
		c.errorf("Root ast node must be of type *ast.Root")
	}
	return c.errors
}
