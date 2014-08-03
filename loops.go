// 5. Loops .. For

package main

import "fmt"

// Go has only one loop construct
func main() {
	sum := 0
	// usage of semicolon :). brackets are not even optional
	for i := 0; i < 10; i++ {  // {} required
		sum += i
	}
	fmt.Println(sum)

	// you can leave pre and post statments empty
	sum2 := 1;
	for ; sum2 < 1000 ; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// but it is just easyer to drop the semicolons
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	//infinite loop
	// for {}
}