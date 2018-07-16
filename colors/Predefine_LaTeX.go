package colors

import (
	"strings"
)

var LaTeX _LaTeX

type _LaTeX struct {
	_c             map[string]*U32
	Apricot        *U32
	Aquamarine     *U32
	Bittersweet    *U32
	Black          *U32
	Blue           *U32
	BlueGreen      *U32
	BlueViolet     *U32
	BrickRed       *U32
	Brown          *U32
	BurntOrange    *U32
	CadetBlue      *U32
	CarnationPink  *U32
	Cerulean       *U32
	CornflowerBlue *U32
	Cyan           *U32
	Dandelion      *U32
	DarkOrchid     *U32
	Emerald        *U32
	ForestGreen    *U32
	Fuchsia        *U32
	Goldenrod      *U32
	Gray           *U32
	Green          *U32
	GreenYellow    *U32
	JungleGreen    *U32
	Lavender       *U32
	LimeGreen      *U32
	Magenta        *U32
	Mahogany       *U32
	Maroon         *U32
	Melon          *U32
	MidnightBlue   *U32
	Mulberry       *U32
	NavyBlue       *U32
	OliveGreen     *U32
	Orange         *U32
	OrangeRed      *U32
	Orchid         *U32
	Peach          *U32
	Periwinkle     *U32
	PineGreen      *U32
	Plum           *U32
	ProcessBlue    *U32
	Purple         *U32
	RawSienna      *U32
	Red            *U32
	RedOrange      *U32
	RedViolet      *U32
	Rhodamine      *U32
	RoyalBlue      *U32
	RoyalPurple    *U32
	RubineRed      *U32
	Salmon         *U32
	SeaGreen       *U32
	Sepia          *U32
	SkyBlue        *U32
	SpringGreen    *U32
	Tan            *U32
	TealBlue       *U32
	Thistle        *U32
	Turquoise      *U32
	Violet         *U32
	VioletRed      *U32
	White          *U32
	WildStrawberry *U32
	Yellow         *U32
	YellowGreen    *U32
	YellowOrange   *U32
}

func (s _LaTeX) Keys() []string {
	var keys = make([]string, 0, len(s._c))
	for k := range s._c {
		keys = append(keys, k)
	}
	return keys
}

func (s _LaTeX) Find(name string) *U32 {
	return s._c[strings.ToLower(name)]
}
func (s _LaTeX) Type() Expression {
	return ExpressionPredefinedLaTeX
}
func (s _LaTeX) ToString(c U32) string {
	for k, v := range s._c {
		if *v == c {
			return k
		}
	}
	return ""
}

// https://en.wikibooks.org/wiki/LaTeX/Colors
func init() {
	LaTeX = _LaTeX{
		Apricot:        &U32{251, 185, 130, 0xFF},
		Aquamarine:     &U32{0, 181, 190, 0xFF},
		Bittersweet:    &U32{192, 79, 23, 0xFF},
		Black:          &U32{34, 30, 31, 0xFF},
		Blue:           &U32{45, 47, 146, 0xFF},
		BlueGreen:      &U32{0, 179, 184, 0xFF},
		BlueViolet:     &U32{71, 57, 146, 0xFF},
		BrickRed:       &U32{182, 50, 28, 0xFF},
		Brown:          &U32{121, 37, 0, 0xFF},
		BurntOrange:    &U32{247, 146, 29, 0xFF},
		CadetBlue:      &U32{116, 114, 154, 0xFF},
		CarnationPink:  &U32{242, 130, 180, 0xFF},
		Cerulean:       &U32{0, 162, 227, 0xFF},
		CornflowerBlue: &U32{65, 176, 228, 0xFF},
		Cyan:           &U32{0, 174, 239, 0xFF},
		Dandelion:      &U32{253, 188, 66, 0xFF},
		DarkOrchid:     &U32{164, 83, 138, 0xFF},
		Emerald:        &U32{0, 169, 157, 0xFF},
		ForestGreen:    &U32{0, 155, 85, 0xFF},
		Fuchsia:        &U32{140, 54, 140, 0xFF},
		Goldenrod:      &U32{255, 223, 66, 0xFF},
		Gray:           &U32{148, 150, 152, 0xFF},
		Green:          &U32{0, 166, 79, 0xFF},
		GreenYellow:    &U32{223, 230, 116, 0xFF},
		JungleGreen:    &U32{0, 169, 154, 0xFF},
		Lavender:       &U32{244, 158, 196, 0xFF},
		LimeGreen:      &U32{141, 199, 62, 0xFF},
		Magenta:        &U32{236, 0, 140, 0xFF},
		Mahogany:       &U32{169, 52, 31, 0xFF},
		Maroon:         &U32{175, 50, 53, 0xFF},
		Melon:          &U32{248, 158, 123, 0xFF},
		MidnightBlue:   &U32{0, 103, 149, 0xFF},
		Mulberry:       &U32{169, 60, 147, 0xFF},
		NavyBlue:       &U32{0, 110, 184, 0xFF},
		OliveGreen:     &U32{60, 128, 49, 0xFF},
		Orange:         &U32{245, 129, 55, 0xFF},
		OrangeRed:      &U32{237, 19, 90, 0xFF},
		Orchid:         &U32{175, 114, 176, 0xFF},
		Peach:          &U32{247, 150, 90, 0xFF},
		Periwinkle:     &U32{121, 119, 184, 0xFF},
		PineGreen:      &U32{0, 139, 114, 0xFF},
		Plum:           &U32{146, 38, 143, 0xFF},
		ProcessBlue:    &U32{0, 176, 240, 0xFF},
		Purple:         &U32{153, 71, 155, 0xFF},
		RawSienna:      &U32{151, 64, 6, 0xFF},
		Red:            &U32{237, 27, 35, 0xFF},
		RedOrange:      &U32{242, 96, 53, 0xFF},
		RedViolet:      &U32{161, 36, 107, 0xFF},
		Rhodamine:      &U32{239, 85, 159, 0xFF},
		RoyalBlue:      &U32{0, 113, 188, 0xFF},
		RoyalPurple:    &U32{97, 63, 153, 0xFF},
		RubineRed:      &U32{237, 1, 125, 0xFF},
		Salmon:         &U32{246, 146, 137, 0xFF},
		SeaGreen:       &U32{63, 188, 157, 0xFF},
		Sepia:          &U32{103, 24, 0, 0xFF},
		SkyBlue:        &U32{70, 197, 221, 0xFF},
		SpringGreen:    &U32{198, 220, 103, 0xFF},
		Tan:            &U32{218, 157, 118, 0xFF},
		TealBlue:       &U32{0, 174, 179, 0xFF},
		Thistle:        &U32{216, 131, 183, 0xFF},
		Turquoise:      &U32{0, 180, 206, 0xFF},
		Violet:         &U32{88, 66, 155, 0xFF},
		VioletRed:      &U32{239, 88, 160, 0xFF},
		White:          &U32{0, 0, 0, 0xFF},
		WildStrawberry: &U32{238, 41, 103, 0xFF},
		Yellow:         &U32{255, 242, 0, 0xFF},
		YellowGreen:    &U32{152, 204, 112, 0xFF},
		YellowOrange:   &U32{250, 162, 26, 0xFF},
	}
	//
	LaTeX._c = map[string]*U32{
		"apricot":        LaTeX.Apricot,
		"aquamarine":     LaTeX.Apricot,
		"bittersweet":    LaTeX.Bittersweet,
		"black":          LaTeX.Bittersweet,
		"blue":           LaTeX.Blue,
		"bluegreen":      LaTeX.Blue,
		"blueviolet":     LaTeX.BlueViolet,
		"brickred":       LaTeX.BlueViolet,
		"brown":          LaTeX.Brown,
		"burntorange":    LaTeX.Brown,
		"cadetblue":      LaTeX.CadetBlue,
		"carnationpink":  LaTeX.CadetBlue,
		"cerulean":       LaTeX.Cerulean,
		"cornflowerblue": LaTeX.Cerulean,
		"cyan":           LaTeX.Cyan,
		"dandelion":      LaTeX.Cyan,
		"darkorchid":     LaTeX.DarkOrchid,
		"emerald":        LaTeX.DarkOrchid,
		"forestgreen":    LaTeX.ForestGreen,
		"fuchsia":        LaTeX.ForestGreen,
		"goldenrod":      LaTeX.Goldenrod,
		"gray":           LaTeX.Goldenrod,
		"green":          LaTeX.Green,
		"greenyellow":    LaTeX.Green,
		"junglegreen":    LaTeX.JungleGreen,
		"lavender":       LaTeX.JungleGreen,
		"limegreen":      LaTeX.LimeGreen,
		"magenta":        LaTeX.LimeGreen,
		"mahogany":       LaTeX.Mahogany,
		"maroon":         LaTeX.Mahogany,
		"melon":          LaTeX.Melon,
		"midnightblue":   LaTeX.Melon,
		"mulberry":       LaTeX.Mulberry,
		"navyblue":       LaTeX.Mulberry,
		"olivegreen":     LaTeX.OliveGreen,
		"orange":         LaTeX.OliveGreen,
		"orangered":      LaTeX.OrangeRed,
		"orchid":         LaTeX.OrangeRed,
		"peach":          LaTeX.Peach,
		"periwinkle":     LaTeX.Peach,
		"pinegreen":      LaTeX.PineGreen,
		"plum":           LaTeX.PineGreen,
		"processblue":    LaTeX.ProcessBlue,
		"purple":         LaTeX.ProcessBlue,
		"rawsienna":      LaTeX.RawSienna,
		"red":            LaTeX.RawSienna,
		"redorange":      LaTeX.RedOrange,
		"redviolet":      LaTeX.RedOrange,
		"rhodamine":      LaTeX.Rhodamine,
		"royalblue":      LaTeX.Rhodamine,
		"royalpurple":    LaTeX.RoyalPurple,
		"rubinered":      LaTeX.RoyalPurple,
		"salmon":         LaTeX.Salmon,
		"seagreen":       LaTeX.Salmon,
		"sepia":          LaTeX.Sepia,
		"skyblue":        LaTeX.Sepia,
		"springgreen":    LaTeX.SpringGreen,
		"tan":            LaTeX.SpringGreen,
		"tealblue":       LaTeX.TealBlue,
		"thistle":        LaTeX.TealBlue,
		"turquoise":      LaTeX.Turquoise,
		"violet":         LaTeX.Turquoise,
		"violetred":      LaTeX.VioletRed,
		"white":          LaTeX.VioletRed,
		"wildstrawberry": LaTeX.WildStrawberry,
		"yellow":         LaTeX.WildStrawberry,
		"yellowgreen":    LaTeX.YellowGreen,
		"yelloworange":   LaTeX.YellowGreen,
	}
}
