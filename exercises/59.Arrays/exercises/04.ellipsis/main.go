// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------


package main

import "fmt"

func main() {
	// Declare and initialize arrays using ellipsis
	names := [...]string{"Einstein", "Tesla", "Shepard"}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [...]byte{72, 69, 76, 76, 79}
	ratios := [...]float64{3.14}
	alives := [...]bool{true, false, true, false}

	// Try to assign to the zero array and observe the error
	// Uncommenting the following line will result in an error since zero is an array with no elements
	// zero := [...]int{}

	// Loop through each array using an ordinary loop and print their elements
	fmt.Println("names")
	fmt.Println("====================")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Println()

	fmt.Println("distances")
	fmt.Println("====================")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Println()

	fmt.Println("data")
	fmt.Println("====================")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("====================")
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %f\n", i, ratios[i])
	}

	fmt.Println()

	fmt.Println("alives")
	fmt.Println("====================")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// Uncommenting the following loop will result in an error since zero is an array with no elements
	// fmt.Println()
	// fmt.Println("zero")
	// fmt.Println("====================")
	// for i := 0; i < len(zero); i++ {
	// 	fmt.Printf("zero[%d]: %d\n", i, zero[i])
	// }

	fmt.Println()

	// Loop through each array using for range and print their elements
	fmt.Println("names")
	fmt.Println("====================")
	for i, name := range names {
		fmt.Printf("names[%d]: %q\n", i, name)
	}

	fmt.Println()

	fmt.Println("distances")
	fmt.Println("====================")
	for i, distance := range distances {
		fmt.Printf("distances[%d]: %d\n", i, distance)
	}

	fmt.Println()

	fmt.Println("data")
	fmt.Println("====================")
	for i, datum := range data {
		fmt.Printf("data[%d]: %d\n", i, datum)
	}

	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("====================")
	for i, ratio := range ratios {
		fmt.Printf("ratios[%d]: %f\n", i, ratio)
	}

	fmt.Println()

	fmt.Println("alives")
	fmt.Println("====================")
	for i, alive := range alives {
		fmt.Printf("alives[%d]: %t\n", i, alive)
	}
}
