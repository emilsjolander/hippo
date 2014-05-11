

function PLUS_float_float(one,two) {
	return one+two
}

function MINUS_float_float(one,two) {
	return one-two
}

function ASTERISK_float_float(one,two) {
	return one*two
}

function SLASH_float_float(one,two) {
	return one/two
}

function PLUS_int_int(one,two) {
	return one+two
}

function MINUS_int_int(one,two) {
	return one-two
}

function ASTERISK_int_int(one,two) {
	return one*two
}

function SLASH_int_int(one,two) {
	return one/two
}

function PLUS_string_string(one,two) {
	return one+two
}

function CARET_LEFT_int_int(one,two) {
	return one<two
}

function CARET_RIGHT_int_int(one,two) {
	return one>two
}

function EQUALS_int_int(one,two) {
	return one==two
}

function CARET_LEFT_float_float(one,two) {
	return one<two
}

function CARET_RIGHT_float_float(one,two) {
	return one>two
}

function EQUALS_float_float(one,two) {
	return one==two
}

function print_float(o) {
	console.log(o)
}

function print_int(o) {
	console.log(o)
}

function print_string(o) {
	console.log(o)
}

function print_bool(o) {
	console.log(o)
}

function vec2_float_float(x,y){return {x:x,y:y}}function dot_vec2_vec2(v1,v2){return PLUS_float_float(ASTERISK_float_float(v1.x,v2.x)
,ASTERISK_float_float(v1.y,v2.y)
)
}function fib_int(i){return (CARET_LEFT_int_int(i,2)
?1:PLUS_int_int(fib_int(MINUS_int_int(i,2)
)
,fib_int(MINUS_int_int(i,1)
)
)
)}function CARET_UP_int_int(i,e){return CARET_UPPRIM_int_int_int(i,i,e)
}function CARET_UPPRIM_int_int_int(v,i,e){return (EQUALS_int_int(e,0)
?1:(EQUALS_int_int(e,1)
?v:CARET_UPPRIM_int_int_int(ASTERISK_int_int(v,i)
,i,MINUS_int_int(e,1)
)
))}print_float(dot_vec2_vec2(vec2_float_float(1.0,1.0)
,vec2_float_float(2.0,2.0)
)
)
print_int(fib_int(10)
)
print_int(CARET_UP_int_int(10,3)
)
