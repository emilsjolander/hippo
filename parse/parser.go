package parse

import (
	"fmt"
	"github.com/emilsjolander/hippo/lex"
)

type parser struct {
	lexemes   chan lex.Lexeme
	current   lex.Lexeme
	errors    []Error
	eof       bool
	hasPeaked bool
	peaked    lex.Lexeme
}

func (p *parser) next() lex.Lexeme {
	if p.hasPeaked {
		p.hasPeaked = false
		return p.peaked
	}
	if !p.eof {
		p.current = <-p.lexemes
		if p.current.Tok == lex.EOF {
			p.eof = true
		}
	}
	return p.current
}

func (p *parser) peak() lex.Lexeme {
	l := p.next()
	p.hasPeaked = true
	p.peaked = l
	return l
}

func (p *parser) errorf(cause string, args ...interface{}) errorNode {
	err := Error{
		Cause: fmt.Sprintf(cause, args...),
		Start: p.current.Start,
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
	err.End = p.current.Start
	p.errors = append(p.errors, err)
	return errorNode{
		err: err.Cause,
	}
}

func Parse(lexemes chan lex.Lexeme) Node {
	p := &parser{
		lexemes: lexemes,
	}
	root := parseRoot(p)
	return root
}
