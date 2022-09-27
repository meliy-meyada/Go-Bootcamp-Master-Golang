package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	min, err0 := strconv.Atoi(os.Args[1])
	max, err1 := strconv.Atoi(os.Args[2])
	if err0 != nil || err1 != nil || min >= max {
		fmt.Println("Wrong numbers")
		return
	}
	var sum int
	for i := min; i <= max; i++ {
		sum += i
		
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}