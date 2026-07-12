package newmake

import "testing"

func TestExample(t *testing.T) {
	got := Example()
	want := "age=26 user=Maria numbers=[10] len=1 cap=3 nil=true emptyNil=false"
	if got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
