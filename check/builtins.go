package check

import "github.com/emilsjolander/hippo/ast"

var (
	builtinTypes = []*typ{
		&typ{
			name: "float",
		},
		&typ{
			name: "int",
		},
		&typ{
			name: "string",
		},
	}
	builtinFuncs = []*function{
		&function{
			name: "+",
			typ:  "float",
			args: []ast.Property{
				ast.Property{
					Typ: "float",
				},
				ast.Property{
					Typ: "float",
				},
			},
		},
		&function{
			name: "*",
			typ:  "float",
			args: []ast.Property{
				ast.Property{
					Typ: "float",
				},
				ast.Property{
					Typ: "float",
				},
			},
		},
		&function{
			name: "print",
			typ:  "_",
			args: []ast.Property{
				ast.Property{
					Typ: "float",
				},
			},
		},
	}
)
