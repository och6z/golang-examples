package main

import (
	"testing"

	"example.com/m/cstructs"
	"example.com/m/names"
)

func TestNames(t *testing.T) {
	t.Run("Hello, World!", func(t *testing.T) {
		got := names.HelloWorldPrefix + names.HelloWorld
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("world!", func(t *testing.T) {
		got := names.World()
		want := "world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
func TestCstructs(t *testing.T) {
	got := cstructs.Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
