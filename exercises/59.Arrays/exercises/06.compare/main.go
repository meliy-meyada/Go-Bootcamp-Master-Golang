// ---------------------------------------------------------
// EXERCISE: Compare the Arrays
//
//  1. Uncomment the code
//
//  2. Fix the problems so that arrays become comparable
//
// EXPECTED OUTPUT
//  true
//  true
//  false
// ---------------------------------------------------------
package main

import "fmt"

func main() {
	// Define two arrays of different lengths
	week := [...]string{"Monday", "Tuesday"}
	wend := [...]string{"Saturday", "Sunday"}

	// Compare the two arrays
	fmt.Println(week != wend)

	// Define two arrays of the same type and length
	// but different type aliases
	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}

	// Compare the two arrays
	fmt.Println(evens == evens2)

	// Define two arrays of the same length but different
	// element types
	// Use uint8 for one of the arrays instead of byte here.
	// Aliased types are the same types, so uint8 and byte are equivalent.
	image := [5]byte{'h', 'i'}
	var data [5]byte

	// Compare the two arrays
	fmt.Println(data == image)
}