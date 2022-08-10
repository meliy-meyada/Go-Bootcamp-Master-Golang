package main

import (
	"fmt"
	"os"
	"strings"
)



const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	// after inner loops breaks
	// this parent loops will look for thr next queries word
	for _, q := range query {
		// "break" will terminate this loop
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1,w)

				// find the first word then break
				// thr nested loop
				break
			}
		}


	}

}