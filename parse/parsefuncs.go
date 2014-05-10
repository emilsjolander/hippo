package parse

import (
	"errors"

	"github.com/emilsjolander/hippo/ast"
	"github.com/emilsjolander/hippo/lex"
)

func parseRoot(p *parser) ast.Node {
	root := &ast.Root{}
	for l := p.next(); l.Tok != lex.EOF; l = p.next() {
		if l.Tok != lex.OpenParen {
			return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.OpenParen)
		}
		root.Nodes = append(root.Nodes, parseInsideParen(p))
	}
	return root
}

func parseInsideParen(p *parser) ast.Node {
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

func parseFuncDeclaration(p *parser) ast.Node {
	decl := &ast.FuncDeclaration{}
	decl.StartPos = p.current().Start

	// parse until start of function body
	properties, err := parseProperties(p, lex.OpenParen)
	if err != nil {
		return p.errorf(err.Error())
	}

	// make sure function name and type exists
	if len(properties) < 1 {
		return p.errorf("Unexpected token %v, expected %v", p.current().Tok, lex.Identifier)
	}

	// make sure function has atleast one argument
	if len(properties) < 2 {
		return p.errorf("Function must have atleast one argument")
	}

	decl.Name = properties[0].Name
	decl.Typ = properties[0].Typ
	decl.Args = properties[1:]

	// make sure function body exists
	if l := p.next(); l.Tok == lex.Identifier {
		decl.Body = parseExpression(p)
	} else {
		return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
	}

	// make sure function is closed directly after body
	if l := p.next(); l.Tok != lex.CloseParen {
		return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.CloseParen)
	}

	decl.EndPos = p.current().Start
	return decl
}

func parseTypeDeclaration(p *parser) ast.Node {
	decl := &ast.TypeDeclaration{}
	decl.StartPos = p.current().Start

	// make sure type name exists
	if l := p.next(); l.Tok == lex.Identifier {
		decl.Name = l.Val
	} else {
		return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
	}

	// parse until end of type declaration
	properties, err := parseProperties(p, lex.CloseParen)
	if err != nil {
		return p.errorf(err.Error())
	}
	decl.Properties = properties

	// make sure type has atleast one property
	if len(decl.Properties) == 0 {
		return p.errorf("Too few properties for type %v", decl.Name)
	}

	decl.EndPos = p.current().Start
	return decl
}

func parseExpression(p *parser) ast.Node {
	// identifier of expression has allready been parsed, get it
	expr := &ast.Expression{
		Name: p.current().Val,
	}
	expr.StartPos = p.current().Start

	// parse args of expression
	for {
		l := p.next()
		switch l.Tok {
		case lex.Identifier:
			// a variable as an argument
			expr.Args = append(expr.Args, parseIdentifier(p))
		case lex.OpenParen:
			// a sub expression as a argument
			if l = p.next(); l.Tok == lex.Identifier {
				expr.Args = append(expr.Args, parseExpression(p))
			} else {
				return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
			}
		case lex.CloseParen:
			// no more args, break switch and for-ever loop
			goto finished
		default:
			// a literal argument
			if l.Tok.IsLiteral() {
				expr.Args = append(expr.Args, &ast.Literal{
					StartPos: l.Start,
					Typ:      l.Tok.String(),
					Val:      l.Val,
				})
			} else {
				return p.errorf("Unexpected token %v, expected on of %v, %v, %v, or a type literal",
					l.Tok, lex.Identifier, lex.OpenParen, lex.CloseParen)
			}
		}
	}

finished:
	expr.EndPos = p.current().Start
	return expr
}

func parseIdentifier(p *parser) ast.Node {
	// identifier name has allready been parsed, get it
	iden := &ast.Identifier{
		Parts: []string{p.current().Val},
	}
	iden.StartPos = p.current().Start

	// loop until all sub identifiers have been added
	for p.peak().Tok == lex.Dot {
		p.next() // skip dot
		l := p.next()
		if l.Tok != lex.Identifier {
			return p.errorf("Unexpected token %v, expected %v", l.Tok, lex.Identifier)
		}
		iden.Parts = append(iden.Parts, l.Val)
	}

	iden.EndPos = p.current().Start
	return iden
}

func parseProperties(p *parser, stop lex.Token) ([]ast.Property, error) {
	// parse properties defined as name:type
	var props []ast.Property
	for l := p.next(); l.Tok != stop; l = p.next() {
		name := l
		colon := p.next()
		typ := p.next()
		if name.Tok != lex.Identifier || colon.Tok != lex.Colon || typ.Tok != lex.Identifier {
			return props, errors.New("Property expected with syntax name:type")
		}
		props = append(props, ast.Property{
			Name: name.Val,
			Typ:  typ.Val,
		})
	}
	return props, nil
}
