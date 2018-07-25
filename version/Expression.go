package axis

import (
	"regexp"
	"strconv"
)

// INTEGER 	::= ([1-9][0-9]*|0)
// VERSION 	::= INTEGER "." INTEGER "." INTEGER

var re_exp = regexp.MustCompile(`^(?P<MAJOR>[1-9][0-9]*|0)\.(?P<MINOR>[1-9][0-9]*|0)\.(?P<PATCH>[1-9][0-9]*|0)$`)

func Marshal(a Version) string {
	return a.String()
}

func Unmarshal(v string) (Version) {
	res := re_exp.FindStringSubmatch(v)
	if len(res) == 0{
		return Invalid
	}
	mj, err := strconv.ParseInt(res[1], 10, 32)
	if err != nil {
		return Invalid
	}
	mn, err := strconv.ParseInt(res[2], 10, 32)
	if err != nil {
		return Invalid
	}
	pc, err := strconv.ParseInt(res[3], 10, 32)
	if err != nil {
		return Invalid
	}
	return Version{
		Major:int(mj),
		Minor:int(mn),
		Patch:int(pc),
	}
}
