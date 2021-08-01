// the package name is the base name of its source directory
package names

// The visibility of a name outside a package is determined
// by whether its first character is upper case.
// Use HelloWorld or helloWorld to write multiword names.
const (
	HelloWorldPrefix string = "Hello, "
	HelloWorld       string = "World!"
	helloWorld       string = "world!"
)

func World() string {
	return helloWorld
}
