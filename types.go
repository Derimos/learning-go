// 3. Types

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe bool 		= false
	MaxInt uint64 	= 1<<64 - 1 // seek this out! 
	z complex128 	= cmplx.Sqrt(-5 + 12i) // it complex math variables. Remember math classes
)

func main() {

	const f = "%T(%v)\n"

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// type converstion 
	fmt.Println("\nType Converstion:")
	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(3 * 3 + 4 * 4))
	fmt.Printf(f, f1, f1) // missing value why?

	var z1 int = int(f1)
	fmt.Printf(f, x, x)
	fmt.Printf(f, y, y)
	fmt.Printf(f, f1, f1)
	fmt.Printf(f, z1, z1)
}

// Basic type:

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128