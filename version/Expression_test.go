package axis

import "testing"

func TestUnmarshal(t *testing.T) {
	var target = "0.1.15"
	if res := Unmarshal(target); res != Invalid{
		t.Log(res)
	}else {
		t.Error("Fail to parse")
	}
}
