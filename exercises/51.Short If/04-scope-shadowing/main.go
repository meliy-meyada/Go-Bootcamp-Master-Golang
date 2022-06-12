package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n int
		err error
	)
	
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")

	} else if n, err = strconv.Atoi(a[1]); err != nil {

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {

		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}
	fmt.Println("n is", n)
}