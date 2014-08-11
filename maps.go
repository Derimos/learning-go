// 10. Maps

package main 

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Advertiser struct {
	id int
	campaignid int16
}

var m map[string]Vertex

func main() {
	// A map maps keys to values
	m = make(map[string]Vertex) // made with a make construct
	m["Digimedia Tech"] = Vertex {
		40.01, -70.555,
	}

	fmt.Println(m["Digimedia Tech"])

	// Let's make something similar
	var intmap map[string]int

	intmap = make(map[string]int)
	intmap["number ten"] = 10

	fmt.Println(intmap)
	fmt.Println(intmap["number ten"]) 

	// so lets make map of Advertisers using Map literals
	advertisermap := map[string]Advertiser{ // remember that in case of := we don't have to define advertisermap variable
		"C and A" : Advertiser {
			123, 123, // comma is mandatory
		}, // comma is mandatory

		"Hurtamurtalalalla" : Advertiser {
			333, 444,
		},
		"Trade Doubler" : { // or we can omit type
			4444, 6666,
		},
	}

	fmt.Println(advertisermap) 

	// lets try to make int key
	var intkeymap map[int]float64 // this actually is a "type" map
	intkeymap = make(map[int]float64)
	intkeymap[234] = 5.444 // it's not an array

	fmt.Printf("\nMap with int key is of type %T and has value %v. \n", intkeymap, intkeymap)

	// or you can quickly
	somemap := map[float32]string {
		5.444 : "something interesting!!", // float key with string value
		6.45332 : " some other value! not",
	}

	fmt.Printf("\nWhole map: %v \nand one element of map: %v\n\n", somemap, somemap[5.444])

	// How maps mutate
	// Changing, inserting, deleting and checking of map elements

	clientmap := make(map[string]int)
	 // insert value
	clientmap["Number of clients"] = 100
	fmt.Println("Number of clients:", clientmap["Number of clients"])

	// change that value
	clientmap["Number of clients"] = 200 
	fmt.Println("Number of clients:", clientmap["Number of clients"])	

	// delete that element
	delete(clientmap, "Number of clients")

	// checking for that specific key
	v, ok := clientmap["Number of clents"] // ok will be false because we don't have that key in this map
	fmt.Println("The value:", v, "Present?", ok)


	// Show exercise!
	
}