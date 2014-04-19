package lex

import "errors"

// look for number or open paren
func expression(l *lexer) lexFunc {
	r := l.next()
	switch r {
	case '(':
		l.emit(TokenOpenParen)
		return expression
	default:
		l.backup()
		return number
	}
}

// look for close paren
func endExpression(l *lexer) lexFunc {
	r := l.next()
	if r == ')' {
		l.emit(TokenCloseParen)
		return endExpression
	}
	l.backup()
	return operator
}

// look for an int or a float
func number(l *lexer) lexFunc {
	n := l.acceptMany("0123456789")
	if l.acceptOne(".") {
		n = l.acceptMany("0123456789")
	}

	if n < 0 {
		l.emitError(errors.New("Unexpected token, expected number"))
		return nil
	}

	l.emit(TokenNumber)
	return endExpression
}

// look for any operator
func operator(l *lexer) lexFunc {
	r := l.next()
	switch r {
	case '+':
	case '-':
	case '*':
	case '/':
	case '^':
	default:
		l.emitError(errors.New("Unexpected token, expected operator: "))
		return nil
	}
	l.emit(TokenOperator)
	return expression
}
