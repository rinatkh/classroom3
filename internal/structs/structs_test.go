package structs

import "testing"

func TestExampleOutput(t *testing.T) {
	got := Example()
	want := "before=Maria\nafter=Masha\nempty={ID:0 Name: Active:false}\nactive=true"

	if got != want {
		t.Fatalf("unexpected output:\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}
}
