package names

import "testing"

func TestHelloworld(t *testing.T) {
	got := HelloWorldPrefix + helloWorld
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
