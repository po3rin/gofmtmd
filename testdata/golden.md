# this is test

## with package and import

```go
package main

import "fmt"

func main() {
	fmt.Println("test!!")
}
```

## without package and import

```go
func main() {
	fmt.Println("test!!")
}
```

## expr only

```go
a := 1 + 1
```

## unexpected lines

```go
b := "has lines"






```

## code block (not fenced code)

`a := 1 + 1`

`a::=1+1`

`println!("Hello, world!");`

`Point::new(1.0, 1.0)`

## other language (no format)

```rust
fn main() {
println!("Hello, world!");
}
```
