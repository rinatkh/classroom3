package arrays

import "testing"

func TestExampleOutput(t *testing.T) {
	got := Example()
	want := "first=10\nlast=30\nlen=3\nzero=[0 0]\noriginal=[1 2 3]\ncopy=[99 2 3]\nequal=true"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}
