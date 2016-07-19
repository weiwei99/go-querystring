package query

import "testing"

func TestValues(t *testing.T) {
	a := map[string]int{"a": 0}
	v, _ := Values(a)
	if v.Encode() != "a=0" {
		t.Errorf("v:%s", v.Encode())
	}

	b := map[string]interface{}{"b": "", "c": nil}
	v, _ = Values(b)
	if v.Encode() != "b=" {
		t.Errorf("v:%s", v.Encode())
	}
}
