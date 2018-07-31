package version

import (
	"strconv"
)

var Invalid = Version{
	-1,-1, -1,
}

type Version struct {
	Major int
	Minor int
	Patch int
}


func (s Version) String() string {
	if s == Invalid{
		return "Invalid"
	}
	res := strconv.FormatInt(int64(s.Major), 10) + "." + strconv.FormatInt(int64(s.Minor), 10)
	if s.Patch != 0{
		res += "." + strconv.FormatInt(int64(s.Patch), 10)
	}
	return res
}

