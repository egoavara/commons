package axis

import "fmt"

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
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

