package main

import (
	"fmt"
	"strconv"
)

func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Printf("There was no error, '%d' is %T\n", n, n)
	} else {
		fmt.Printf("ERROR: %s is not a number.\n", err)
	}
}