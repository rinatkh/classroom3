package panics

import "testing"

func TestExampleOutput(t *testing.T) {
	got := Example()
	want := "1) recover outside defer -> <nil>\n" +
		"2) without panic -> start -> end -> defer cleanup\n" +
		"3) with panic -> start -> before panic -> recover=unexpected state -> defer cleanup"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}
