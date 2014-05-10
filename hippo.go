package main

import (
	"fmt"

	"github.com/emilsjolander/hippo/ast"
	"github.com/emilsjolander/hippo/check"
	"github.com/emilsjolander/hippo/lex"
	"github.com/emilsjolander/hippo/parse"
)

const program = `

# type definition
(type vec2
	x:float
	y:float)

# function definition
(func dot:float v1:vec2 v2:vec2
	(+ 	(* v1.x v2.x)
		(* v1.y v2.y)))

# executed when running script
(print (dot (vec2 1.0 1.0) (vec2 2.0 2.0)))

`

func main() {
	lexemes, errored := lex.Lex(program)
	for _, l := range lexemes {
		fmt.Println(l)
	}
	if errored {
		return
	}

	astNode, errors := parse.Parse(lexemes)
	if errors != nil {
		for _, e := range errors {
			fmt.Println(e)
		}
		return
	}
	ast.Print(astNode)

	errors = check.Check(astNode)
	if errors != nil {
		for _, e := range errors {
			fmt.Println(e)
		}
		return
	}
}
