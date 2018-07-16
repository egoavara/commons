package colors

import (
	"strings"
)

var HTML _HTML

type _HTML struct {
	_c map[string]*U32

	AliceBlue            *U32
	AntiqueWhite         *U32
	Aqua                 *U32
	Aquamarine           *U32
	Azure                *U32
	Beige                *U32
	Bisque               *U32
	Black                *U32
	BlanchedAlmond       *U32
	Blue                 *U32
	BlueViolet           *U32
	Brown                *U32
	BurlyWood            *U32
	CadetBlue            *U32
	Chartreuse           *U32
	Chocolate            *U32
	Coral                *U32
	CornflowerBlue       *U32
	Cornsilk             *U32
	Crimson              *U32
	Cyan                 *U32
	DarkBlue             *U32
	DarkCyan             *U32
	DarkGoldenRod        *U32
	DarkGray             *U32
	DarkGrey             *U32
	DarkGreen            *U32
	DarkKhaki            *U32
	DarkMagenta          *U32
	DarkOliveGreen       *U32
	DarkOrange           *U32
	DarkOrchid           *U32
	DarkRed              *U32
	DarkSalmon           *U32
	DarkSeaGreen         *U32
	DarkSlateBlue        *U32
	DarkSlateGray        *U32
	DarkSlateGrey        *U32
	DarkTurquoise        *U32
	DarkViolet           *U32
	DeepPink             *U32
	DeepSkyBlue          *U32
	DimGray              *U32
	DimGrey              *U32
	DodgerBlue           *U32
	FireBrick            *U32
	FloralWhite          *U32
	ForestGreen          *U32
	Fuchsia              *U32
	Gainsboro            *U32
	GhostWhite           *U32
	Gold                 *U32
	GoldenRod            *U32
	Gray                 *U32
	Grey                 *U32
	Green                *U32
	GreenYellow          *U32
	HoneyDew             *U32
	HotPink              *U32
	IndianRed            *U32
	Indigo               *U32
	Ivory                *U32
	Khaki                *U32
	Lavender             *U32
	LavenderBlush        *U32
	LawnGreen            *U32
	LemonChiffon         *U32
	LightBlue            *U32
	LightCoral           *U32
	LightCyan            *U32
	LightGoldenRodYellow *U32
	LightGray            *U32
	LightGrey            *U32
	LightGreen           *U32
	LightPink            *U32
	LightSalmon          *U32
	LightSeaGreen        *U32
	LightSkyBlue         *U32
	LightSlateGray       *U32
	LightSlateGrey       *U32
	LightSteelBlue       *U32
	LightYellow          *U32
	Lime                 *U32
	LimeGreen            *U32
	Linen                *U32
	Magenta              *U32
	Maroon               *U32
	MediumAquaMarine     *U32
	MediumBlue           *U32
	MediumOrchid         *U32
	MediumPurple         *U32
	MediumSeaGreen       *U32
	MediumSlateBlue      *U32
	MediumSpringGreen    *U32
	MediumTurquoise      *U32
	MediumVioletRed      *U32
	MidnightBlue         *U32
	MintCream            *U32
	MistyRose            *U32
	Moccasin             *U32
	NavajoWhite          *U32
	Navy                 *U32
	OldLace              *U32
	Olive                *U32
	OliveDrab            *U32
	Orange               *U32
	OrangeRed            *U32
	Orchid               *U32
	PaleGoldenRod        *U32
	PaleGreen            *U32
	PaleTurquoise        *U32
	PaleVioletRed        *U32
	PapayaWhip           *U32
	PeachPuff            *U32
	Peru                 *U32
	Pink                 *U32
	Plum                 *U32
	PowderBlue           *U32
	Purple               *U32
	RebeccaPurple        *U32
	Red                  *U32
	RosyBrown            *U32
	RoyalBlue            *U32
	SaddleBrown          *U32
	Salmon               *U32
	SandyBrown           *U32
	SeaGreen             *U32
	SeaShell             *U32
	Sienna               *U32
	Silver               *U32
	SkyBlue              *U32
	SlateBlue            *U32
	SlateGray            *U32
	SlateGrey            *U32
	Snow                 *U32
	SpringGreen          *U32
	SteelBlue            *U32
	Tan                  *U32
	Teal                 *U32
	Thistle              *U32
	Tomato               *U32
	Turquoise            *U32
	Violet               *U32
	Wheat                *U32
	White                *U32
	WhiteSmoke           *U32
	Yellow               *U32
	YellowGreen          *U32
}

func (s _HTML) Keys() []string {
	var keys = make([]string, 0, len(s._c))
	for k := range s._c {
		keys = append(keys, k)
	}
	return keys
}

func (s _HTML) Find(name string) *U32 {
	return s._c[strings.ToLower(name)]
}
func (s _HTML) Type() Expression {
	return ExpressionPredefinedHTML
}
func (s _HTML) ToString(c U32) string {
	for k, v := range s._c {
		if *v == c {
			return k
		}
	}
	return ""
}

// https://www.w3schools.com/Colors/colors_names.asp
func init() {
	HTML = _HTML{
		AliceBlue:            &U32{0xF0, 0xF8, 0xFF, 0xFF},
		AntiqueWhite:         &U32{0xFA, 0xEB, 0xD7, 0xFF},
		Aqua:                 &U32{0x00, 0xFF, 0xFF, 0xFF},
		Aquamarine:           &U32{0x7F, 0xFF, 0xD4, 0xFF},
		Azure:                &U32{0xF0, 0xFF, 0xFF, 0xFF},
		Beige:                &U32{0xF5, 0xF5, 0xDC, 0xFF},
		Bisque:               &U32{0xFF, 0xE4, 0xC4, 0xFF},
		Black:                &U32{0x00, 0x00, 0x00, 0xFF},
		BlanchedAlmond:       &U32{0xFF, 0xEB, 0xCD, 0xFF},
		Blue:                 &U32{0x00, 0x00, 0xFF, 0xFF},
		BlueViolet:           &U32{0x8A, 0x2B, 0xE2, 0xFF},
		Brown:                &U32{0xA5, 0x2A, 0x2A, 0xFF},
		BurlyWood:            &U32{0xDE, 0xB8, 0x87, 0xFF},
		CadetBlue:            &U32{0x5F, 0x9E, 0xA0, 0xFF},
		Chartreuse:           &U32{0x7F, 0xFF, 0x00, 0xFF},
		Chocolate:            &U32{0xD2, 0x69, 0x1E, 0xFF},
		Coral:                &U32{0xFF, 0x7F, 0x50, 0xFF},
		CornflowerBlue:       &U32{0x64, 0x95, 0xED, 0xFF},
		Cornsilk:             &U32{0xFF, 0xF8, 0xDC, 0xFF},
		Crimson:              &U32{0xDC, 0x14, 0x3C, 0xFF},
		Cyan:                 &U32{0x00, 0xFF, 0xFF, 0xFF},
		DarkBlue:             &U32{0x00, 0x00, 0x8B, 0xFF},
		DarkCyan:             &U32{0x00, 0x8B, 0x8B, 0xFF},
		DarkGoldenRod:        &U32{0xB8, 0x86, 0x0B, 0xFF},
		DarkGray:             &U32{0xA9, 0xA9, 0xA9, 0xFF},
		DarkGrey:             &U32{0xA9, 0xA9, 0xA9, 0xFF},
		DarkGreen:            &U32{0x00, 0x64, 0x00, 0xFF},
		DarkKhaki:            &U32{0xBD, 0xB7, 0x6B, 0xFF},
		DarkMagenta:          &U32{0x8B, 0x00, 0x8B, 0xFF},
		DarkOliveGreen:       &U32{0x55, 0x6B, 0x2F, 0xFF},
		DarkOrange:           &U32{0xFF, 0x8C, 0x00, 0xFF},
		DarkOrchid:           &U32{0x99, 0x32, 0xCC, 0xFF},
		DarkRed:              &U32{0x8B, 0x00, 0x00, 0xFF},
		DarkSalmon:           &U32{0xE9, 0x96, 0x7A, 0xFF},
		DarkSeaGreen:         &U32{0x8F, 0xBC, 0x8F, 0xFF},
		DarkSlateBlue:        &U32{0x48, 0x3D, 0x8B, 0xFF},
		DarkSlateGray:        &U32{0x2F, 0x4F, 0x4F, 0xFF},
		DarkSlateGrey:        &U32{0x2F, 0x4F, 0x4F, 0xFF},
		DarkTurquoise:        &U32{0x00, 0xCE, 0xD1, 0xFF},
		DarkViolet:           &U32{0x94, 0x00, 0xD3, 0xFF},
		DeepPink:             &U32{0xFF, 0x14, 0x93, 0xFF},
		DeepSkyBlue:          &U32{0x00, 0xBF, 0xFF, 0xFF},
		DimGray:              &U32{0x69, 0x69, 0x69, 0xFF},
		DimGrey:              &U32{0x69, 0x69, 0x69, 0xFF},
		DodgerBlue:           &U32{0x1E, 0x90, 0xFF, 0xFF},
		FireBrick:            &U32{0xB2, 0x22, 0x22, 0xFF},
		FloralWhite:          &U32{0xFF, 0xFA, 0xF0, 0xFF},
		ForestGreen:          &U32{0x22, 0x8B, 0x22, 0xFF},
		Fuchsia:              &U32{0xFF, 0x00, 0xFF, 0xFF},
		Gainsboro:            &U32{0xDC, 0xDC, 0xDC, 0xFF},
		GhostWhite:           &U32{0xF8, 0xF8, 0xFF, 0xFF},
		Gold:                 &U32{0xFF, 0xD7, 0x00, 0xFF},
		GoldenRod:            &U32{0xDA, 0xA5, 0x20, 0xFF},
		Gray:                 &U32{0x80, 0x80, 0x80, 0xFF},
		Grey:                 &U32{0x80, 0x80, 0x80, 0xFF},
		Green:                &U32{0x00, 0x80, 0x00, 0xFF},
		GreenYellow:          &U32{0xAD, 0xFF, 0x2F, 0xFF},
		HoneyDew:             &U32{0xF0, 0xFF, 0xF0, 0xFF},
		HotPink:              &U32{0xFF, 0x69, 0xB4, 0xFF},
		IndianRed:            &U32{0xCD, 0x5C, 0x5C, 0xFF},
		Indigo:               &U32{0x4B, 0x00, 0x82, 0xFF},
		Ivory:                &U32{0xFF, 0xFF, 0xF0, 0xFF},
		Khaki:                &U32{0xF0, 0xE6, 0x8C, 0xFF},
		Lavender:             &U32{0xE6, 0xE6, 0xFA, 0xFF},
		LavenderBlush:        &U32{0xFF, 0xF0, 0xF5, 0xFF},
		LawnGreen:            &U32{0x7C, 0xFC, 0x00, 0xFF},
		LemonChiffon:         &U32{0xFF, 0xFA, 0xCD, 0xFF},
		LightBlue:            &U32{0xAD, 0xD8, 0xE6, 0xFF},
		LightCoral:           &U32{0xF0, 0x80, 0x80, 0xFF},
		LightCyan:            &U32{0xE0, 0xFF, 0xFF, 0xFF},
		LightGoldenRodYellow: &U32{0xFA, 0xFA, 0xD2, 0xFF},
		LightGray:            &U32{0xD3, 0xD3, 0xD3, 0xFF},
		LightGrey:            &U32{0xD3, 0xD3, 0xD3, 0xFF},
		LightGreen:           &U32{0x90, 0xEE, 0x90, 0xFF},
		LightPink:            &U32{0xFF, 0xB6, 0xC1, 0xFF},
		LightSalmon:          &U32{0xFF, 0xA0, 0x7A, 0xFF},
		LightSeaGreen:        &U32{0x20, 0xB2, 0xAA, 0xFF},
		LightSkyBlue:         &U32{0x87, 0xCE, 0xFA, 0xFF},
		LightSlateGray:       &U32{0x77, 0x88, 0x99, 0xFF},
		LightSlateGrey:       &U32{0x77, 0x88, 0x99, 0xFF},
		LightSteelBlue:       &U32{0xB0, 0xC4, 0xDE, 0xFF},
		LightYellow:          &U32{0xFF, 0xFF, 0xE0, 0xFF},
		Lime:                 &U32{0x00, 0xFF, 0x00, 0xFF},
		LimeGreen:            &U32{0x32, 0xCD, 0x32, 0xFF},
		Linen:                &U32{0xFA, 0xF0, 0xE6, 0xFF},
		Magenta:              &U32{0xFF, 0x00, 0xFF, 0xFF},
		Maroon:               &U32{0x80, 0x00, 0x00, 0xFF},
		MediumAquaMarine:     &U32{0x66, 0xCD, 0xAA, 0xFF},
		MediumBlue:           &U32{0x00, 0x00, 0xCD, 0xFF},
		MediumOrchid:         &U32{0xBA, 0x55, 0xD3, 0xFF},
		MediumPurple:         &U32{0x93, 0x70, 0xDB, 0xFF},
		MediumSeaGreen:       &U32{0x3C, 0xB3, 0x71, 0xFF},
		MediumSlateBlue:      &U32{0x7B, 0x68, 0xEE, 0xFF},
		MediumSpringGreen:    &U32{0x00, 0xFA, 0x9A, 0xFF},
		MediumTurquoise:      &U32{0x48, 0xD1, 0xCC, 0xFF},
		MediumVioletRed:      &U32{0xC7, 0x15, 0x85, 0xFF},
		MidnightBlue:         &U32{0x19, 0x19, 0x70, 0xFF},
		MintCream:            &U32{0xF5, 0xFF, 0xFA, 0xFF},
		MistyRose:            &U32{0xFF, 0xE4, 0xE1, 0xFF},
		Moccasin:             &U32{0xFF, 0xE4, 0xB5, 0xFF},
		NavajoWhite:          &U32{0xFF, 0xDE, 0xAD, 0xFF},
		Navy:                 &U32{0x00, 0x00, 0x80, 0xFF},
		OldLace:              &U32{0xFD, 0xF5, 0xE6, 0xFF},
		Olive:                &U32{0x80, 0x80, 0x00, 0xFF},
		OliveDrab:            &U32{0x6B, 0x8E, 0x23, 0xFF},
		Orange:               &U32{0xFF, 0xA5, 0x00, 0xFF},
		OrangeRed:            &U32{0xFF, 0x45, 0x00, 0xFF},
		Orchid:               &U32{0xDA, 0x70, 0xD6, 0xFF},
		PaleGoldenRod:        &U32{0xEE, 0xE8, 0xAA, 0xFF},
		PaleGreen:            &U32{0x98, 0xFB, 0x98, 0xFF},
		PaleTurquoise:        &U32{0xAF, 0xEE, 0xEE, 0xFF},
		PaleVioletRed:        &U32{0xDB, 0x70, 0x93, 0xFF},
		PapayaWhip:           &U32{0xFF, 0xEF, 0xD5, 0xFF},
		PeachPuff:            &U32{0xFF, 0xDA, 0xB9, 0xFF},
		Peru:                 &U32{0xCD, 0x85, 0x3F, 0xFF},
		Pink:                 &U32{0xFF, 0xC0, 0xCB, 0xFF},
		Plum:                 &U32{0xDD, 0xA0, 0xDD, 0xFF},
		PowderBlue:           &U32{0xB0, 0xE0, 0xE6, 0xFF},
		Purple:               &U32{0x80, 0x00, 0x80, 0xFF},
		RebeccaPurple:        &U32{0x66, 0x33, 0x99, 0xFF},
		Red:                  &U32{0xFF, 0x00, 0x00, 0xFF},
		RosyBrown:            &U32{0xBC, 0x8F, 0x8F, 0xFF},
		RoyalBlue:            &U32{0x41, 0x69, 0xE1, 0xFF},
		SaddleBrown:          &U32{0x8B, 0x45, 0x13, 0xFF},
		Salmon:               &U32{0xFA, 0x80, 0x72, 0xFF},
		SandyBrown:           &U32{0xF4, 0xA4, 0x60, 0xFF},
		SeaGreen:             &U32{0x2E, 0x8B, 0x57, 0xFF},
		SeaShell:             &U32{0xFF, 0xF5, 0xEE, 0xFF},
		Sienna:               &U32{0xA0, 0x52, 0x2D, 0xFF},
		Silver:               &U32{0xC0, 0xC0, 0xC0, 0xFF},
		SkyBlue:              &U32{0x87, 0xCE, 0xEB, 0xFF},
		SlateBlue:            &U32{0x6A, 0x5A, 0xCD, 0xFF},
		SlateGray:            &U32{0x70, 0x80, 0x90, 0xFF},
		SlateGrey:            &U32{0x70, 0x80, 0x90, 0xFF},
		Snow:                 &U32{0xFF, 0xFA, 0xFA, 0xFF},
		SpringGreen:          &U32{0x00, 0xFF, 0x7F, 0xFF},
		SteelBlue:            &U32{0x46, 0x82, 0xB4, 0xFF},
		Tan:                  &U32{0xD2, 0xB4, 0x8C, 0xFF},
		Teal:                 &U32{0x00, 0x80, 0x80, 0xFF},
		Thistle:              &U32{0xD8, 0xBF, 0xD8, 0xFF},
		Tomato:               &U32{0xFF, 0x63, 0x47, 0xFF},
		Turquoise:            &U32{0x40, 0xE0, 0xD0, 0xFF},
		Violet:               &U32{0xEE, 0x82, 0xEE, 0xFF},
		Wheat:                &U32{0xF5, 0xDE, 0xB3, 0xFF},
		White:                &U32{0xFF, 0xFF, 0xFF, 0xFF},
		WhiteSmoke:           &U32{0xF5, 0xF5, 0xF5, 0xFF},
		Yellow:               &U32{0xFF, 0xFF, 0x00, 0xFF},
		YellowGreen:          &U32{0x9A, 0xCD, 0x32, 0xFF},
	}
	HTML._c = map[string]*U32{
		"aliceblue":            HTML.AliceBlue,
		"antiquewhite":         HTML.AntiqueWhite,
		"aqua":                 HTML.Aqua,
		"aquamarine":           HTML.Aquamarine,
		"azure":                HTML.Azure,
		"beige":                HTML.Beige,
		"bisque":               HTML.Bisque,
		"black":                HTML.Black,
		"blanchedalmond":       HTML.BlanchedAlmond,
		"blue":                 HTML.Blue,
		"blueviolet":           HTML.BlueViolet,
		"brown":                HTML.Brown,
		"burlywood":            HTML.BurlyWood,
		"cadetblue":            HTML.CadetBlue,
		"chartreuse":           HTML.Chartreuse,
		"chocolate":            HTML.Chocolate,
		"coral":                HTML.Coral,
		"cornflowerblue":       HTML.CornflowerBlue,
		"cornsilk":             HTML.Cornsilk,
		"crimson":              HTML.Crimson,
		"cyan":                 HTML.Cyan,
		"darkblue":             HTML.DarkBlue,
		"darkcyan":             HTML.DarkCyan,
		"darkgoldenrod":        HTML.DarkGoldenRod,
		"darkgray":             HTML.DarkGray,
		"darkgrey":             HTML.DarkGrey,
		"darkgreen":            HTML.DarkGreen,
		"darkkhaki":            HTML.DarkKhaki,
		"darkmagenta":          HTML.DarkMagenta,
		"darkolivegreen":       HTML.DarkOliveGreen,
		"darkorange":           HTML.DarkOrange,
		"darkorchid":           HTML.DarkOrchid,
		"darkred":              HTML.DarkRed,
		"darksalmon":           HTML.DarkSalmon,
		"darkseagreen":         HTML.DarkSeaGreen,
		"darkslateblue":        HTML.DarkSlateBlue,
		"darkslategray":        HTML.DarkSlateGray,
		"darkslategrey":        HTML.DarkSlateGrey,
		"darkturquoise":        HTML.DarkTurquoise,
		"darkviolet":           HTML.DarkViolet,
		"deeppink":             HTML.DeepPink,
		"deepskyblue":          HTML.DeepSkyBlue,
		"dimgray":              HTML.DimGray,
		"dimgrey":              HTML.DimGrey,
		"dodgerblue":           HTML.DodgerBlue,
		"firebrick":            HTML.FireBrick,
		"floralwhite":          HTML.FloralWhite,
		"forestgreen":          HTML.ForestGreen,
		"fuchsia":              HTML.Fuchsia,
		"gainsboro":            HTML.Gainsboro,
		"ghostwhite":           HTML.GhostWhite,
		"gold":                 HTML.Gold,
		"goldenrod":            HTML.GoldenRod,
		"gray":                 HTML.Gray,
		"grey":                 HTML.Grey,
		"green":                HTML.Green,
		"greenyellow":          HTML.GreenYellow,
		"honeydew":             HTML.HoneyDew,
		"hotpink":              HTML.HotPink,
		"indianred":            HTML.IndianRed,
		"indigo":               HTML.Indigo,
		"ivory":                HTML.Ivory,
		"khaki":                HTML.Khaki,
		"lavender":             HTML.Lavender,
		"lavenderblush":        HTML.LavenderBlush,
		"lawngreen":            HTML.LawnGreen,
		"lemonchiffon":         HTML.LemonChiffon,
		"lightblue":            HTML.LightBlue,
		"lightcoral":           HTML.LightCoral,
		"lightcyan":            HTML.LightCyan,
		"lightgoldenrodyellow": HTML.LightGoldenRodYellow,
		"lightgray":            HTML.LightGray,
		"lightgrey":            HTML.LightGrey,
		"lightgreen":           HTML.LightGreen,
		"lightpink":            HTML.LightPink,
		"lightsalmon":          HTML.LightSalmon,
		"lightseagreen":        HTML.LightSeaGreen,
		"lightskyblue":         HTML.LightSkyBlue,
		"lightslategray":       HTML.LightSlateGray,
		"lightslategrey":       HTML.LightSlateGrey,
		"lightsteelblue":       HTML.LightSteelBlue,
		"lightyellow":          HTML.LightYellow,
		"lime":                 HTML.Lime,
		"limegreen":            HTML.LimeGreen,
		"linen":                HTML.Linen,
		"magenta":              HTML.Magenta,
		"maroon":               HTML.Maroon,
		"mediumaquamarine":     HTML.MediumAquaMarine,
		"mediumblue":           HTML.MediumBlue,
		"mediumorchid":         HTML.MediumOrchid,
		"mediumpurple":         HTML.MediumPurple,
		"mediumseagreen":       HTML.MediumSeaGreen,
		"mediumslateblue":      HTML.MediumSlateBlue,
		"mediumspringgreen":    HTML.MediumSpringGreen,
		"mediumturquoise":      HTML.MediumTurquoise,
		"mediumvioletred":      HTML.MediumVioletRed,
		"midnightblue":         HTML.MidnightBlue,
		"mintcream":            HTML.MintCream,
		"mistyrose":            HTML.MistyRose,
		"moccasin":             HTML.Moccasin,
		"navajowhite":          HTML.NavajoWhite,
		"navy":                 HTML.Navy,
		"oldlace":              HTML.OldLace,
		"olive":                HTML.Olive,
		"olivedrab":            HTML.OliveDrab,
		"orange":               HTML.Orange,
		"orangered":            HTML.OrangeRed,
		"orchid":               HTML.Orchid,
		"palegoldenrod":        HTML.PaleGoldenRod,
		"palegreen":            HTML.PaleGreen,
		"paleturquoise":        HTML.PaleTurquoise,
		"palevioletred":        HTML.PaleVioletRed,
		"papayawhip":           HTML.PapayaWhip,
		"peachpuff":            HTML.PeachPuff,
		"peru":                 HTML.Peru,
		"pink":                 HTML.Pink,
		"plum":                 HTML.Plum,
		"powderblue":           HTML.PowderBlue,
		"purple":               HTML.Purple,
		"rebeccapurple":        HTML.RebeccaPurple,
		"red":                  HTML.Red,
		"rosybrown":            HTML.RosyBrown,
		"royalblue":            HTML.RoyalBlue,
		"saddlebrown":          HTML.SaddleBrown,
		"salmon":               HTML.Salmon,
		"sandybrown":           HTML.SandyBrown,
		"seagreen":             HTML.SeaGreen,
		"seashell":             HTML.SeaShell,
		"sienna":               HTML.Sienna,
		"silver":               HTML.Silver,
		"skyblue":              HTML.SkyBlue,
		"slateblue":            HTML.SlateBlue,
		"slategray":            HTML.SlateGray,
		"slategrey":            HTML.SlateGrey,
		"snow":                 HTML.Snow,
		"springgreen":          HTML.SpringGreen,
		"steelblue":            HTML.SteelBlue,
		"tan":                  HTML.Tan,
		"teal":                 HTML.Teal,
		"thistle":              HTML.Thistle,
		"tomato":               HTML.Tomato,
		"turquoise":            HTML.Turquoise,
		"violet":               HTML.Violet,
		"wheat":                HTML.Wheat,
		"white":                HTML.White,
		"whitesmoke":           HTML.WhiteSmoke,
		"yellow":               HTML.Yellow,
		"yellowgreen":          HTML.YellowGreen,
	}
}
