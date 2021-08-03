# Control structures

## Examples

### IfStatement

"If" statements specify the conditional execution of two branches
according to the value of a boolean expression.

```golang
var name = ""
if name == "" {
    name = "World!"
}
fmt.Println(HelloWorldPrefix + name)
```

 Output:

```
Hello, World!
```

### IfInitializationStatement

The expression may be preceded by a simple statement,
which executes before the expression is evaluated.

```golang
if name := "World!"; name != "" {
    fmt.Println(HelloWorldPrefix + name)
}
```

 Output:

```
Hello, World!
```

### IfElseStatement

If the expression evaluates to true, the "if" branch is executed,
otherwise, if present, the "else" branch is executed.

```golang
if name := "World!"; name == "" {
    fmt.Println(HelloWorldPrefix + "")
} else {
    fmt.Println(HelloWorldPrefix + name)
}
```

 Output:

```
Hello, World!
```

### IfElseIfStatement

The else if statement will be placed between the
if statement and the else statement.

```golang
if name := "World!"; name == "" {
    fmt.Println(HelloWorldPrefix + "")
} else if name == "World!" {
    fmt.Println(HelloWorldPrefix + name)
} else {
    fmt.Println(HelloWorldPrefix + "Golang!")
}
```

 Output:

```
Hello, World!
```