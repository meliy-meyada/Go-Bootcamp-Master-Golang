package main

import "fmt"



func main(){
	// Without
	
	i := 142

	switch {
	case i > 100:
		fmt.Print("big positive number")
	case i > 0:
		fmt.Print("positive number")
	default:
		fmt.Print("number")
	}

	fmt.Println()

	// With
	//  i := 142

	// switch {
	// case i > 100:
	// 	fmt.Print("big ")
	// 	fallthrough
	// case i > 0:
	// 	fmt.Print("positive ")
	// 	fallthrough
	// default:
	// 	fmt.Print("number")
	// }

	// fmt.Println()
}