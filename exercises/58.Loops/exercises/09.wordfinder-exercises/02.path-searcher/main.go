package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get and split the PATH env variable 

	//SplitList for the path env variable separator for the path env variable
	words := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Search for: ")
		scanner.Scan()
		query := scanner.Text()

		if query == "q"{
			break
		}

		for i, w := range words {
			q, w := strings.ToLower(query), strings.ToLower(w)
			
			if strings.Contains(w, q){
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}