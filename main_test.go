package main

import (
	"testing"

	"example.com/m/names"
)

func TestNames(t *testing.T) {
	got := names.HelloWorldPrefix + names.HelloWorld
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
