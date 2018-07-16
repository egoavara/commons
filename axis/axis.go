package axis

type Axis uint32

func (s Axis) HasHorizontal() bool {
	return s&Horizontal == Horizontal
}
func (s Axis) HasVertical() bool {
	return s&Vertical == Vertical
}
func (s Axis) HasBoth() bool {
	return s&Vertical == Vertical && s&Horizontal == Horizontal
}

func (s Axis) Has() (vert, hori bool) {
	return s&Vertical == Vertical, s&Horizontal == Horizontal
}

type ExculsiveAxis Axis

func (s ExculsiveAxis) Axis() Axis {
	switch s {
	case ExculsiveVertical:
		fallthrough
	case ExculsiveHorizontal:
		return Axis(s)
	}
	return Invalid
}

const (
	None       = Axis(0x0000)
	Horizontal = Axis(0x0001)
	Vertical   = Axis(0x0002)
	Both       = Axis(0x0003)
	Invalid    = Axis(0xF0)
)
const (
	ExculsiveHorizontal = ExculsiveAxis(Horizontal)
	ExculsiveVertical   = ExculsiveAxis(Vertical)
)
