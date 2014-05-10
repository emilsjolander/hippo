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
		return "error"
	case EOF:
		return "eof"
	case OpenParen:
		return "openparen"
	case CloseParen:
		return "closeparen"
	case Colon:
		return "colon"
	case Type:
		return "type"
	case Function:
		return "func"
	case Identifier:
		return "identifier"
	case Dot:
		return "dot"
	case Float:
		return "float"
	case Integer:
		return "int"
	case String:
		return "string"
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
