package cstructs_test

import (
	"fmt"
	"testing"

	"example.com/m/cstructs"
)

// "If" statements specify the conditional execution of two branches
// according to the value of a boolean expression.
func Example_ifStatement() {
	name := ""
	if name == "" {
		name = "World!"
	}
	fmt.Println(HelloWorldPrefix + name)
	// Output: Hello, World!
}
func TestHelloIf(t *testing.T) {
	got := cstructs.HelloIf()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// The expression may be preceded by a simple statement,
// which executes before the expression is evaluated.
func Example_ifInitializationStatement() {
	if name := "World!"; name != "" {
		fmt.Println(HelloWorldPrefix + name)
	}
	// Output: Hello, World!
}

// func TestHelloIfInitialization(t *testing.T) {
// 	got := cstructs.HelloInitialization()
// 	want := "Hello, World!"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// If the expression evaluates to true, the "if" branch is executed,
// otherwise, if present, the "else" branch is executed.
func Example_ifElseStatement() {
	if name := "World!"; name == "" {
		fmt.Println(HelloWorldPrefix + "")
	} else {
		fmt.Println(HelloWorldPrefix + name)
	}
	// Output: Hello, World!
}
func TestHelloIfElse(t *testing.T) {
	got := cstructs.HelloIfElse()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// The else if statement will be placed between the
// if statement and the else statement.
func Example_ifElseIfStatement() {
	if name := "World!"; name == "" {
		fmt.Println(HelloWorldPrefix + "")
	} else if name == "World!" {
		fmt.Println(HelloWorldPrefix + name)
	} else {
		fmt.Println(HelloWorldPrefix + "Golang!")
	}
	// Output: Hello, World!
}
func TestHelloIfElseIf(t *testing.T) {
	got := cstructs.HelloIfElseIf()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
