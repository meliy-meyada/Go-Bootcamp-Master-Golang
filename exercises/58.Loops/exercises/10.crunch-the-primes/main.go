
// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strconv"
)

// isPrime return true id n is a prime number, or false otherwise.
func isPrime(n int) bool {
	switch {
		// Special cases for 2 and 3
		case n == 2 || n == 3:
			return true
		// if n is <= 1 oe divisible by 2 or 3, its not a prime.
		case n <= 1 || n%2 == 0 || n%3 == 0:
			return false
	}
	// check divisibility by numbers of the form 6k + 1, up to sqrt(n)
	for i, w := 5, 2; i*i <= n; {
		if  n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}
	return true
}
func main() {
	// Loop over the command-line arguments and print the prime ones.
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non numeric arguments
			continue
		}
		if isPrime(n){
			fmt.Print(n, " ")
		}
	}
	// print a newline at the end
	fmt.Println()
}