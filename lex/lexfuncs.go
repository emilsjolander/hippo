package lex

import (
	"fmt"
)

const (
	numbers = "0123456789"
	alphas  = "abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ+-/^*%"
)

func root(l *lexer) lexFunc {
	l.ignoreSpace()

	if l.acceptOne("#") {
		for l.next() != '\n' {
			// skip until new line
		}
		l.ignore()
		return root
	}

	if l.acceptOne("0123456789") {
		l.backup()
		return numberLiteral
	}

	if l.acceptOne("\"") {
		l.backup()
		return stringLiteral
	}

	if l.acceptOne("(") {
		l.emit(OpenParen)
		l.ignoreSpace()
		return identifier
	}

	if l.acceptOne(")") {
		l.emit(CloseParen)
		return root
	}

	if l.acceptOne(":") {
		l.emit(Colon)
		return identifier
	}

	return identifier
}

func identifier(l *lexer) lexFunc {
	if !l.acceptOne(alphas) {
		l.backup()
		l.emitError(fmt.Errorf("Unexpected token for identifier: %c", l.peak()))
		return nil
	}
	l.acceptMany(alphas + numbers)

	switch l.match() {
	case "type":
		l.emit(Type)
		l.ignoreSpace()
		return identifier
	case "func":
		l.emit(Function)
		l.ignoreSpace()
		return identifier
	}

	l.emit(Identifier)

	if l.acceptOne(".") {
		l.emit(Dot)
		return identifier
	}

	return root
}

func numberLiteral(l *lexer) lexFunc {
	l.acceptMany(numbers)
	if l.next() == '.' {
		l.acceptMany(numbers)
		l.emit(Float)
	} else {
		l.backup()
		l.emit(Integer)
	}
	return root
}

func stringLiteral(l *lexer) lexFunc {
	l.next()
	l.ignore()
	for c := l.next(); c != '"'; {
		if c == eof {
			l.emitError(fmt.Errorf("Unexpected end of file, string literal not closed"))
			l.emit(EOF)
			return nil
		}
		c = l.next()
	}
	l.backup()
	l.emit(String)
	l.next()
	l.ignore()
	return root
}
