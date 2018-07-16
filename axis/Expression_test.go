package axis

import "testing"

func TestMarshal(t *testing.T) {
	mb := Marshal(Both)
	t.Log("Marshal Both -> ", mb)
	if mb != "(hori|vert)" {
		t.Error("Unexpected Result Both ", mb)
	}

	mv := Marshal(Vertical)
	t.Log("Marshal Vertical -> ", mv)
	if mv != "(vert)" {
		t.Error("Unexpected Result Vertical ", mv)
	}

	mh := Marshal(Horizontal)
	t.Log("Marshal Horizontal -> ", mh)
	if mh != "(hori)" {
		t.Error("Unexpected Result Horizontal ", mh)
	}

	mn := Marshal(None)
	t.Log("Marshal None -> ", mn)
	if mn != "()" {
		t.Error("Unexpected Result None ", mn)
	}

	mi := Marshal(Invalid)
	t.Log("Marshal Invalid -> ", mi)
	if mi != "<Invalid>" {
		t.Error("Unexpected Result Invalid ", mi)
	}
}

func TestUnmarshal(t *testing.T) {
	var testset = []string{
		"random data",
		"<Invalid>",
		"(h",
		"(x)",
		"()",
		"(h)",
		"(hori)",
		"(ve)",
		"(vert)",
		"(vertical)",
		"(h|v)",
		"(h|h|h)",
		"(h|h|v)",
		"(h|horizontal|v)",
		"(v|hori|v|v|v)",
	}
	var testresult = []Axis{
		Invalid,
		Invalid,
		Invalid,
		Invalid,
		None,
		Horizontal,
		Horizontal,
		Vertical,
		Vertical,
		Vertical,
		Both,
		Horizontal,
		Both,
		Both,
		Both,
	}
	for i, v := range testset {
		if temp := Unmarshal(v); temp != testresult[i]{
			t.Error(v, " expected ", testresult[i], ", but got ", temp)
		}
		t.Log("Pass ", v, testresult[i])
	}
}