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

### SwitchTrue

```golang
switch true {
case i5 < 5:
    fmt.Println(assertion)
case i5 == 5:
    fmt.Println(assertion)
case i5 > 5:
    fmt.Println(assertion)
}
```

 Output:

```
assertion true
```

### Switch

```golang
switch {
case i5 < 5:
    fmt.Println(assertion)
case i5 == 5:
    fmt.Println(assertion)
case i5 > 5:
    fmt.Println(assertion)
}
```

 Output:

```
assertion true
```

### SwitchInitializationTrue

```golang
switch x := 5; true {
case i5 < x:
    fmt.Println(assertion)
case i5 == x:
    fmt.Println(assertion)
case i5 > x:
    fmt.Println(assertion)
}
```

 Output:

```
assertion true
```

### SwitchExpressionDefault

```golang
switch i5 {
case 0:
    fmt.Println(assertion)
case 1:
    fmt.Println(assertion)
case 2:
    fmt.Println(assertion)
case 3:
    fmt.Println(assertion)
case 4:
    fmt.Println(assertion)
case 5:
    fmt.Println(assertion)
default:
    fmt.Println("assertion default")
}
```

 Output:

```
assertion true
```

### SwitchExpressionLeftToRightDefault

```golang
switch i5 {
case 0, 1, 2, 3:
    fmt.Println(assertion)
case 4, 5, 6, 7:
    fmt.Println(assertion)
case 8, 9:
    fmt.Println(assertion)
default:
    fmt.Println("assertion default")
}
```

 Output:

```
assertion true
```

### SwitchExpressionTopToBottomDefault

```golang
switch i5 {
case 0:
case 1:
case 2:
case 3:
case 4:
    fmt.Println(assertion)
case 5:
    fmt.Println(assertion)
default:
    fmt.Println("assertion default")
}
```

 Output:

```
assertion true
```

### SwitchExpressionFallthroughDefault

```golang
switch i5 {
case 0:
    dummy := 0
    _ = dummy
    fallthrough
case 1:
    dummy := 0
    _ = dummy
    fallthrough
case 2:
    dummy := 0
    _ = dummy
    fallthrough
case 3:
    dummy := 0
    _ = dummy
    fallthrough
case 4:
    dummy := 0
    _ = dummy
    fmt.Println(assertion)
case 5:
    dummy := 0
    _ = dummy
    fallthrough
default:
    dummy := 0
    _ = dummy
    fmt.Println("assertion default")
}
```

 Output:

```
assertion default
```

### SwitchVariableExpressionFallthroughDefault

```golang
count := 0
switch i5 {
case 0:
    count = count + 1
    fallthrough
case 1:
    count = count + 1
    fallthrough
case 2:
    count = count + 1
    fallthrough
case 3:
    count = count + 1
    fallthrough
case 4:
    count = count + 1
    fmt.Printf(assertion, "\n%d\n", count)
case 5:
    count = count + 1
    fallthrough
case 6:
    count = count + 1
    fallthrough
case 7:
    count = count + 1
    fallthrough
case 8:
    count = count + 1
    fallthrough
case 9:
    count = count + 1
    fallthrough
default:
    count = count + 1
    fmt.Printf("assertion default\n%d\n", count)
}
```

 Output:

```
assertion default
6
```

### SwitchExpressionHelloDefault

```golang
switch hello {
case "wowie":
    fmt.Println(assertion)
case "hello":
    fmt.Println(assertion)
case "jumpn":
    fmt.Println(assertion)
default:
    fmt.Println("assertion default")
}
```

 Output:

```
assertion true
```

### SwitchExpressionDefaultVariable

```golang
fired := false
switch i := i5 + 2; i {
case i7:
    fired = true
default:
    fired = false
}
fmt.Println(fired)
```

 Output:

```
true
```

### SwitchDefault

```golang
fired := false
switch {
default:
    fired = true
}
fmt.Println(fired)
```

 Output:

```
true
```