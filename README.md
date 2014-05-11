Lanuage specification for hippo-lang
====================================
Hippo is a toy project i'm working on. It's a statically typed language with inspirations from lisp.

Keywords
--------
| keyword       | description                      |
| ------------- | -------------------------------- |
| type      	  | define a type                  |
| func      	  | define a function              |

Tokens
------
| type         	| token     		       		   |
| ------------- | -------------------------------- |
| Colon    		| :                                |
| OpenParen     | (                                |
| CloseParen    | )                                |
| Comment       | #                                |

Primitives
----------
| type       	| example		       		 	   |
| ------------- | -------------------------------- |
| string      	| "1"            		           |
| int      	   	| 1 							   |
| float       	| 1.0 						       |

Code sample
-----------
```lisp

# type definition
(type vec2
	x:float
	y:float)

# function definition
(func dot:float v1:vec2 v2:vec2
	(+ 	(* v1.x v2.x)
		(* v1.y v2.y)))

# fibonacci function
(func fib:int i:int
	(if (< i 2) 1
		(+  (fib (- i 2))
			(fib (- i 1)))))

# pow operator
(func ^:int i:int e:int (^' i i e))
(func ^':int v:int i:int e:int
	(if (= e 0) 1 
		(if (= e 1) v
			(^' (* v i) i (- e 1)))))

# executed when running script
(print (dot (vec2 1.0 1.0) (vec2 2.0 2.0)))
(print (fib 10))
(print (^ 10 3))

```
