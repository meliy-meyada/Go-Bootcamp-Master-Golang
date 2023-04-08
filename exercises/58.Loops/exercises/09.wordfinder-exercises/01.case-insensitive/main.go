package main

import (
	"flag"
	"fmt"
	"strings"
)



const corpus = "Lazy cat jumps again and again and again"

func main() {
	var wordsMap = make(map[string][]int)
	
	// Parse command-line arguments
	flag.Parse()
	query := flag.Args()

	// split the corpus string into words
	words := strings.Split(corpus, " ")

	// Create a map where the key are the words and the values are their indices in the wordas slice
	for i, word := range words {
		word = strings.ToLower(word)
		wordsMap[word] = append(wordsMap[word], i+1)
	}

	// search for each query word in the  map and print its index in the word slice
	for _, q := range query {
		q = strings.ToLower(q)
		if indices, ok := wordsMap[q]; ok {
			for _, index := range indices {
				fmt.Printf("#%-2d: %q\n", index, q)
			}
		} 
	}




}