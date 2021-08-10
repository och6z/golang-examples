package cstructs_test

import (
	"fmt"
	"testing"

	"example.com/m/cstructs"
)

func ExampleIfTrue() {
	count := 0
	if true {
		count = count + 1
	}
	fmt.Println(count)
	// output: 1
}
func TestIfTrue(t *testing.T) {
	got := cstructs.IfTrue()
	want := 1

	assertCorrectMessageInt(got, want, t)
}
