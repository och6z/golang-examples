# Control structures

## Examples

### IfTrue

```golang
count := 0
if true {
    count = count + 1
}
fmt.Println(count)
```

 Output:

```
1
```

### IfFalse

```golang
count := 0
if false {
    count = count + 1
}
fmt.Println(count)
```

 Output:

```
0
```

### IfInitializationTrue

```golang
count := 0
if one := 1; true {
    count = count + one
}
fmt.Println(count)
```

 Output:

```
1
```

### IfInitializationFalse

```golang
count := 0
if one := 1; false {
    count = count + one
}
fmt.Println(count)
```

 Output:

```
0
```

### IfExpression

```golang
count := 0
if i5 < i7 {
    count = count + 1
}
fmt.Println(count)
```

 Output:

```
1
```

### IfElseTrue

```golang
count := 0
if true {
    count = count + 1
} else {
    count = count - 1
}
fmt.Println(count)
```

 Output:

```
1
```

### IfElseFalse

```golang
count := 0
if false {
    count = count + 1
} else {
    count = count - 1
}
fmt.Println(count)
```

 Output:

```
-1
```

### IfElseInitializationFalse

```golang
count := 0
if t := 1; false {
    count = count + 1
    _ = t
    t := 7
    _ = t
} else {
    count = count - t
}
fmt.Println(count)
```

 Output:

```
-1
```

### IfElseVariableFalse

```golang
count := 0
t := 1
if false {
    count = count + 1
    t := 7
    _ = t
} else {
    count = count - t
}
fmt.Println(count)
```

 Output:

```
-1
```

### Variables

```golang
var (
    i5        = 5
    assertion = "assertion true"
    hello     = "hello"
    i7        = 7
)
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
default:
    fmt.Println("assertion default")
}
```

 Output:

```
assertion default
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