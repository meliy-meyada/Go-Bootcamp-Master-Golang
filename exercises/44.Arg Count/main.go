package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		arg = os.Args
		count = len(arg) -1
	)

	if count == 0 {
		fmt.Println("Give me args")
	} else if count == 1 {
		fmt.Printf("There is one: %q\n", arg[1])
	} else if count == 2 {
		fmt.Printf(`There are two: "%s %s"`+"\n", arg[1], arg[2])
	} else {
		fmt.Printf("There are %d arguments.\n", count)
	}

}
