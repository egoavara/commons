package align

type Align uint32

func (s Align) Horizontal() Align {
	return s & MaskHorizontal
}
func (s Align) Vertical() Align {
	return s & MaskVertical
}

const (
	Left           = Align(0x0000)
	Horizontal     = Align(0x0001)
	Right          = Align(0x0002)
	MaskHorizontal = Align(0x000F)

	Top          = Align(0x0000)
	Vertical     = Align(0x0010)
	Bottom       = Align(0x0020)
	MaskVertical = Align(0x00F0)

	Center = Horizontal | Vertical
)
