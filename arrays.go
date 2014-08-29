// 8. Array and Slices

package main

import "fmt"

func main() {

	var a [2]string // array of 2 string values // [n]T
	a[0] = "Hello"
	a[1] = "World"

a[2] = "fool"

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

	// Slices can be resliced : creates a new slice value that points to the same array
	// s[lo:hi] : from lo to hi-1
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	fmt.Println("s[:3] ==", s[:3]) // if we don't have lo index it is 0
	fmt.Println("s[4:] ==", s[4:]) // if we don't have hi index it will be len(s)


	// Making Slices
	slicea := make( []int, 5) // len is five, so slice with 5 zeroes will be created

	slicea[0] = 1;	// this is ok
	// slicea[10] = 12;  index out of range so error
	printSlice("slicea", slicea)


	sliceb := make( []int, 0, 5) // empty slice will be created with capacity of 5
	printSlice("sliceb", sliceb)

	// Nil slices: The zero value of a slice is nil
	// Nil has length and capacity of 0
	var z []int;
	if z == nil {
		fmt.Println("It's nil!")
	}

}

// notice that here we don't have return type notifier
// and this is just another example of function
func printSlice(s string, x []int) /* type is missing*/ {

	// cap is capacity
	fmt.Printf("%s len=%d cap=%d %v \n", s, len(x), cap(x), x) 
	// %s is some sort of Interfacing but we'll leave it for now
}
