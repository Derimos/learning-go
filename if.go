// 6. If construct

package main 

import (
	"fmt"
	"math"
)

// regular syntax
func sqrt(x float64) string {
	if x < 0 { // () are gone and {} are mandatory
		return sqrt(-x) + "i"
	} 
	return fmt.Sprint(math.Sqrt(x)) // wtf is Sprint
}

// short syntax
func pow(x, n, lim float64) float64 {
	// v is visible only in scope of If
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g/n", v, lim)
	}
	return lim
}


func main() {
	
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))

	fmt.Println(pow(3, 2, 10))

}