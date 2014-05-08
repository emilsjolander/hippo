package lex

import (
	"fmt"
)

type Pos struct {
	Row int
	Col int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.Row, p.Col)
}
