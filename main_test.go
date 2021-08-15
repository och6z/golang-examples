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
	t.Run("If True", func(t *testing.T) {
		got := cstructs.IfTrue()
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If False", func(t *testing.T) {
		got := cstructs.IfFalse()
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If initialization True", func(t *testing.T) {
		got := cstructs.IfInitializationTrue()
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If initialization False", func(t *testing.T) {
		got := cstructs.IfInitializationFalse()
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If expression", func(t *testing.T) {
		got := cstructs.IfExpression()
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If else True", func(t *testing.T) {
		got := cstructs.IfElseTrue()
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If else False", func(t *testing.T) {
		got := cstructs.IfElseFalse()
		want := -1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If else initialization False", func(t *testing.T) {
		got := cstructs.IfElseInitializationFalse()
		want := -1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("If else variable False", func(t *testing.T) {
		got := cstructs.IfElseVariableFalse()
		want := -1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
