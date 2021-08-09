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

	assertCorrectMessage(got, want, t)
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

	assertCorrectMessage(got, want, t)
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

	assertCorrectMessage(got, want, t)
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
		fmt.Println("assertion default")
	}
	// output: assertion true
}
func TestSwitchExpressionDefault(t *testing.T) {
	got := cstructs.SwitchExpressionDefault()
	want := "assertion true"

	assertCorrectMessage(got, want, t)
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
		fmt.Println("assertion default")
	}
	// output: assertion true
}
func TestSwitchExpressionLeftToRightDefault(t *testing.T) {
	got := cstructs.SwitchExpressionLeftToRightDefault()
	want := "assertion true"

	assertCorrectMessage(got, want, t)
}
func ExampleSwitchExpressionTopToBottomDefault() {
	switch i5 {
	case 0:
	case 1:
	case 2:
	case 3:
	case 4:
		fmt.Println(assertion)
	case 5:
		fmt.Println(assertion)
	default:
		fmt.Println("assertion default")
	}
	// output: assertion true
}
func TestSwitchExpressionTopToBottomDefault(t *testing.T) {
	got := cstructs.SwitchExpressionTopToBottomDefault()
	want := "assertion true"

	assertCorrectMessage(got, want, t)
}
func ExampleSwitchExpressionFallthroughDefault() {
	switch i5 {
	case 0:
		dummy := 0
		_ = dummy
		fallthrough
	case 1:
		dummy := 0
		_ = dummy
		fallthrough
	case 2:
		dummy := 0
		_ = dummy
		fallthrough
	case 3:
		dummy := 0
		_ = dummy
		fallthrough
	case 4:
		dummy := 0
		_ = dummy
		fmt.Println(assertion)
	case 5:
		dummy := 0
		_ = dummy
		fallthrough
	default:
		dummy := 0
		_ = dummy
		fmt.Println("assertion default")
	}
	// output: assertion default
}
func TestSwitchExpressionFallthroughDefault(t *testing.T) {
	got := cstructs.SwitchExpressionFallthroughDefault()
	want := "assertion default"

	assertCorrectMessage(got, want, t)
}
func ExampleSwitchVariableExpressionFallthroughDefault() {
	count := 0
	switch i5 {
	case 0:
		count = count + 1
		fallthrough
	case 1:
		count = count + 1
		fallthrough
	case 2:
		count = count + 1
		fallthrough
	case 3:
		count = count + 1
		fallthrough
	case 4:
		count = count + 1
		fmt.Printf(assertion, "\n%d\n", count)
	case 5:
		count = count + 1
		fallthrough
	case 6:
		count = count + 1
		fallthrough
	case 7:
		count = count + 1
		fallthrough
	case 8:
		count = count + 1
		fallthrough
	case 9:
		count = count + 1
		fallthrough
	default:
		count = count + 1
		fmt.Printf("assertion default\n%d\n", count)
	}
	// output: assertion default
	// 6
}
func TestSwitchVariableExpressionFallthroughDefault(t *testing.T) {
	got := cstructs.SwitchVariableExpressionFallthroughDefault()
	want := "assertion default\n6\n"

	assertCorrectMessage(got, want, t)
}
func ExampleSwitchExpressionHelloDefault() {
	switch hello {
	case "wowie":
		fmt.Println(assertion)
	case "hello":
		fmt.Println(assertion)
	case "jumpn":
		fmt.Println(assertion)
	default:
		fmt.Println("assertion default")
	}
	// output: assertion true
}
func TestSwitchExpressionHelloDefault(t *testing.T) {
	got := cstructs.SwitchExpressionHelloDefault()
	want := "assertion true"

	assertCorrectMessage(got, want, t)
}
