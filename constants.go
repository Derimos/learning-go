// 4. Constants

package main

import "fmt"

// can be character, string or boolean values
const Pi = 3.14

const (
	Big = 1 << 100 // didn't figured it out by now
	Small = Big >> 99
)
func needInt(x int) int {
	return x * 10 + 1
}

func main() {
	const World = "世界" // Japaneese .. UTF-8 

	fmt.Println("Hello", World)

	fmt.Println("Hello", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// simple declaration won't work
	// const Simple := "simple" 

	// Get back to this some other time
	fmt.Printf("%T(%v)\n", Small, Small)
	fmt.Println(needInt(Small))


}