// Maps exercise

// Implement WordCount. It should return a map of the counts of each “word” in the string s. 

package main

import (
	"fmt"
	"strings" // we need this to access various string functions
)

func WordCount(s string) map[string]int {

	wordsmap := make(map[string]int) 
	words := strings.Fields(s) // split function argument by words. It will create array of strings

	for _ , word := range words {
		wordsmap[word] = len(word)
	}
    return wordsmap
}

func main() {
    
	fmt.Println(WordCount("We are having tech friday yeah!"))

	fmt.Println(WordCount(""))

}