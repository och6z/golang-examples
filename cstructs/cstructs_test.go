package cstructs_test

import (
	"testing"

	"example.com/m/cstructs"
)

func TestHelloworld(t *testing.T) {
	got := cstructs.Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
