package cstructs_test

import (
	"fmt"
	"testing"

	"example.com/m/cstructs"
)

func ExampleForConditionAbsent() {
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

func TestForConditionAbsent(t *testing.T) {
	got := cstructs.ForConditionAbsent()
	want := 6

	assertCorrectMessageInt(got, want, t)
}

func ExampleForCondition() {
	sum = 0
	for sum < 100 {
		sum = sum + 9
	}
	fmt.Println(sum)
	// output: 108
}

func TestForCondition(t *testing.T) {
	got := cstructs.ForCondition()
	want := 99 + 9

	assertCorrectMessageInt(got, want, t)
}
