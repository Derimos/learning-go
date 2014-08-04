// 7. Structs: initialization, access, pointers, literals, new function

package main

import "fmt"

// struct is a collection of fields (or type with a collection of fields)
type Vertex struct {
	X int
	Y int
	// or X, Y int
}

func main() {

	// 1 .initialization
	fmt.Println(Vertex{1, 2})

	// or like this
	v := Vertex{1,2}

	// 2. struct filed are accessed using a dot
	v.X = 4 // notice = and not := because we already know what is type of v
	// or 
	v.X, v.Y = 4 , 5
	fmt.Println(v.X, v.Y)

	var r Vertex // initialize r with {0, 0}
	// t := Vertex  this will make error
	r.X = 7
	fmt.Println(r)

	// 3. POINTERS
	p := Vertex{1, 2}
	q := &p

	q.X = 5

	fmt.Println(p)

	// 4. Struct literals
	var (
		pvert = Vertex{4, 5}
		qvert = &Vertex{8, 9}
		rvert = Vertex{X: 2} // and Y field is 0, we don't have to follow any order {Y: 1, X: 2}
		zvert = Vertex{} //  {0, 0}
	)

	fmt.Println(pvert, qvert, rvert, zvert)


	// 5. new function
	// var t *T = new(T) allocates empty (zeroed) T value and return pointer to it

	newvert := new(Vertex)
	newvert.X , newvert.Y = 20, 22

	fmt.Println(newvert)

	newint := new(int)
	fmt.Println(newint)


}