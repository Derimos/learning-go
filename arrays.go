// 8. Array and Slices

package main

import "fmt"

func main() {

	var a [2]string // array of 2 string values // [n]T
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])

	fmt.Println(a)

	// 10 zeroed value of type int
	var integerarray [10] int
	fmt.Println(integerarray)
	fmt.Println(len(integerarray))

	// Also array can't resize

	// But we have slices
	// Slices are pointers to array of values // []T
	p := []int{1, 2, 3, 4, 5}
	fmt.Println(p)
	
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}