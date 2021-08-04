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
func ExampleSwitch() {
	switch {
	case i5 < 5:
		fmt.Println(assertion)
	case i5 == 5:
		fmt.Println(assertion)
	case i5 > 5:
		fmt.Println(assertion)
	}
	// output: assertion true
}
func TestSwitch(t *testing.T) {
	got := cstructs.Switch()
	want := "assertion true"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func ExampleSwitchInitializationTrue() {
	switch x := 5; true {
	case i5 < x:
		fmt.Println(assertion)
	case i5 == x:
		fmt.Println(assertion)
	case i5 > x:
		fmt.Println(assertion)
	}
	// output: assertion true
}
func TestSwitchInitializationTrue(t *testing.T) {
	got := cstructs.SwitchInitializationTrue()
	want := "assertion true"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func ExampleSwitchExpressionDefault() {
	switch i5 {
	case 0:
		fmt.Println(assertion)
	case 1:
		fmt.Println(assertion)
	case 2:
		fmt.Println(assertion)
	case 3:
		fmt.Println(assertion)
	case 4:
		fmt.Println(assertion)
	case 5:
		fmt.Println(assertion)
	default:
		fmt.Println("default")
	}
	// output: assertion true
}
func TestSwitchExpressionDefault(t *testing.T) {
	got := cstructs.SwitchExpressionDefault()
	want := "assertion true"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func ExampleSwitchExpressionLeftToRightDefault() {
	switch i5 {
	case 0, 1, 2, 3:
		fmt.Println(assertion)
	case 4, 5, 6, 7:
		fmt.Println(assertion)
	case 8, 9:
		fmt.Println(assertion)
	default:
		fmt.Println("default")
	}
	// output: assertion true
}
func TestExpressionSwitchLeftToRightDefault(t *testing.T) {
	got := cstructs.SwitchExpressionLeftToRightDefault()
	want := "assertion true"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
