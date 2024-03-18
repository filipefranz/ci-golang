package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(30, 40)

	if total != 70 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, 70)
	}
}

func TestSub(t *testing.T) {
	total := Sub(30, 40)

	if total != -10 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, -10)
	}
}

func TestMult(t *testing.T) {
	total := Mult(30, 40)

	if total != 1200 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, 1200)
	}
}

func TestDiv(t *testing.T) {
	total := Div(30, 40)

	if total != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, 0)
	}
}
