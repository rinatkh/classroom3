package slices

import "testing"

func TestExampleOutput(t *testing.T) {
	got := Example()
	want := "base=[10 99 30 40 77]\npart=[99 30 40 77]\nclone=[555 30 40 77]\nlen=4\ncap=4"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}

func TestAppendMemoryExampleOutput(t *testing.T) {
	got := AppendMemoryExample()
	want := "after make -> s=[] len=0 cap=3\n" +
		"after append 10 -> s=[10] len=1 cap=3\n" +
		"after append 20 -> s=[10 20] len=2 cap=3\n" +
		"after append 30 -> s=[10 20 30] len=3 cap=3\n" +
		"after append 40 -> s=[10 20 30 40] len=4 cap=6"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}

func TestSliceInFunctionExampleOutput(t *testing.T) {
	got := SliceInFunctionExample()
	want := "part=[20 30] len=2 cap=3\n" +
		"after set: part=[999 30] len=2 cap=3\n" +
		"after append1: part=[999 30 777] len=3 cap=3\n" +
		"after append2: part=[999 30 777 888] len=4 cap=6\n" +
		"numbers=[10 999 30 777]"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}
