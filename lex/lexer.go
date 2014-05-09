package lex

import (
	"strings"
	"unicode/utf8"
)

type lexer struct {
	input string

	start int
	pos   int
	width int

	startPos Pos

	lexemes []Lexeme
	errored bool
}

const eof = rune(-1)

type lexFunc func(*lexer) lexFunc

func Lex(input string) ([]Lexeme, bool) {
	l := &lexer{
		input:    input,
		lexemes:  make([]Lexeme, 0),
		startPos: Pos{1, 1},
	}
	for state := root; state != nil; {
		if l.peak() == eof {
			l.emit(EOF)
			break
		}
		state = state(l)
	}
	return l.lexemes, l.errored
}

func (l *lexer) updateStartPos() {
	val := l.input[l.start:l.pos]

	if newLineCount := strings.Count(val, "\n"); newLineCount > 0 {
		l.startPos.Row += strings.Count(val, "\n")
		l.startPos.Col = len(val) - strings.LastIndex(val, "\n")
	} else {
		l.startPos.Col += len(val)
	}

	l.start = l.pos
}

func (l *lexer) match() string {
	return l.input[l.start:l.pos]
}

func (l *lexer) emit(tok Token) {
	l.lexemes = append(l.lexemes, Lexeme{
		Tok:   tok,
		Val:   l.match(),
		Start: l.startPos,
	})
	l.updateStartPos()
}

func (l *lexer) emitError(err error) {
	l.lexemes = append(l.lexemes, Lexeme{
		Tok:   Error,
		Val:   err.Error(),
		Start: l.startPos,
	})
	l.errored = true
	l.updateStartPos()
}

func (l *lexer) ignoreSpace() {
	l.acceptMany(" \n\r\t")
	l.ignore()
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

func (l *lexer) backup() {
	l.pos -= l.width
	l.width = 0
}

func (l *lexer) ignore() {
	l.updateStartPos()
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
	totalWidth := 0
	for l.acceptOne(valid) {
		totalWidth += l.width
		n++
	}
	l.width = totalWidth
	return n
}
