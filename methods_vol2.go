// 14. Methods with pointer receivers

package main 

import (
	"fmt"
	"math"
)

// first define some type
// in this case for example plain Vertex.
// This kind of type is also called: named type, 
// because we basically give name to some predefined type

type Vertex struct {
	X, Y float64
}

// and then we will define two methods for Vertex type
// and both of them will have pointer receivers
// So we can assosiate methods with named type or pointer to a name type

// And what is a method with a pointer to a named type :)
// And Vertex is method receiver

func (v *Vertex) Scale(f float64) { // doesn't have any return type
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4} // pointer to a name type Vertex

	v.Scale(5)

	fmt.Println(v, v.Abs())

	// There are two reasons to use a pointer receiver. 
	// First, to avoid copying the value on each method 
	// call (more efficient if the value type is a large struct). 
	// Second, so that the method can modify the value that its receiver points to.


	a := &Advertiser{} // empty struct
	a.CreateSimpleId("some simple id not!")

	fmt.Println(a.id)

}

type Advertiser struct {
	id, name, campaignId string
}

func (a *Advertiser) CreateSimpleId(simpleId string) {
	a.id = simpleId 
}


