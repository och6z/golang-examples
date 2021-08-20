package cstructs_test

import (
	"fmt"
	"testing"

	"example.com/m/cstructs"
)

func ExampleForSingleConditionAbsent() {
	i = 0
	for {
		i = i + 1
		if i > 5 {
			break
		}
	}
	fmt.Println(i)
	// output: 6
}

func TestForSingleConditionAbsent(t *testing.T) {
	got := cstructs.ForSingleConditionAbsent()
	want := 6

	assertCorrectMessageInt(got, want, t)
}
