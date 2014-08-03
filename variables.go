// 2. Variables

package main

import "fmt"

// var defines list of variables, type is last
var i int
var c, python, java bool // default value will be false

// initialization; list of variables with same type, then type and values
var a, b int = 1, 2

// when initializing type can be omited
var t, f, s = true, false, "some string"

// this is all for outside of functions
// outside of a function every construct
// must start with a keyword

func main() {

	fmt.Println(i, c, python, java)

	fmt.Println(t, f, s)

	// while in function we can do this:
	var i, j int = 1, 2 // same as before

	k := 3 // short variable declaration, without var keyword

	// or list declaration
	is, isnot, maybestring := true, false, "perhaps string" 

	fmt.Println(i, j, k, is, isnot, maybestring)

}