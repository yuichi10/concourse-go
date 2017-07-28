package util

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Add was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSub(t *testing.T) {
	total := Sub(5, 5)
	if total != 0 {
		t.Errorf("Sub was incorrect, got: %d, want: %d.", total, 0)
	}
}
