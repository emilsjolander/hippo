package lex

import (
	"fmt"
)

type Token int

const (
	Error Token = iota
	EOF
	OpenParen
	CloseParen
	Colon
	Type
	Function
	Identifier
	Dot

	tokenLiteralStart
	Float
	Integer
	String
	tokenLiteralEnd
)

func (t Token) String() string {
	switch t {
	case Error:
		return "Error"
	case EOF:
		return "EOF"
	case OpenParen:
		return "OpenParen"
	case CloseParen:
		return "CloseParen"
	case Colon:
		return "Colon"
	case Type:
		return "Type"
	case Function:
		return "Function"
	case Identifier:
		return "Identifier"
	case Dot:
		return "Dot"
	case Float:
		return "Float"
	case Integer:
		return "Integer"
	case String:
		return "String"
	}
	panic("Forgot to add string representation for some token")
}

func (t Token) IsLiteral() bool {
	return t > tokenLiteralStart && t < tokenLiteralEnd
}

type Lexeme struct {
	Tok   Token
	Val   string
	Start Pos
}

func (l Lexeme) String() string {
	switch l.Tok {
	case EOF:
		return fmt.Sprintf("%v - EOF", l.Start)
	default:
		return fmt.Sprintf("%v - %v: %s", l.Start, l.Tok, l.Val)
	}
}
