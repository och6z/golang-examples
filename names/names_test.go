package names

import (
	"fmt"
	"testing"
)

func TestHelloworld(t *testing.T) {
	got := HelloWorldPrefix + helloWorld
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func Example_helloworld() {
	fmt.Println(HelloWorldPrefix + helloWorld)
	// Output: Hello, world!
}
func ExampleHelloworld() {
	fmt.Println(Helloworld())
	// Output: Hello, world!
}
