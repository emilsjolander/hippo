package parse

import (
	"fmt"

	"github.com/emilsjolander/hippo/ast"
	"github.com/emilsjolander/hippo/lex"
)

type parser struct {
	lexemes []lex.Lexeme
	errors  []error
	pos     int
}

func Parse(lexemes []lex.Lexeme) (ast.Node, []error) {
	p := &parser{
		lexemes: lexemes,
		pos:     -1,
	}
	root := parseRoot(p)
	return root, p.errors
}

func (p *parser) next() lex.Lexeme {
	p.pos++
	if p.pos >= len(p.lexemes) {
		p.pos = len(p.lexemes) - 1
	}
	return p.lexemes[p.pos]
}

func (p *parser) current() lex.Lexeme {
	return p.lexemes[p.pos]
}

func (p *parser) backup() {
	p.pos--
}

func (p *parser) peak() lex.Lexeme {
	l := p.next()
	p.backup()
	return l
}

func (p *parser) errorf(cause string, args ...interface{}) *ast.Error {
	err := Error{
		Cause: fmt.Sprintf(cause, args...),
		Start: p.current().Start,
	}
	scope := 1
	for scope > 0 {
		l := p.next()
		switch l.Tok {
		case lex.OpenParen:
			scope++
		case lex.CloseParen:
			scope--
		case lex.EOF:
			scope = 0
		}
	}
	err.End = p.current().Start
	p.errors = append(p.errors, err)
	return &ast.Error{
		Err: err.Cause,
	}
}
