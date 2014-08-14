// 15. Interfaces
package main 
import (
	"fmt"
	"math"
)

// Interface type is define by set of methods
type Abser interface {
	Abs() float64 // in this case one method :)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// Value of interface type can hold any value that
	// implements those methods.
	// In this case f is of type MyFloat, and we have Abs function
	// in this type
	a = f


	v := Vertex{3, 4} // this is not a pointer

	// so if we would assign:
	// a = v will produce error
	// instead we must a to a referece of a v
	a = &v // because Abster is define only on *Vertex

	fmt.Println(a.Abs())
}

// so we define some named type
type MyFloat float64

// and define method of that type
// or in other words this method's receiver is type MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// let us define struct type 
type Vertex struct {
	X,Y float64
}
// and then func of that type, but that has pointer as receiver
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
