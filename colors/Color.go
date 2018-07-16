package colors

import (
	"github.com/go-gl/mathgl/mgl32"
	"image/color"
	"math"
)

var (
	U32Model = color.ModelFunc(func(c color.Color) color.Color {
		if _, ok := c.(U32); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return U32{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
	})
	CompressedU32Model = color.ModelFunc(func(c color.Color) color.Color {
		if _, ok := c.(CompressedU32); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return CompressedU32(r<<24 | g<<16 | b<<8 | a)
	})
	U64Model = color.ModelFunc(func(c color.Color) color.Color {
		if _, ok := c.(U64); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return U64{uint16(r), uint16(g), uint16(b), uint16(a)}
	})
	CompressedU64Model = color.ModelFunc(func(c color.Color) color.Color {
		if _, ok := c.(CompressedU64); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return CompressedU32(r<<48 | g<<32 | b<<16 | a)
	})
	F128Model = color.ModelFunc(func(c color.Color) color.Color {
		if _, ok := c.(F128); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		return F128{float32(r) / math.MaxUint16, float32(g) / math.MaxUint16, float32(b) / math.MaxUint16, float32(a) / math.MaxUint16}
	})
	HSLAModel = color.ModelFunc(func(c color.Color) color.Color {
		if _, ok := c.(HSLA); ok {
			return c
		}
		r, g, b, a := c.RGBA()
		var res HSLA
		res.A = mgl32.Clamp(float32(a)/math.MaxUint16, 0, 1)
		//
		nr := mgl32.Clamp(float32(r)/math.MaxUint16, 0, 1)
		ng := mgl32.Clamp(float32(g)/math.MaxUint16, 0, 1)
		nb := mgl32.Clamp(float32(b)/math.MaxUint16, 0, 1)
		//
		min := float32(math.Min(math.Min(float64(nr), float64(ng)), float64(nb)))
		max := float32(math.Max(math.Max(float64(nr), float64(ng)), float64(nb)))
		res.L = (max + min) / 2
		if min == max {
			return res
		}
		var delta = max - min
		if res.L > .5 {
			res.S = delta / (2 - max - min)
		} else {
			res.S = delta / (max + min)
		}
		switch max {
		case nr:
			if ng > nb {
				res.H = (ng-nb)/delta + 0
			} else {
				res.H = (ng-nb)/delta + 6
			}
		case ng:
			res.H = (nb-nr)/delta + 2
		case nb:
			res.H = (nr-ng)/delta + 4
		}
		res.H /= 6
		return res
	})
)

type U32 struct {
	R, G, B, A uint8
}
func (s U32) RGBA() (r, g, b, a uint32) {
	r = uint32(s.R)
	r |= r << 8
	g = uint32(s.G)
	g |= g << 8
	b = uint32(s.B)
	b |= b << 8
	a = uint32(s.A)
	a |= a << 8
	return
}

type CompressedU32 uint32
func (s CompressedU32) RGBA() (r, g, b, a uint32) {
	r = uint32(s>>24) & 0xFF
	g = uint32(s>>16) & 0xFF
	b = uint32(s>>8) & 0xFF
	a = uint32(s>>0) & 0xFF
	return
}

type U64 struct {
	R, G, B, A uint16
}
func (s U64) RGBA() (r, g, b, a uint32) {
	r = uint32(s.R)
	g = uint32(s.G)
	b = uint32(s.B)
	a = uint32(s.A)
	return
}

type CompressedU64 uint64
func (s CompressedU64) RGBA() (r, g, b, a uint32) {
	r = uint32(s>>48) & 0xFFFF
	g = uint32(s>>32) & 0xFFFF
	b = uint32(s>>16) & 0xFFFF
	a = uint32(s>>0) & 0xFFFF
	return
}

type F128 struct {
	R, G, B, A float32
}
func (s F128) RGBA() (r, g, b, a uint32) {
	return uint32(s.R * math.MaxUint16), uint32(s.G * math.MaxUint16), uint32(s.B * math.MaxUint16), uint32(s.A * math.MaxUint16)
}

type HSLA struct {
	H float32
	S float32
	L float32
	A float32
}
func (s HSLA) RGBA() (r, g, b, a uint32) {
	a = uint32(s.A * math.MaxUint16)
	if s.S <= 0.001 {
		out := uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
		return out, out, out, a
	}

	i := floor(s.H * 6)
	f := s.H*6 - i
	p := s.L * (1 - s.S)
	q := s.L * (1 - s.S*f)
	t := s.L * (1 - s.S*(1-f))

	switch i {
	case 0:
		r = uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
		g = uint32(mgl32.Clamp(t*math.MaxUint16+0.5, 0, math.MaxUint16))
		b = uint32(mgl32.Clamp(p*math.MaxUint16+0.5, 0, math.MaxUint16))
	case 1:
		r = uint32(mgl32.Clamp(q*math.MaxUint16+0.5, 0, math.MaxUint16))
		g = uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
		b = uint32(mgl32.Clamp(p*math.MaxUint16+0.5, 0, math.MaxUint16))
	case 2:
		r = uint32(mgl32.Clamp(p*math.MaxUint16+0.5, 0, math.MaxUint16))
		g = uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
		b = uint32(mgl32.Clamp(t*math.MaxUint16+0.5, 0, math.MaxUint16))
	case 3:
		r = uint32(mgl32.Clamp(p*math.MaxUint16+0.5, 0, math.MaxUint16))
		g = uint32(mgl32.Clamp(q*math.MaxUint16+0.5, 0, math.MaxUint16))
		b = uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
	case 4:
		r = uint32(mgl32.Clamp(t*math.MaxUint16+0.5, 0, math.MaxUint16))
		g = uint32(mgl32.Clamp(p*math.MaxUint16+0.5, 0, math.MaxUint16))
		b = uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
	default:
		r = uint32(mgl32.Clamp(s.L*math.MaxUint16+0.5, 0, math.MaxUint16))
		g = uint32(mgl32.Clamp(p*math.MaxUint16+0.5, 0, math.MaxUint16))
		b = uint32(mgl32.Clamp(q*math.MaxUint16+0.5, 0, math.MaxUint16))
	}
	return
}
func floor(a float32) float32 {
	return float32(math.Floor(float64(a)))
}
