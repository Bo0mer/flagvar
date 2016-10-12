package flagvar_test

import (
	"testing"

	"github.com/Bo0mer/flagvar"
)

func TestArray(t *testing.T) {
	want0 := "one"
	want1 := "two"
	var a flagvar.Array
	err := a.Set(want0)
	if err != nil {
		t.Fatal(err)
	}
	err = a.Set(want1)
	if err != nil {
		t.Fatal(err)
	}

	if len(a) != 2 {
		t.Errorf("expected array of len 2, got %d\n", len(a))
	}

	if a[0] != want0 {
		t.Errorf("expected %q, got %q\n", want0, a[0])
	}
	if a[1] != want1 {
		t.Errorf("expected %q, got %q\n", want1, a[1])
	}

	wantStr := "[one two]"
	if gotStr := a.String(); gotStr != wantStr {
		t.Errorf("expected %q, got %q\n", wantStr, gotStr)
	}
}
