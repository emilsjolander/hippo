package gen

import (
	"fmt"
	"os"

	"github.com/emilsjolander/hippo/ast"
)

type Generator interface {
	Write(root ast.Node, dir string)
}

type basicGenerator struct {
	file *os.File
}

func (g *basicGenerator) putf(output string, args ...interface{}) {
	_, err := g.file.WriteString(fmt.Sprintf(output, args...))
	if err != nil {
		panic(err)
	}
}

func propertyTypes(props []ast.Property) []string {
	types := make([]string, len(props))
	for i, p := range props {
		types[i] = p.Typ
	}
	return types
}

func overloadedName(name string, types []string) string {
	for _, t := range types {
		name += "_" + t
	}
	return name
}

func translateName(name string) string {
	translated := ""
	for _, r := range name {
		switch r {
		case '+':
			translated += "PLUS"
		case '-':
			translated += "MINUS"
		case '*':
			translated += "ASTERISK"
		case '/':
			translated += "SLASH"
		case '^':
			translated += "CARET_UP"
		case '<':
			translated += "CARET_LEFT"
		case '>':
			translated += "CARET_RIGHT"
		case '=':
			translated += "EQUALS"
		case '%':
			translated += "PERCENT"
		default:
			translated += string(r)
		}
	}
	return translated
}
