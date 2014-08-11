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

	fmt.Println("Hypot = ", hypot(3, 4))

}