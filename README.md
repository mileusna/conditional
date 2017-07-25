# Go/Golang replacement for ternary if/else operator [![GoDoc](https://godoc.org/github.com/mileusna/conditional?status.svg)](https://godoc.org/github.com/mileusna/conditional)

Package conditinal is go/golang replacement for ternary if/else operator

Many languages have ternary operator which makes it easy and short to assign value inline, for example

```
// pseudocode, not the go code :(
val := (x==y) ? "Value OK" : "Value not OK" 
```

Unfortunately, [there is no ternary testing operator in Go](https://golang.org/doc/faq#Does_Go_have_a_ternary_form) so the shortest way to write this code in go would be

```Go
val := "Value not OK"
if x==y {
    val = "Value OK"
}
```
or if you prefere perhapse more readable but longer version

```Go
var val string
if x==y {
    val = "Value OK"
} else {
    val = "Value not OK"
}
```
 
This is where conditional package steps in. It provides fuctions that replaces ternary operator for each basic type in go. We can now write conditional assignment in only one Go line:

```Go
val := conditional.String(x==y, "Value OK", "Value not OK")
```
Package conditional also provides fuctions for all Go basic types
```Go
n := conditional.Int(x==y, 20, 0)
u := conditional.UInt(true, 23, 15)
f := conditional.Float64(true, 23.2, 15.1)
// etc. etc.

// even for interface{}
i := conditional.Interface(x==y, "Great", 10)
```
Example:
```Go
package main

import (
    "fmt"
    "github.com/mileusna/conditional"
)

func main() {
    x := 2
    y := 3
    fmt.Println(conditional.Int(x==y, 20, 0))
    fmt.Println(conditional.String(x==y, "Value OK", "Value not OK"))
    fmt.Println(conditional.UInt(true, 23, 15))
    fmt.Println(conditional.Float64(true, 23.4, 15.1))
}
```
