package cstructs_test

import (
	"testing"

	"example.com/m/cstructs"
)

// If
var (
	HelloWorldPrefix = cstructs.HelloWorldPrefix
)

func assertCorrectMessageInt(got, want int, t *testing.T) {
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

func assertCorrectMessage(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessageBool(got, want bool, t *testing.T) {
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
