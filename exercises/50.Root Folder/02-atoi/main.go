package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){

	n, err := strconv.Atoi(os.Args[1])

	if err == nil {
		fmt.Printf("Converted number: %d\n", n)
	}else {
		fmt.Printf("Returned error value: %T\n", err)
	}
}