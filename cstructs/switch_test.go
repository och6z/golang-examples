package cstructs_test

import (
	"fmt"
	"testing"

	"example.com/m/cstructs"
)

func ExampleSwitchTrue() {
	switch true {
	case i5 < 5:
		fmt.Println(assertion)
	case i5 == 5:
		fmt.Println(assertion)
	case i5 > 5:
		fmt.Println(assertion)
	}
	// output: assertion true
}
func TestSwitchTrue(t *testing.T) {
	got := cstructs.SwitchTrue()
	want := "assertion true"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
