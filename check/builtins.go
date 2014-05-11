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
		&typ{
			name: "bool",
		},
		&typ{
			name: "void",
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
			name: "-",
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
			name: "/",
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
			name: "+",
			typ:  "int",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "-",
			typ:  "int",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "*",
			typ:  "int",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "/",
			typ:  "int",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "<",
			typ:  "bool",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: ">",
			typ:  "bool",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "=",
			typ:  "bool",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "<",
			typ:  "bool",
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
			name: ">",
			typ:  "bool",
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
			name: "=",
			typ:  "bool",
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
			name: "+",
			typ:  "string",
			args: []ast.Property{
				ast.Property{
					Typ: "string",
				},
				ast.Property{
					Typ: "string",
				},
			},
		},
		&function{
			name: "print",
			typ:  "void",
			args: []ast.Property{
				ast.Property{
					Typ: "float",
				},
			},
		},
		&function{
			name: "print",
			typ:  "void",
			args: []ast.Property{
				ast.Property{
					Typ: "int",
				},
			},
		},
		&function{
			name: "print",
			typ:  "void",
			args: []ast.Property{
				ast.Property{
					Typ: "string",
				},
			},
		},
		&function{
			name: "print",
			typ:  "void",
			args: []ast.Property{
				ast.Property{
					Typ: "bool",
				},
			},
		},
	}
)
