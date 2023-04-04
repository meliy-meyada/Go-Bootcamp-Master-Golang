package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


const validOps = "%*/+-"

type input struct {
	op string
	size int
}

func parseInput(args []string) (*input, error){
	if len(args) < 2 {
		return nil, fmt.Errorf("usage: [op=%s] [size]", validOps)
	}

	op := args[0]
	if !strings.ContainsAny(op, validOps){
		return nil, fmt.Errorf("invalid operator. Valid ops: %ss", validOps)
	}
	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		return nil, fmt.Errorf("invalid size")
	}
	return &input{op, size}, nil
}

func multiplicationTable(in *input) string {
	table := ""

	// print the header
	table += fmt.Sprintf("%5s", in.op)
	for i := 0; i <= in.size; i++ {
		table += fmt.Sprintf("%5d", i)
	}
	table += "\n"
	
	// print the rows
	for i := 0; i <= in.size; i++ {
		table += fmt.Sprintf("%5d", i)

		// print the cells
		for j := 0; j <= in.size; j++ {
			var res int

			switch in.op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i % j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}
			table += fmt.Sprintf("%5d", res)
		}
		table += "\n"
	}
	return table
}

func main() {
	in, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	table :=  multiplicationTable(in)
	fmt.Println(table)

}