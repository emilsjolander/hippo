package ast

import (
	"github.com/emilsjolander/hippo/lex"
)

type Node interface {
	Start() lex.Pos
	End() lex.Pos
	Type() string
}

type basicNode struct {
	StartPos lex.Pos
	EndPos   lex.Pos
	Typ      string
}

func (n *basicNode) Start() lex.Pos {
	return n.StartPos
}

func (n *basicNode) End() lex.Pos {
	return n.EndPos
}

func (n *basicNode) Type() string {
	return n.Typ
}

type Error struct {
	basicNode
	Err string
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

type Property struct {
	Name string
	Typ  string
}

type TypeDeclaration struct {
	basicNode
	Name       string
	Properties []Property
}

type FuncDeclaration struct {
	basicNode
	Name string
	Args []Property
	Body Node
}

type Root struct {
	Nodes []Node
}

func (r *Root) Start() lex.Pos {
	if len(r.Nodes) == 0 {
		return lex.Pos{Row: 0, Col: 0}
	}
	return r.Nodes[0].Start()
}

func (r *Root) End() lex.Pos {
	if len(r.Nodes) == 0 {
		return lex.Pos{Row: 0, Col: 0}
	}
	return r.Nodes[len(r.Nodes)-1].End()
}

func (r *Root) Type() string {
	return ""
}

type Literal struct {
	StartPos lex.Pos
	Typ      string
	Val      string
}

func (l *Literal) Start() lex.Pos {
	return l.StartPos
}

func (l *Literal) End() lex.Pos {
	return lex.Pos{
		Row: l.StartPos.Row,
		Col: l.StartPos.Col + len(l.Val),
	}
}

func (l *Literal) Type() string {
	return l.Typ
}
