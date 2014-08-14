// Interfaces exercise
package main 

import (
	"fmt"
)

func main() {

	// This is easy Printf is just like C Printf
	fmt.Printf("Tech %s\n", "Friday")

	// But %s will print things other than strings

	// So Printf defines that %s can print any value that 
	// has String() method defined on that type, but with specific signature
	// This one: String() string. So no arguments and it must return string


	fmt.Printf("Tech %s\n", new(Friday)) 
	// remeber that new will return zeroed struct and return
	// pointer to that struct
}

// We will define some named type
type Friday struct{}

// and method witch has pointer receiver to that type
func (f *Friday) String() string {
	return "FRIDAY!!!!!"
}

// Because type Friday has a String() method
// it automatically implements Stringer interface
// and any value of type Friday when printed by Printf and %s
// will automatically call newly defined String method
