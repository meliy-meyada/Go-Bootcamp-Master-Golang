package main

import (
	"os"
	"fmt"
	"strings"
)



const corpus = "Lazy cat jumps again and again and again"

func main() {
	
	words := strings.Fields(corpus)
	query := os.Args[1:]
	
	if len(query) == 0 {
		fmt.Println("Please provide one or more search queries.")
		return
	}

	for _, q := range query {
		// case insensitive search
		q = strings.ToLower(q)

		if q == "and" ||q == "or" ||q == "the" {
			continue
		}
		found := false
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				found = true
			}
		}
		if !found {
			fmt.Printf("%q not found\n", q)
		}
	}
}