// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//	Your goal is to sort the given numbers from the command-line.
//
//	1. Get the numbers from the command-line.
//
//	2. Create an array and assign the given numbers to that array.
//
//	3. Sort the given numbers and print them.
//
// RESTRICTION
//   - Maximum 5 numbers can be provided
//   - If one of the arguments is not a valid number, skip it
//
// HINTS
//
//   - You can use the bubble-sort algorithm to sort the numbers.
//     Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//   - When swapping the elements, do not check for the last element.
//
//     Or, you will receive this error:
//     "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//
//	go run main.go
//	  Please give me up to 5 numbers.
//
//	go run main.go 6 5 4 3 2 1
//	  Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//	go run main.go 5 4 3 2 1
//	  [1 2 3 4 5]
//
//	go run main.go 5 4 a c 1
//	  [0 0 1 4 5]
//
// ---------------------------------------------------------
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	
	if len(args) == 0 || len(args) > 5 {
		fmt.Println("Please give me up to 5 numbers.")
		return

	}
	var nums [5]int

	for i, v := range args{
		num , err := strconv.Atoi(v)
		if err != nil {
			num = 0
		}
		nums[i] = num
	}

	// Bubble-sort algorithm
	for range nums {
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	fmt.Println(nums)
}