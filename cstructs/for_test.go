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

	assertCorrectMessageInt(got, want, t)
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

	assertCorrectMessageInt(got, want, t)
}

func ExampleForRangeClauseArray() {
	sum = 0
	for _, value := range []int{0, 1, 2, 3, 4} {
		sum += value
	}
	fmt.Println(sum)
	// output: 10
}

func TestForRangeClauseArray(t *testing.T) {
	got := cstructs.ForRangeClauseArray()
	want := 10

	assertCorrectMessageInt(got, want, t)
}
