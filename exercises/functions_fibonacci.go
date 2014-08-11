// Fibonacci closure
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int. :D
func fibonacci() func() int {
	previous := 0
	next := 1
	sum := 0

	return func() int {

		sum = previous + next;
		previous = next;
		next = sum;

		return sum;

	}
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}