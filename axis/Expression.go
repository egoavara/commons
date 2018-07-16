package axis

import (
	"regexp"
	"strings"
)

// Both 		= ($horizontal|$vertical)
// Horizontal 	= ($horizontal)
// Vertical 	= ($vertical)
// None 		= ()
// Invalid 		= $else, <Invalid>
//
// $horizontal 	= "h", "H", "hori", "Hori", "horizontal", "Horizontal", "HORIZONTAL" ...
// 				= any letter start with 'h'
//
// $vertical 	= "v", "V", "vert", "Vert", "vertical", "Vertical", "VERTICAL" ...
// 				= any letter start with 'v'

var re_exp = regexp.MustCompile(`\(([a-zA-Z]+)?(\s*\|\s*[a-zA-Z]+)*\)`)

func Marshal(a Axis) string {
	switch a {
	case Both:
		return "(hori|vert)"
	case Vertical:
		return "(vert)"
	case Horizontal:
		return "(hori)"
	case None:
		return "()"
	}
	return "<Invalid>"
}

func Unmarshal(v string) (a Axis) {
	if re_exp.MatchString(v) {
		v = v[1 : len(v)-1]
		if len(v) == 0 {
			return None
		}
		for _, temp := range strings.Split(v, "|") {
			switch strings.TrimSpace(temp)[0] {
			case 'h':
				a |= Horizontal
			case 'v':
				a |= Vertical
			default:
				return Invalid
			}
		}
		return a
	}
	return Invalid
}
