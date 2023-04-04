package main

import (
	"fmt"
	"strconv"
)

func multiplicationTable(size int) string {
	if size <= 0 {
		return "Wrong size"
	}

	// create a buffer to build the table
	table := ""

	// print the header
	table += fmt.Sprintf("%5s", "X")
	for i := 0; i <= size; i++ {
		table += fmt.Sprintf("%5d", i)
	}
	table += "\n"

	for i := 0; i <= size; i++ {
		// print the vertical header
		table += fmt.Sprintf("%5d", i)

		// print the cells
		for j := 0; j <= size; j++ {
			table += fmt.Sprintf("%5d", i*j)
		}
		table += "\n"
	}

	return table
}

func main() {
	sizeStr := "5"
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		fmt.Println("Invalid size")
		return
	}

	table := multiplicationTable(size)
	fmt.Println(table)
}
