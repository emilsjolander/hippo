

function PLUS_float_float(one,two) {
	return one+two
}

function MINUS_float_float(one,two) {
	return one+two
}

function ASTERISK_float_float(one,two) {
	return one*two
}

function SLASH_float_float(one,two) {
	return one+two
}

function PLUS_int_int(one,two) {
	return one+two
}

function MINUS_int_int(one,two) {
	return one+two
}

function ASTERISK_int_int(one,two) {
	return one*two
}

function SLASH_int_int(one,two) {
	return one+two
}

function PLUS_string_string(one,two) {
	return one+two
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

function vec2_float_float(x,y){return {x:x,y:y}}function dot_vec2_vec2(v1,v2){return PLUS_float_float(ASTERISK_float_float(v1.x,v2.x)
,ASTERISK_float_float(v1.y,v2.y)
)
}function PLUSPLUS_float(f){return PLUS_float_float(f,1.0)
}print_string(PLUS_string_string("hello ","world")
)
print_float(dot_vec2_vec2(vec2_float_float(1.0,1.0)
,vec2_float_float(2.0,2.0)
)
)
print_float(PLUSPLUS_float(1.0)
)
