package main

import (
	"fmt"
	"os"
	"strings"
)


const corpus = "the lazy cat jumps again and gain and again"
func main(){

	words := strings.Fields(corpus)
	query := os.Args[1:]


	// the labels and other names do not share thr same scope
	// this will work even though queries label exists

	// var queries string
	// _= queries

	// this label labels the parent loop below
	// label's scope is the whole func main() 

queries:
			for _, q := range query {
			
			search:
						for i, w := range words {
							switch q {
							case "and", "or", "the":
									break search
							}

							if q == w {
								fmt.Printf("#%-2d: %q\n", i+1, w)

								// find the first word then quit
								continue queries
							}
						}
			}

}