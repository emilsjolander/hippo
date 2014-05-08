package parse

import (
	"github.com/emilsjolander/hippo/lex"
)

type Node interface {
	Start() lex.Pos
	End() lex.Pos
}

type basicNode struct {
	start lex.Pos
	end   lex.Pos
}

func (n basicNode) Start() lex.Pos {
	return n.start
}

func (n basicNode) End() lex.Pos {
	return n.end
}

type errorNode struct {
	basicNode
	err string
}

type Identifier struct {
	basicNode
	Parts []string
}

type Expression struct {
	basicNode
	Name string
	Args []Node
}

type TypeDeclaration struct {
	basicNode
	Name       string
	Properties []Property
}

type FuncDeclaration struct {
	basicNode
	Property
	Args []Property
	Body Node
}

type Root struct {
	Nodes []Node
}

func (r Root) Start() lex.Pos {
	if len(r.Nodes) == 0 {
		return lex.Pos{Row: 0, Col: 0}
	}
	return r.Nodes[0].Start()
}

func (r Root) End() lex.Pos {
	if len(r.Nodes) == 0 {
		return lex.Pos{Row: 0, Col: 0}
	}
	return r.Nodes[len(r.Nodes)-1].End()
}

type Literal struct {
	start lex.Pos
	Typ   lex.Token
	Val   string
}

func (l Literal) Start() lex.Pos {
	return l.start
}

func (l Literal) End() lex.Pos {
	return lex.Pos{
		Row: l.start.Row,
		Col: l.start.Col + len(l.Val),
	}
}
