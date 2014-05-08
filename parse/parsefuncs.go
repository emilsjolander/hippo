package parse

import (
	"errors"
	"github.com/emilsjolander/hippo/lex"
)

func parseRoot(p *parser) Node {
	root := Root{}
	for l := p.next(); l.Tok != lex.EOF; l = p.next() {
		if l.Tok != lex.OpenParen {
			return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.OpenParen)
		}
		root.Nodes = append(root.Nodes, parseInsideParen(p))
	}
	return root
}

func parseInsideParen(p *parser) Node {
	l := p.next()
	switch l.Tok {
	case lex.Function:
		return parseFuncDeclaration(p)
	case lex.Type:
		return parseTypeDeclaration(p)
	case lex.Identifier:
		return parseExpression(p)
	}
	return p.errorf("Unexpected token %v, expected on of %v, %v, or %v",
		l.Tok, lex.Function, lex.Type, lex.Identifier)
}

func parseFuncDeclaration(p *parser) Node {
	decl := FuncDeclaration{}
	decl.start = p.current.Start

	// parse until start of function body
	properties, err := parseProperties(p, lex.OpenParen)
	if err != nil {
		return p.errorf(err.Error())
	}

	if len(properties) < 1 {
		return p.errorf("Unexpected token %v, expected %v", p.current.Tok, lex.Identifier)
	}
	if len(properties) < 2 {
		return p.errorf("Function must have atleast one argument")
	}

	decl.Name = properties[0].Name
	decl.Typ = properties[0].Typ
	decl.Args = properties[1:]

	if l := p.next(); l.Tok == lex.Identifier {
		decl.Body = parseExpression(p)
	} else {
		return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
	}

	if l := p.next(); l.Tok != lex.CloseParen {
		return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.CloseParen)
	}

	decl.end = p.current.Start
	return decl
}

func parseTypeDeclaration(p *parser) Node {
	decl := TypeDeclaration{}
	decl.start = p.current.Start

	if l := p.next(); l.Tok == lex.Identifier {
		decl.Name = l.Val
	} else {
		return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
	}

	properties, err := parseProperties(p, lex.CloseParen)
	if err != nil {
		return p.errorf(err.Error())
	}
	decl.Properties = properties

	if len(decl.Properties) == 0 {
		return p.errorf("Too few properties for type %v", decl.Name)
	}

	decl.end = p.current.Start
	return decl
}

func parseExpression(p *parser) Node {
	expr := Expression{
		Name: p.current.Val,
	}
	expr.start = p.current.Start

	// parse args
	for {
		l := p.next()
		switch l.Tok {
		case lex.Identifier:
			expr.Args = append(expr.Args, parseIdentifier(p))
		case lex.OpenParen:
			if l = p.next(); l.Tok == lex.Identifier {
				expr.Args = append(expr.Args, parseExpression(p))
			} else {
				return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
			}
		case lex.CloseParen:
			goto finished
		default:
			if l.Tok.IsLiteral() {
				expr.Args = append(expr.Args, Literal{
					start: l.Start,
					Typ:   l.Tok,
					Val:   l.Val,
				})
			} else {
				return p.errorf("Unexpected token %v, expected on of %v, %v, %v, or a type literal",
					l.Tok, lex.Identifier, lex.OpenParen, lex.CloseParen)
			}
		}
	}

finished:
	expr.end = p.current.Start
	return expr
}

func parseIdentifier(p *parser) Node {
	iden := Identifier{
		Parts: []string{p.current.Val},
	}

	for p.peak().Tok == lex.Dot {
		p.next() // skip dot
		l := p.next()
		if l.Tok != lex.Identifier {
			return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
		}
		iden.Parts = append(iden.Parts, l.Val)
	}

	return iden
}

func parseProperties(p *parser, stop lex.Token) ([]Property, error) {
	var props []Property
	for l := p.next(); l.Tok != stop; l = p.next() {
		name := l
		colon := p.next()
		typ := p.next()
		if name.Tok != lex.Identifier || colon.Tok != lex.Colon || typ.Tok != lex.Identifier {
			return props, errors.New("Property expected with syntax name:type")
		}
		props = append(props, Property{
			Name: name.Val,
			Typ:  typ.Val,
		})
	}
	return props, nil
}
