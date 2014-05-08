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

(type vec3
	x:float
	y:float
	z:float)

(func dot:float v1:vec3 v2:vec3
	(+ (* v1.x v2.x)
		 (* v1.y v2.y)
		 (* v1.z v2.z)))

(print (dot (vec3 1 1 1) (vec3 2 2 2)))

```
