package main

import (
	"fmt"

	"github.com/emilsjolander/hippo/check"
	"github.com/emilsjolander/hippo/gen"
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

# implement increment operator
(func ++:float f:float (+ f 1.0))

# executed when running script
(print (dot (vec2 1.0 1.0) (vec2 2.0 2.0)))
(print (++ 1.0))

`

func main() {
	lexemes, errored := lex.Lex(program)
	if errored {
		for _, l := range lexemes {
			fmt.Println(l)
		}
		return
	}

	root, errors := parse.Parse(lexemes)
	if errors != nil {
		for _, e := range errors {
			fmt.Println(e)
		}
		return
	}

	errors = check.Check(root)
	if errors != nil {
		for _, e := range errors {
			fmt.Println(e)
		}
		return
	}

	new(gen.JS).Write(root, "bin")
}
