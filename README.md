Lanuage specification for hippo-lang
====================================
Hippo is a toy project i'm working on.

Keywords
--------
| keyword       | description		       		 |
| ------------- |--------------------------------|
| import      	| import a directory             |
| def     		| bind a type instance to a name |
| match      	| define a match block           |

Operators
---------
| operator      | description		       		 |
| ------------- |--------------------------------|
| +          	| add two numbers             	 |
| -          	| subtract two numbers           |
| *          	| multiply two numbers           |
| /          	| divide two numbers             |
| ^          	| raise a number to another      |
| (          	| opening parentheses            |
| )          	| closing parentheses            |
| //          	| single line comment            |
| /*         	| multi line comment start       |
| */         	| multi line comment end         |
| _          	| empty matcher             	 |

Type literals
-------------
| type       	| example		       		 	 |
| ------------- |--------------------------------|
| string      	| ""             				 |
| int     		| 1 							 |
| float      	| 1.0           				 |
| function      | [] ->           				 |

Builtin functions
-----------------
| function      | description		       		 |
| ------------- |--------------------------------|
| push      	| push an element onto a list    |
| join     		| join two lists 				 |
| pop      		| pop an element from a list     |
| head      	| get the head of a list         |
| tail      	| get the tail of a list         |
| len      		| get the length of a list       |
| print      	| print a value				     |


Code sample
-----------
```
import github.com/emilsjolander/hippo/stdlib
import helpers

def pi 3.14159

def fib [n] ->
	match n:
		0 	-> 0
		1 	-> 1
		_ 	-> fib(n-1) + fib(n-2)

def map [l, f] ->
	match len(l):
		0 	-> []
		1 	-> [f(head(l))]
		_ 	-> push (f(head(l)), map(tail(l), f))
```
