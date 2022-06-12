package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")

	}else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q %T %[1]q is not a number.\n", a[1], err)

	} else {
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

}