Lanuage specification for hippo-lang
====================================
Hippo is a toy project i'm working on. It's a typed language in the list family.

Keywords
--------
| keyword       | description                      |
| ------------- | -------------------------------- |
| type      	  | define a type                    |
| func      	  | define a function                |

Tokens
------
| token         | regex     		       		        |
| ------------- | -------------------------------- |
| Space    		 | [ \n]+                           |
| Colon    		 | :                                |
| OpenParen     | (                                |
| CloseParen    | )                                |
| Comment       | //.*                             |

Primitives
----------
| type       	 | example		       		 	        |
| ------------- | -------------------------------- |
| string      	| "1"            		         		 |
| int      	   | 1 							                 |
| float       	| 1.0 						                 |

Code sample
-----------
```lisp

(type vec2
	x:float
	y:float)

(func dot:float v1:vec2 v2:vec2
	(+ 	(* v1.x v2.x)
		(* v1.y v2.y)))

(print (dot (vec2 1 1) (vec2 2 2)))

```
