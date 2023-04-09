// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	var names [3]string = [3]string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	var books [8]string = [8]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Harry Potter and the Philosopher's Stone",
		"To Kill a Mockingbird",
		"The Great Gatsby",
		"1984",
		"The Catcher in the Rye",
		"The Lord of the Rings",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
