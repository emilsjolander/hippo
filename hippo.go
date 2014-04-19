package main

import (
	"fmt"

	"github.com/emilsjolander/hippo/parse"

	"github.com/emilsjolander/hippo/lex"
)

func main() {
	tokens := lex.Lex("5+5*(2-4)^2")
	root := parse.Parse(tokens)
	fmt.Println(root.Value())
}
