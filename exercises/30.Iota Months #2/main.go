package main

import "fmt"


func main(){ 

	// const (
	//  _ = iota
	// 	Jan 
	// 	Feb
	// 	Mar
	// )

	const _ = iota - 1


	const (
		Jan = 1 + iota
		Feb
		Mar
	)

	fmt.Println(Jan, Feb, Mar)



}