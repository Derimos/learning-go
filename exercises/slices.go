// Slices Exercise 

package main 

import (
	"fmt"
	// "code.google.com/p/go-tour/pic"
	)

// Make a 2 dimensional slice
// Run this in golang's tour to have a nice image displayed
func Pic(dx, dy int) [][]uint8 {

	var sliceY = make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		var sliceX = make([]uint8, dx)
		for j := range sliceX {
				sliceX[j] = uint8(j*i)
		}
		sliceY[i] = sliceX
	}

	return sliceY
}

func main() {

	var p = Pic(30,30)

	fmt.Print(p)

	// pic.Show(Pic)
}