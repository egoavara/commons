package axis

import "testing"

func TestAxis_Has(t *testing.T) {
	var v, h bool
	v, h = Both.Has()
	if v != true && h != true{
		t.Error("Invalid result")
	}

	v, h = Vertical.Has()
	if v != true && h != false{
		t.Error("Invalid result")
	}

	v, h = Horizontal.Has()
	if v != false && h != true{
		t.Error("Invalid result")
	}

	v, h = None.Has()
	if v != false && h != false{
		t.Error("Invalid result")
	}

	v, h = Invalid.Has()
	if v != false && h != false{
		t.Error("Invalid result")
	}
}

func TestAxis_HasBoth(t *testing.T) {
	if !Both.HasBoth(){
		t.Error("Invalid result")
	}
	if Horizontal.HasBoth(){
		t.Error("Invalid result")
	}
	if Vertical.HasBoth(){
		t.Error("Invalid result")
	}
	if None.HasBoth(){
		t.Error("Invalid result")
	}
	if Invalid.HasBoth(){
		t.Error("Invalid result")
	}
}

func TestAxis_HasHorizontal(t *testing.T) {
	if !Both.HasHorizontal(){
		t.Error("Invalid result")
	}
	if !Horizontal.HasHorizontal(){
		t.Error("Invalid result")
	}
	if Vertical.HasHorizontal(){
		t.Error("Invalid result")
	}
	if None.HasHorizontal(){
		t.Error("Invalid result")
	}
	if Invalid.HasHorizontal(){
		t.Error("Invalid result")
	}
}

func TestAxis_HasVertical(t *testing.T) {
	if !Both.HasVertical(){
		t.Error("Invalid result")
	}
	if Horizontal.HasVertical(){
		t.Error("Invalid result")
	}
	if !Vertical.HasVertical(){
		t.Error("Invalid result")
	}
	if None.HasVertical(){
		t.Error("Invalid result")
	}
	if Invalid.HasVertical(){
		t.Error("Invalid result")
	}
}

func TestExculsiveAxis_Axis(t *testing.T) {
	if ExculsiveHorizontal.Axis() != Horizontal{
		t.Error("Invalid result")
	}
	if ExculsiveVertical.Axis() != Vertical{
		t.Error("Invalid result")
	}
	if ExculsiveAxis(0xF0).Axis() != Invalid{
		t.Error("Invalid result")
	}
}