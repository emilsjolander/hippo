

function PLUS_float_float(f1,f2) {
	return f1+f2
}

function ASTERISK_float_float(f1,f2) {
	return f1*f2
}

function print_float(f) {
	console.log(f)
}

function vec2_float_float(x,y){return {x:x,y:y}}function dot_vec2_vec2(v1,v2){return PLUS_float_float(ASTERISK_float_float(v1.x,v2.x),ASTERISK_float_float(v1.y,v2.y))}print_float(dot_vec2_vec2(vec2_float_float(1.0,1.0),vec2_float_float(2.0,2.0)))