Lanuage specification for hippo-lang
====================================
Hippo is a cool new functional language with type inferense

Keywords
--------
def		bind a type instance to a name
match	define a match block
->		equivalence
_		empty matcher
import	import a directory

Operators
---------
+		add two numbers
-		subtract two numbers
*		multiply two numbers	
/		divide two numbers
^		raise a number to the power of the other number	
(		opening parentheses
)		closing parentheses
//		single line comment
/*		multi line comment start
*/		multi line comment end

Type literals
-------------
string 		""
int			1
float		1.0
function 	[] ->

Builtin functions
-----------------
push
join
pop
head
tail
len
print

Code sample
-----------
import github.com/emilsjolander/stdlib
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
