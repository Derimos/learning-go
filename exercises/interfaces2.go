// Interfaces exercise 2.
// Using func with plain type receiver (and not receiver to struct)
// as we said, this functionality is unique for Golang

package main

import (
	"fmt"
)

type Company int

const (
	Mnkys Company  = iota // value 0 // iota is like enum : Mnkys = 0, Digi = 1 ... etc
	Digimediatech 
)

func (c Company) String() string {
	return "Beograd"
}

func main() {
	fmt.Printf("We have an office in: %s \n", Digimediatech)
}