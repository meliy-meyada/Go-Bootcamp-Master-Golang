package main

import (
	"fmt"
	"os"
)

func main(){

	var (
			l = len(os.Args) - 1
			n1 = os.Args[1]
			n2 = os.Args[2]
			n3 = os.Args[3]
	
	)

	fmt.Println("There are ", l, " people !") 
	fmt.Println("Hello greet ", n1, "!")
	fmt.Println("Hello greet ", n2, "!")
	fmt.Println("Hello greet ", n3, "!")
	fmt.Println("Nice to meet you all!.")
	
}