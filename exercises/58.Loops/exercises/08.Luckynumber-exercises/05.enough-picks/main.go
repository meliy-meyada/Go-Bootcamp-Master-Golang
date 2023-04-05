package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers between 0 and your chosen number.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	// Set the seed for the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Get the command-line arguments.
	args := os.Args[1:]

	// Check if the user provided an argument.
	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	// Convert the argument to an integer.
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	// Check if the integer is positive.
	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	// Set the minimum value for the random number generation.
	min := guess

	// Loop through the turns.
	for turn := 0; turn < maxTurns; turn++ {
		// Generate a random number between 0 and the user's guess.
		n := rand.Intn(min + 1)

		// Check if the user guessed the number.
		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	// The user did not guess the number.
	fmt.Println("☠️  YOU LOST... Try again?")
}
