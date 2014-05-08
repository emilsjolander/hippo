package main

import (
	"fmt"
	"github.com/emilsjolander/hippo/lex"
	"github.com/emilsjolander/hippo/parse"
)

const program = `

(type vec2
	x:float
	y:float)

(func dot:float v1:vec2 v2:vec2
	(+ (* v1.x v2.x)
		 (* v1.y v2.y)))

(print (dot (vec2 1 1) (vec2 2 2)))

`

func main() {
	lexemes := lex.Lex(program)
	for l := range lexemes {
		fmt.Println(l)
	}

	lexemes = lex.Lex(program)
	ast := parse.Parse(lexemes)
	parse.PrintAst(ast)
}
