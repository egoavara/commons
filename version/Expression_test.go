package version

import (
	"testing"
	"encoding/json"
	"strings"
)

func TestUnmarshal(t *testing.T) {
	var target = "0.1.15"
	if res := Unmarshal(target); res != Invalid{
		t.Log(res)
	}else {
		t.Error("Fail to parse")
	}
}

type testtype struct {
	Version Version `json:"version"`
}

var testjson = `
{
	"version": "1.2.3"
}`

func TestVersion_UnmarshalJSON(t *testing.T) {
	temp := new(testtype)
	d := json.NewDecoder(strings.NewReader(testjson))
	d.Decode(temp)
	v := Version{1,2,3}
	if temp.Version != v{
		t.Error("result fail")
	}else {
		t.Log("SUCCESS")
	}

}