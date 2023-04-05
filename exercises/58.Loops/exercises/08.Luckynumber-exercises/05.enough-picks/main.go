package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxGuesses = 5 // the maximum number of guesses allowed
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide a number as an argument.")
		return
	}

	guess, err := toInt(args[0])
	if err != nil {
		fmt.Println("Please provide a valid number as an argument.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please provide a positive number as an argument.")
		return
	}

	for turn := 0; turn < maxGuesses; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			if turn == 0 {
				fmt.Println("🥇 First try! You're a winner!")
			} else {
				fmt.Printf("🎉  You win in %d guesses!\n", turn+1)
			}
			return
		}
	}

	fmt.Println("☠️  You lost... Try again?")
}

// toInt converts a string to an integer and returns an error if the conversion fails
func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
