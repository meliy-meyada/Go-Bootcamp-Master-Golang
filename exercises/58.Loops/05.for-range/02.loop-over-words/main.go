package main

import (
	"fmt"
	"strings"
)


func main(){

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)
	// 1	2	 3 	4 	5 	6 	7 	8

	// 1st way:
	// for i := 0; i < len(words); i++ {
	// 	fmt.Printf("#%-2d: %q\n", i+1, words[i])
	// }

	// 2nd way:
	// var (
	// 	i int
	// 	v string 
	// )

	// for i, v := range words {
	// 	fmt.Printf("#%-2d: %q\n", i+1, v)
	// }
	// fmt.Printf("Last value of i and v %d %q\n", i, v)


	// 3rd way best:
	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
	
}