package parse

import "math"

type Node interface {
	Value() float64
}

type number struct {
	val float64
}

func (n *number) Value() float64 {
	return n.val
}

type operator struct {
	sign  string
	left  Node
	right Node
}

func (o *operator) Value() float64 {
	l := o.left.Value()
	r := o.right.Value()

	switch o.sign {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	case "^":
		return math.Pow(l, r)
	}

	panic("invalid operator type")
}

func (o *operator) precedence() int {
	switch o.sign {
	case "+":
		return 0
	case "-":
		return 0
	case "*":
		return 1
	case "/":
		return 1
	case "^":
		return 2
	}

	panic("invalid operator type")
}
