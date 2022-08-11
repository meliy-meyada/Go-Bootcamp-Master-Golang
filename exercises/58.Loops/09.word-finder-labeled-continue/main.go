package main

import (
	"fmt"
	"os"
	"strings"
)


const corpus = "lazy cat jumps gain and again and again"
func main(){

	words := strings.Fields(corpus)
	query := os.Args[1:]

	queries:
					for _, q := range query {
							for i, w := range words {
								if q == w {
									fmt.Printf("#%-2d: %q\n", i+1, w)

									// Skip duplicate words
									continue queries
								}
							}
					}


}