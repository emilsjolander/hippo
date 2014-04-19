package lex

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type TokenType int

const (
	TokenError TokenType = iota
	TokenOpenParen
	TokenCloseParen
	TokenOperator
	TokenNumber
	TokenEOF
)

type Token struct {
	Typ TokenType
	Val string
}

func (t Token) String() string {
	switch t.Typ {
	case TokenEOF:
		return "EOF"
	case TokenError:
		return t.Val
	}
	if len(t.Val) > 10 {
		return fmt.Sprintf("%.10q...", t.Val)
	}
	return t.Val
}

type lexer struct {
	input  string
	start  int
	pos    int
	width  int
	tokens chan Token
}

const eof = rune(-1)

type lexFunc func(*lexer) lexFunc

func Lex(input string) chan Token {
	l := &lexer{
		input:  input,
		tokens: make(chan Token),
	}
	go func() {
		for state := expression; state != nil; {
			// skip white space
			l.acceptMany(" ")
			l.ignore()

			if l.peak() == eof {
				l.emit(TokenEOF)
				break
			}

			state = state(l)
		}
		close(l.tokens)
	}()
	return l.tokens
}

func (l *lexer) emit(typ TokenType) {
	l.tokens <- Token{typ, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) emitError(err error) {
	l.ignore()
	l.tokens <- Token{TokenError, err.Error()}
}

func (l *lexer) next() rune {
	for _, r := range l.input[l.pos:] {
		l.width = utf8.RuneLen(r)
		l.pos += l.width
		return r
	}
	l.width = 0
	return eof
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peak() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) acceptOne(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptMany(valid string) int {
	n := 0
	for l.acceptOne(valid) {
		n++
	}
	return n
}
