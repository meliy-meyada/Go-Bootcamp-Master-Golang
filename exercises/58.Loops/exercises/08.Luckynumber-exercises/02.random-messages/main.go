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
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Parse the command line argument
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	// Generate a random number to guess
	target := rand.Intn(guess) + 1

	// Play the game for a limited number of turns
	for turn := 0; turn < maxTurns; turn++ {
		// Prompt the user for their guess
		fmt.Printf("Guess #%d: ", turn+1)
		var response int
		fmt.Scanln(&response)

		// Check if the guess is correct
		if response == target {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}

			// Ask the user if they want to play again
			fmt.Print("Play again? (y/n): ")
			var playAgain string
			fmt.Scanln(&playAgain)
			if playAgain == "y" {
				main()
			} else {
				return
			}
		}

		// Prompt the user to guess again
		fmt.Println("Sorry, try again.")
	}

	// The user lost, so ask if they want to play again
	fmt.Println("☠️  YOU LOST...")
	fmt.Print("Play again? (y/n): ")
	var playAgain string
	fmt.Scanln(&playAgain)
	if playAgain == "y" {
		main()
	}
}
