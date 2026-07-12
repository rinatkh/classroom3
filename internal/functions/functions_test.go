package functions

import "testing"

func TestExampleOutput(t *testing.T) {
	got := Example()
	want := "divide=5\n" +
		"err=nil\n" +
		"user=Maria\n" +
		"ok=true\n" +
		"sum=10\n" +
		"product=12\n" +
		"counter=1/2\n" +
		"defer=first-second-third\n" +
		"args=captured=first current=second"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}
