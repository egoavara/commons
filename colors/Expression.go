package colors

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"image/color"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Hexadecimal
	re_rgb      = regexp.MustCompile(`#(?P<r>[0-9a-zA-Z])(?P<g>[0-9a-zA-Z])(?P<b>[0-9a-zA-Z])`)
	re_rgba     = regexp.MustCompile(`#(?P<r>[0-9a-zA-Z])(?P<g>[0-9a-zA-Z])(?P<b>[0-9a-zA-Z])(?P<a>[0-9a-zA-Z])`)
	re_rrggbb   = regexp.MustCompile(`#(?P<r>[0-9a-zA-Z]{2})(?P<g>[0-9a-zA-Z]{2})(?P<b>[0-9a-zA-Z]{2})`)
	re_rrggbbaa = regexp.MustCompile(`#(?P<r>[0-9a-zA-Z]{2})(?P<g>[0-9a-zA-Z]{2})(?P<b>[0-9a-zA-Z]{2})(?P<a>[0-9a-zA-Z]{2})`)
	// RGBA Constructor
	re_constructor_rgb  = regexp.MustCompile(`rgb\(\s*(?P<r>.+)\s*,\s*(?P<g>.+)\s*,\s*(?P<b>.+)\s*\)`)
	re_constructor_rgba = regexp.MustCompile(`rgba\(\s*(?P<r>.+)\s*,\s*(?P<g>.+)\s*,\s*(?P<b>.+)\s*,\s*(?P<a>.+)\s*\)`)
	// HSLA Constructor
	re_constructor_hsl  = regexp.MustCompile(`hsl\(\s*(?P<h>.+)\s*,\s*(?P<s>.+)\s*,\s*(?P<l>.+)\s*\)`)
	re_constructor_hsla = regexp.MustCompile(`hsla\(\s*(?P<h>.+)\s*,\s*(?P<s>.+)\s*,\s*(?P<l>.+)\s*,\s*(?P<a>.+)\s*\)`)
)

var MarshalFail = ""
var UnmarshalFail color.Color = nil

type Expression uint8

const (
	ExpressionHex             Expression = iota
	ExpressionRGBA            Expression = iota
	ExpressionHSLA            Expression = iota
	ExpressionPredefinedHTML  Expression = iota
	ExpressionPredefinedLaTeX Expression = iota
	ExpressionUnknown         Expression = math.MaxUint8
)

func Marshal(c color.Color, expression Expression, optimization bool) string {
	if c == nil {
		return MarshalFail
	}
	switch expression {
	default:
		fallthrough
	case ExpressionHex:
		r, g, b, a := c.RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		a >>= 8
		if optimization && a == math.MaxUint8 {
			if optimization && r>>4 == r&0x0F && g>>4 == g&0x0F && b>>4 == b&0x0F {
				return fmt.Sprintf("#%01x%01x%01x", r>>4, g>>4, b>>4)
			}
			return fmt.Sprintf("#%02x%02x%02x", r, g, b)
		} else {
			if optimization && r>>4 == r&0x0F && g>>4 == g&0x0F && b>>4 == b&0x0F && a>>4 == a&0x0F {
				return fmt.Sprintf("#%01x%01x%01x%01x", r>>4, g>>4, b>>4, a>>4)
			}
			return fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
		}
	case ExpressionRGBA:
		r, g, b, a := c.RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		if optimization && a == math.MaxUint16 {
			return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
		} else {
			return fmt.Sprintf("rgba(%d, %d, %d, %s)", r, g, b, strconv.FormatFloat(float64(a)/math.MaxUint16, 'f', -1, 32))
		}
	case ExpressionHSLA:
		hsla := HSLAModel.Convert(c).(HSLA)
		if optimization && hsla.A == 1 {
			return fmt.Sprintf("hsl(%d, %d%%, %d%%)", uint32(mgl32.Clamp(hsla.H*360, 0, 360)), uint32(hsla.S*100), uint32(hsla.L*100))
		} else {
			return fmt.Sprintf("hsla(%d, %d%%, %d%%, %s)", uint32(mgl32.Clamp(hsla.H*360, 0, 360)), uint32(hsla.S*100), uint32(hsla.L*100), strconv.FormatFloat(float64(hsla.A), 'f', -1, 32))
		}
	case ExpressionPredefinedHTML:
		if cc, ok := c.(U32); ok{
			return HTML.ToString(cc)
		}
	case ExpressionPredefinedLaTeX:
		if cc, ok := c.(U32); ok{
			return LaTeX.ToString(cc)
		}
	}
	return MarshalFail
}
func Unmarshal(value string, define ... Predefine) (color.Color, Expression) {
	// hex type color
	if res := re_rrggbb.FindStringSubmatch(value); len(res) > 0 {
		r, err := strconv.ParseUint(res[1], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		g, err := strconv.ParseUint(res[2], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		b, err := strconv.ParseUint(res[3], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		return U32{
			uint8(r),
			uint8(g),
			uint8(b),
			0xFF,
		}, ExpressionHex
	}
	if res := re_rrggbbaa.FindStringSubmatch(value); len(res) > 0 {
		r, err := strconv.ParseUint(res[1], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		g, err := strconv.ParseUint(res[2], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		b, err := strconv.ParseUint(res[3], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		a, err := strconv.ParseUint(res[4], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		return U32{
			uint8(r),
			uint8(g),
			uint8(b),
			uint8(a),
		}, ExpressionHex
	}
	if res := re_rgba.FindStringSubmatch(value); len(res) > 0 {
		r, err := strconv.ParseUint(res[1], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		g, err := strconv.ParseUint(res[2], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		b, err := strconv.ParseUint(res[3], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		a, err := strconv.ParseUint(res[4], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		return U32{
			uint8(r * 0x11),
			uint8(g * 0x11),
			uint8(b * 0x11),
			uint8(a * 0x11),
		}, ExpressionHex
	}
	if res := re_rgb.FindStringSubmatch(value); len(res) > 0 {
		r, err := strconv.ParseUint(res[1], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		g, err := strconv.ParseUint(res[2], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		b, err := strconv.ParseUint(res[3], 16, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		return U32{
			uint8(r * 0x11),
			uint8(g * 0x11),
			uint8(b * 0x11),
			0xFF,
		}, ExpressionHex
	}
	// constructor type color
	if res := re_constructor_rgb.FindStringSubmatch(value); len(res) > 0 {
		r, err := strconv.ParseUint(res[1], 10, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		g, err := strconv.ParseUint(res[2], 10, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		b, err := strconv.ParseUint(res[3], 10, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		return U32{
			uint8(r),
			uint8(g),
			uint8(b),
			0xFF,
		}, ExpressionRGBA
	}
	if res := re_constructor_rgba.FindStringSubmatch(value); len(res) > 0 {
		r, err := strconv.ParseUint(res[1], 10, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		g, err := strconv.ParseUint(res[2], 10, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		b, err := strconv.ParseUint(res[3], 10, 8)
		if err != nil {
			return UnmarshalFail, ExpressionUnknown
		}
		a, err := strconv.ParseFloat(res[4], 32)
		if err != nil || a < 0 || a > 1 {
			return UnmarshalFail, ExpressionUnknown
		}
		return U32{
			uint8(r),
			uint8(g),
			uint8(b),
			uint8(a*math.MaxUint8 + .5),
		}, ExpressionRGBA
	}
	if res := re_constructor_hsl.FindStringSubmatch(value); len(res) > 0 {
		h, err := strconv.ParseUint(res[1], 10, 64)
		if err != nil || h > 360 {
			return UnmarshalFail, ExpressionUnknown
		}
		s, err := strconv.ParseUint(strings.TrimSuffix(res[2], "%"), 10, 64)
		if err != nil || s > 100 {
			return UnmarshalFail, ExpressionUnknown
		}
		l, err := strconv.ParseUint(strings.TrimSuffix(res[3], "%"), 10, 64)
		if err != nil || l > 100 {
			return UnmarshalFail, ExpressionUnknown
		}
		return HSLA{
			float32(h) / 360,
			float32(s) / 100,
			float32(l) / 100,
			1,
		}, ExpressionHSLA
	}
	if res := re_constructor_hsla.FindStringSubmatch(value); len(res) > 0 {
		h, err := strconv.ParseUint(res[1], 10, 64)
		if err != nil || h > 360 {
			return UnmarshalFail, ExpressionUnknown
		}
		s, err := strconv.ParseUint(strings.TrimSuffix(res[2], "%"), 10, 64)
		if err != nil || s > 100 {
			return UnmarshalFail, ExpressionUnknown
		}
		l, err := strconv.ParseUint(strings.TrimSuffix(res[3], "%"), 10, 64)
		if err != nil || l > 100 {
			return UnmarshalFail, ExpressionUnknown
		}
		a, err := strconv.ParseFloat(res[4], 32)
		if err != nil || a < 0 || a > 1 {
			return UnmarshalFail, ExpressionUnknown
		}
		return HSLA{
			float32(h) / 360,
			float32(s) / 100,
			float32(l) / 100,
			float32(a),
		}, ExpressionHSLA
	}
	// predefined color
	for _, def := range define {
		if res := def.Find(value); res != nil{
			return *res, def.Type()
		}
	}
	return UnmarshalFail, ExpressionUnknown
}
