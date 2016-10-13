package handlers

import "testing"

func TestSanity(t *testing.T) {
	if false {
		t.Error("nothing to see here")
	}
}
func TestInSanity(t *testing.T) {
	if true {
		t.Error("nothing to see here")
	}
}
