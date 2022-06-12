package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("42")

	if err == nil {
		fmt.Printf("There was no error, '%d' is %T\n", n, n)
	} else {
		fmt.Printf("ERROR: %s is not a number.\n", err)
	}
}
