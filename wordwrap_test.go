package wordwrap

import (
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	got := Wrap("the quick brown fox", 10)
	want := []string{"the quick", "brown fox"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapKeepsLongWordsWhole(t *testing.T) {
	got := Wrap("short internationalization", 8)
	want := []string{"short", "internationalization"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapEmpty(t *testing.T) {
	if got := Wrap("   ", 10); len(got) != 0 {
		t.Errorf("Wrap(blank) = %q, want no lines", got)
	}
}
