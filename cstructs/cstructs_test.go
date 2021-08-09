package cstructs_test

import (
	"testing"

	"example.com/m/cstructs"
)

// If
var (
	HelloWorldPrefix = cstructs.HelloWorldPrefix
)

// Switch
var (
	i5        = cstructs.I5
	assertion = "assertion true"
	hello     = cstructs.Hello
)

func assertCorrectMessage(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
