package colors

import (
	"testing"
)

func TestUnmarshal_Hexrgb(t *testing.T) {
	red := U32{255, 0, 0, 255}
	if temp, _ := Unmarshal("#F00"); temp.(U32) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_Hexrgba(t *testing.T) {
	red := U32{255, 0, 0, 255}
	if temp, _ := Unmarshal("#F00F"); temp.(U32) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_Hexrrggbb(t *testing.T) {
	red := U32{255, 0, 0, 255}
	if temp, _ := Unmarshal("#FF0000"); temp.(U32) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_Hexrrggbbaa(t *testing.T) {
	red := U32{255, 0, 0, 255}
	if temp, _ := Unmarshal("#FF0000FF"); temp.(U32) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_rgb(t *testing.T) {
	red := U32{255, 0, 0, 255}
	if temp, _ := Unmarshal("rgb(255, 0, 0)"); temp.(U32) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_rgba(t *testing.T) {
	red := U32{255, 0, 0, 255}
	if temp, _ := Unmarshal("rgba(255, 0, 0, 1.)"); temp.(U32) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_hsl(t *testing.T) {
	red := HSLA{1, 0, 0, 1}
	if temp, _ := Unmarshal("hsl(360, 0%, 0%)"); temp.(HSLA) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_hsla(t *testing.T) {
	red := HSLA{1, 0, 0, 1}
	if temp, _ := Unmarshal("hsla(360, 0%, 0%, 1.)"); temp.(HSLA) != red {
		t.Error("parse fail got", temp)
	}
}
func TestUnmarshal_predefine(t *testing.T) {
	if temp, _ := Unmarshal("AliceBlue"); temp.(U32) != *HTML.AliceBlue {
		t.Error("parse fail got", temp)
	}
}

func TestMarshal_Hexrgb(t *testing.T) {
	if temp := Marshal(U32{0xFF, 0xFF, 0xFF, 0xFF}, ExpressionHex, true); temp != "#fff"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_Hexrgba(t *testing.T) {
	if temp := Marshal(U32{0xFF, 0xFF, 0xFF, 0xEE}, ExpressionHex, true); temp != "#fffe"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_Hexrrggbb(t *testing.T) {

	if temp := Marshal(U32{0x12, 0x34, 0x56, 0xFF}, ExpressionHex, true); temp != "#123456"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_Hexrrggbbaa(t *testing.T) {
	if temp := Marshal(U32{0x12, 0x34, 0x56, 0x78}, ExpressionHex, true); temp != "#12345678"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_rgb(t *testing.T) {
	if temp := Marshal(U32{12, 34, 56, 0xFF}, ExpressionRGBA, true); temp != "rgb(12, 34, 56)"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_rgba(t *testing.T) {
	if temp := Marshal(U32{12, 34, 56, 51}, ExpressionRGBA, true); temp != "rgba(12, 34, 56, 0.2)"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_hsl(t *testing.T) {
	if temp := Marshal(HSLA{0, .5, .5, 1.}, ExpressionHSLA, true); temp != "hsl(0, 50%, 50%)"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_hsla(t *testing.T) {
	if temp := Marshal(HSLA{0, .5, .5, .8}, ExpressionHSLA, true); temp != "hsla(0, 50%, 50%, 0.8)"{
		t.Error("Marshal fail, got ", temp)
	}
}

func TestMarshal_predefine(t *testing.T) {
	if temp := Marshal(HTML.AliceBlue, ExpressionPredefinedHTML, true); temp != "aliceblue"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_Unknown(t *testing.T) {
	if temp := Marshal(U32{0x12,0x34, 0x56, 0x78}, ExpressionUnknown, true); temp != "#12345678"{
		t.Error("Marshal fail, got ", temp)
	}
}
func TestMarshal_nilcolor(t *testing.T) {
	if temp := Marshal(nil, ExpressionUnknown, true); temp != MarshalFail{
		t.Error("Marshal fail, got ", temp)
	}
}