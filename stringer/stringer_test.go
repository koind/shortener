package stringer

import "testing"

var str, shortStr = "5eb63bbbe01eeed093cb22bb8f5acdc3", "5eb63b"

func TestSubstr(t *testing.T) {
	text := Substr(str, 0, 6)

	if text != shortStr {
		t.Errorf("texts does not match %s - %s", shortStr, text)
	}
}
