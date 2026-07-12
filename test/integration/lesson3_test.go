package integration

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCommandOutput(t *testing.T) {
	tests := []struct {
		cmd  string
		want string
	}{
		{cmd: "./cmd/01_errors", want: "1) nil error -> true\n2) errors.New -> email is empty\n3) errors.Is with %w -> true\n4) anti-example with %v -> false\n5) branch with errors.Is -> branch: user not found\n6) errors.As -> true, field=age, reason=must be positive\n7) anti-example compare text -> false"},
		{cmd: "./cmd/02_arrays", want: "first=10\nlast=30\nlen=3\nzero=[0 0]\noriginal=[1 2 3]\ncopy=[99 2 3]\nequal=true"},
		{cmd: "./cmd/03_structs", want: "before=Maria\nafter=Masha\nempty={ID:0 Name: Active:false}\nactive=true"},
		{cmd: "./cmd/04_new_make", want: "age=26 user=Maria numbers=[10] len=1 cap=3 nil=true emptyNil=false"},
		{cmd: "./cmd/05_slices", want: "base=[10 99 30 40 77]\npart=[99 30 40 77]\nclone=[555 30 40 77]\nlen=4\ncap=4"},
		{cmd: "./cmd/06_loops", want: "sum=15\ncountdown=0\niterations=2\nnumbers=[10 20 30]\nactive=1"},
		{cmd: "./cmd/07_functions", want: "divide=5\n" +
			"err=nil\n" +
			"user=Maria\n" +
			"ok=true\n" +
			"sum=10\n" +
			"product=12\n" +
			"counter=1/2\n" +
			"defer=first-second-third\n" +
			"args=captured=first current=second",
		},
		{cmd: "./cmd/08_panics", want: "1) recover outside defer -> <nil>\n" +
			"2) without panic -> start -> end -> defer cleanup\n" +
			"3) with panic -> start -> before panic -> recover=unexpected state -> defer cleanup"},
	}

	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			cmd := exec.Command("go", "run", tt.cmd)
			cmd.Dir = "../.."
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go run failed: %v\n%s", err, out)
			}
			got := strings.TrimSpace(string(out))
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
