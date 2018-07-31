package version

import (
	"regexp"
	"strconv"
)

// INTEGER 	::= ([1-9][0-9]*|0)
// VERSION 	::= INTEGER "." INTEGER "." INTEGER

var re_exp = regexp.MustCompile(`^(?P<MAJOR>[1-9][0-9]*|0)(\.(?P<MINOR>[1-9][0-9]*|0)(\.(?P<PATCH>[1-9][0-9]*|0))?)?$`)

func Marshal(a Version) string {
	return a.String()
}

func Unmarshal(v string) (Version) {
	res := re_exp.FindStringSubmatch(v)
	var err error
	var mj, mn, pc int64
	switch len(res) {
	default:
		fallthrough
	case 0:
		fallthrough
	case 1:
		return Invalid
	case 6:
		if len(res[5]) != 0{
			pc, err = strconv.ParseInt(res[5], 10, 32)
			if err != nil {
				return Invalid
			}
		}
		if len(res[3]) != 0{
			mn, err = strconv.ParseInt(res[3], 10, 32)
			if err != nil {
				return Invalid
			}
		}
		mj, err = strconv.ParseInt(res[1], 10, 32)
		if err != nil {
			return Invalid
		}
	}
	return Version{
		Major:int(mj),
		Minor:int(mn),
		Patch:int(pc),
	}
}
