package parse

import (
	"fmt"

	"github.com/emilsjolander/hippo/lex"
)

type Error struct {
	Start lex.Pos
	End   lex.Pos
	Cause string
}

func (e Error) Error() string {
	return fmt.Sprintf("Error starting at %v, ending at %v", e.Start, e.End)
}
