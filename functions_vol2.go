// 11. More functions

package main

import (
	"fmt"
	"math"
)

func main() {
	
	// functions are values too
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("Type of hypot:%T, hypot value = %f", hypot(3, 4), hypot(3, 4))

	// Welcome to closures
	nextInt := adder()

	fmt.Printf("\n\nTwo has type: %T and value: %d \n",nextInt(), nextInt())
	fmt.Printf("\nTwo has type: %T and value: %d \n",nextInt(), nextInt())
	// sum will be updated every time we call nextInt

	// new function. Has it's own i
	newOne := adder()

	fmt.Printf("\n\nTwo has type: %T and value: %d \n",newOne(), newOne())

	// So the state is unique for that particular function
}

// This function adder returns another function, which we define anonymously in the body of adder. 
// The returned function closes over the variable sum to form a closure.
func adder() func() int {
	sum := 0

	return func() int { // this is just anonimous function
		sum += 1
		return sum
	}
}