package main_test

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
	t.Run("If statement", func(t *testing.T) {
		got := cstructs.HelloIf()
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("If else statement", func(t *testing.T) {
		got := cstructs.HelloIfElse()
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("If else if statement", func(t *testing.T) {
		got := cstructs.HelloIfElseIf()
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
