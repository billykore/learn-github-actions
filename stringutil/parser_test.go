package stringutil

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	want := 100
	got := ParseInt("100")
	if got != want {
		t.Errorf("failed. want %d, got %d", want, got)
	}
}

func TestParseFloat(t *testing.T) {
	want := 5.555
	got := ParseFloat("5.555")
	if got != want {
		t.Errorf("failed. want %f, got %f", want, got)
	}
}
