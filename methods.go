// 13. Methods

package main 

import (
	"fmt"
	"math"
) 

type Vertex struct {
	X, Y float64
}

// You can define methods on struct type
// the method receiver is in this case pointer to Vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Or any other type you define in you package
type MyFloat float64 // so we have our out float64

func (f MyFloat) Abs() float64 { // and Abs method for MyFloat. Notice same func name
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4} 
	fmt.Println(v.Abs())


	f := MyFloat(math.Sqrt2)
	fmt.Println(f.Abs())

}