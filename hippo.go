package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/emilsjolander/hippo/check"
	"github.com/emilsjolander/hippo/gen"
	"github.com/emilsjolander/hippo/lex"
	"github.com/emilsjolander/hippo/parse"
)

/*
	TODO:
		- All nodes should have a types, add Type() to the node interface
		- Nodes without type should have the builtin 'void' type
		- Remove types property from expression node. Types can be taken from arg nodes with new Type() method

*/

func main() {
	scriptName := os.Args[1]
	script, err := ioutil.ReadFile(scriptName + ".hippo")
	if err != nil {
		panic(err)
	}

	lexemes, errored := lex.Lex(string(script))
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

	new(gen.JS).Write(root, "bin", scriptName)
}
