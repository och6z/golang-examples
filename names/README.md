# Names

## Examples

The package name is the base name of its source directory.

```golang
package names
```

The visibility of a name outside a package is determined
by whether its first character is upper case.
Use `HelloWorld` or `helloWorld` to write multiword names.

```golang
const (
    HelloWorldPrefix string = "Hello, "
    HelloWorld       string = "World!"
    helloWorld       string = "world!"
)
```

### Helloworld

```golang
fmt.Println(HelloWorldPrefix + helloWorld)
```

Output:

```
Hello, world!
```

### Helloworld

```golang
fmt.Println(Helloworld())
```

Output:

```
Hello, world!
```
