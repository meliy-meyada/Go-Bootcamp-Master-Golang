// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//
//  2. Add book titles to the array
//
//  3. Create two more copies of the array named: upper and lower
//
//  4. Change the book titles to uppercase in the upper array only
//
//  5. Change the book titles to lowercase in the lower array only
//
//  6. Print all the arrays
//
//  7. Observe that the arrays are not connected when they're copied.
//
// NOTE
//  Check out the strings package, it has functions to convert letters to
//  upper and lower cases.
//
// BONUS
//  Invent your own arrays with different types other than string,
//  and do some manipulations on them.
//
// EXPECTED OUTPUT
//   Note: Don't worry about the book titles here, you can use any title.
//
//   books: ["Kafka's Revenge" "Stay Golden" "Everythingship"]
//   upper: ["KAFKA'S REVENGE" "STAY GOLDEN" "EVERYTHINGSHIP"]
//   lower: ["kafka's revenge" "stay golden" "everythingship"]
// ---------------------------------------------------------

package main

import (
	"fmt"
	"strings"
)

func main() {
	// create an arrays books with book titles
	books := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	// create copies of the books arrays namded upper and lower
	upper := books
	lower := books

	// change the book titles to uppercase in the upper array
	for i := 0; i < len(upper); i++ {
		upper[i] = strings.ToUpper(upper[i])
	}

	// change the books titles to lowercase in the lower array
	for i := 0; i < len(lower); i++ {
		lower[i] = strings.ToLower(lower[i])
	}

	// print all the arrays
	fmt.Println("books: ", books)
	fmt.Println("upper: ", upper)
	fmt.Println("lower: ", lower)

}
