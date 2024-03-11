package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(30, 40)

	if total != 70 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, 70)
	}
}
