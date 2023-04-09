// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

package main

import "fmt"




func main() {
	// refactored arrays to literals
	names := [3]string{"Einstein", "Tesla", "Shepard"}
	distances := [5]int{50, 40, 75, 30, 125}
	data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, false}
	var zero [0]int

	// ordinary for loops
	fmt.Println("names")
	fmt.Println("====================")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Println("\ndistances")
	fmt.Println("====================")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Println("\ndata")
	fmt.Println("====================")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %c\n", i, data[i])
	}

	fmt.Println("\nratios")
	fmt.Println("====================")
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Println("\nalives")
	fmt.Println("====================")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Println("\nzero")
	fmt.Println("====================")
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	// for range loops
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("FOR RANGES")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Println("\nnames")
	fmt.Println("====================")
	for i, name := range names {
		fmt.Printf("names[%d]: %q\n", i, name)
	}

	fmt.Println("\ndistances")
	fmt.Println("====================")
	for i, distance := range distances {
		fmt.Printf("distances[%d]: %d\n", i, distance)
	}

	fmt.Println("\ndata")
	fmt.Println("====================")
	for i, datum := range data {
		fmt.Printf("data[%d]: %c\n", i, datum)
	}

	fmt.Println("\nratios")
	fmt.Println("====================")
	for i, ratio := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratio)
	}

	fmt.Println("\nalives")
	fmt.Println("====================")
	for i, alive := range alives {
		fmt.Printf("alives[%d]: %t\n", i, alive)
	}

	fmt.Println("\nzero")
	fmt.Println("====================")
	for i := range zero {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}
}