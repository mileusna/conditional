/*
Package conditinal is go/golang replacement for ternary if/else operator

Many languages have ternary operator which makes it easy and short to assign value inline, for example

	val := (x==y) ? "Value OK" : "Value not OK"

Unfortunately, Go doesn't have ternary operator so the shortest way to write this code in Go would be

	val := "Value not OK"
	if x==y {
		val = "Value OK"
	}

or if you prefere perhapse more readable but longer version

	var val string
	if x==y {
		val = "Value OK"
	} else {
		val = "Value not OK"
	}

This is where conditional package steps in. It provides fuctions that replaces ternary operator for each basic type in Go

	// previouse example, now in only one line
	s := conditional.String(x==y, "Value OK", "Value not OK")

	// for other basic Go types
	n := conditional.Int(x==y, 20, 0)
	u := conditional.UInt(true, 23, 15)
	f := conditional.Float64(true, 23.1, 15.7)
	// etc. etc.

	// even for interface{}
	i := conditional.Interface(x==y, "Great", 10)

Example:

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
    }
*/
package conditional
