// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := []string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) < 2 {
		fmt.Println("Tell me a book title")
		return
	}

	query := strings.ToLower(os.Args[1])

	var results []string

	for _, book := range books {
		if strings.Contains(strings.ToLower(book), query) {
			results = append(results, book)
		}
	}

	if len(results) == 0 {
		fmt.Printf(`Search Results:
We don't have the book: %q
`, os.Args[1])
		return
	}

	fmt.Println("Search Results:")
	for _, result := range results {
		fmt.Printf("+ %s\n", result)
	}
}






