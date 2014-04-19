package parse

import (
	"strconv"

	"github.com/emilsjolander/hippo/lex"
)

func Parse(input chan lex.Token) Node {
	nodes := rootNodes(input)
	return treeFromNodes(nodes)
}

func rootNodes(input chan lex.Token) []Node {
	nodes := make([]Node, 0)
	for t := range input {
		switch t.Typ {
		case lex.TokenNumber:
			f, _ := strconv.ParseFloat(t.Val, 64)
			n := &number{f}
			nodes = append(nodes, n)
		case lex.TokenOperator:
			n := &operator{sign: t.Val}
			nodes = append(nodes, n)
		case lex.TokenOpenParen:
			n := &number{Parse(input).Value()}
			nodes = append(nodes, n)
		case lex.TokenCloseParen, lex.TokenEOF:
			return nodes
		case lex.TokenError:
			panic(t.Val)
		}
	}
	return nodes
}

func treeFromNodes(nodes []Node) Node {
	if len(nodes) == 1 {
		return nodes[0]
	}
	if len(nodes)%2 == 0 {
		panic("error: number of nodes must be odd")
	}

	i := indexOfRootOperator(nodes)
	op := nodes[i].(*operator)
	op.left = treeFromNodes(nodes[:i])
	op.right = treeFromNodes(nodes[i+1:])
	return op
}

func indexOfRootOperator(nodes []Node) int {
	var index int
	precedence := 999
	for i, n := range nodes {
		switch t := n.(type) {
		case *operator:
			if t.precedence() <= precedence {
				precedence = t.precedence()
				index = i
			}
		}
	}
	return index
}
