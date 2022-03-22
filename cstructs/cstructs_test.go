package cstructs_test

import (
	"testing"

	"example.com/m/cstructs"
)

// If
var (
	HelloWorldPrefix = cstructs.HelloWorldPrefix
)

func assertCorrectMessageInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// Switch
var (
	i5        = cstructs.I5
	assertion = "assertion true"
	hello     = cstructs.Hello
	i7        = cstructs.I7
)

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessageBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

// For
var (
	i   = cstructs.I
	sum = cstructs.Sum
)
