package errorsdemo

import (
	"strings"
	"testing"
)

func TestExampleOutput(t *testing.T) {
	got := Example()

	want := strings.Join([]string{
		"1) nil error -> true",
		"2) errors.New -> email is empty",
		"3) errors.Is with %w -> true",
		"4) anti-example with %v -> false",
		"5) branch with errors.Is -> branch: user not found",
		"6) errors.As -> true, field=age, reason=must be positive",
		"7) anti-example compare text -> false",
	}, "\n")

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}
