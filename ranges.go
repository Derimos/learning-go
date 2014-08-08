// 9. Range is a foreach

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	slice := make( []int, 10)
	for i := range slice {
		slice[i] = 1 << uint(i)
	}

	fmt.Printf("slice = %v \n", slice)


	// if you don't want indexes (or values) for example use underscore _
	for _, value := range slice {
		fmt.Printf("%d ", value)
	}

}