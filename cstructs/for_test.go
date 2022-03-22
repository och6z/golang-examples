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

	assertCorrectMessageInt(t, got, want)
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

	assertCorrectMessageInt(t, got, want)
}

func ExampleForClause() {
	sum = 0
	for i := 0; i <= 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
	// output: 55
}

func TestForClause(t *testing.T) {
	got := cstructs.ForClause()
	want := 55

	assertCorrectMessageInt(t, got, want)
}

func ExampleForClausePostStmtAbsent() {
	sum = 0
	for i := 0; i <= 10; {
		sum = sum + i
		i++
	}
	fmt.Println(sum)
	// output: 55
}

func TestForClausePostStmtAbsent(t *testing.T) {
	got := cstructs.ForClausePostStmtAbsent()
	want := 55

	assertCorrectMessageInt(t, got, want)
}

func ExampleForRangeClauseArray() {
	sum = 0
	for _, value := range [5]int{0, 1, 2, 3, 4} {
		sum += value
	}
	fmt.Println(sum)
	// output: 10
}

func TestForRangeClauseArray(t *testing.T) {
	got := cstructs.ForRangeClauseArray()
	want := 10

	assertCorrectMessageInt(t, got, want)
}

func ExampleForRangeClauseSlice() {
	sum = 0
	arr := [6]int{0, 1, 2, 3, 4, 5}
	for _, value := range arr[0:5] {
		sum += value
	}
	fmt.Println(sum)
	// output: 10
}

func TestForRangeClauseSlice(t *testing.T) {
	got := cstructs.ForRangeClauseSlice()
	want := 10

	assertCorrectMessageInt(t, got, want)
}

func ExampleForRangeClauseString() {
	for _, rn := range "assertion true" {
		fmt.Print(string(rn))
	}
	// output: assertion true
}

func TestForRangeClauseString(t *testing.T) {
	got := cstructs.ForRangeClauseString()
	want := "assertion true"

	assertCorrectMessage(t, got, want)
}
