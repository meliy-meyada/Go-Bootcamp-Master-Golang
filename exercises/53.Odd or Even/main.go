package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
  

	if old := os.Args; len(old) != 2 {
		println("Pick a number")
	} else if old, err := strconv.Atoi(old[1]); err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
	} else if old%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", old)
	} else if old%2 == 0 {
		fmt.Printf("%d is an even number\n", old)
	} else {
		fmt.Printf("%d is an odd number\n", old)
	}

}