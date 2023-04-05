package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick a random number between 1 and 10 (inclusive) or between 1 and the number you provide (exclusive).
Your mission is to guess the number.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "print the picked numbers")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Printf(usage)
		return
	}

	guess, err := strconv.Atoi(flag.Arg(0))
	if err != nil || guess <= 0 {
		fmt.Println("Please provide a positive integer.")
		return
	}

	min := 10
	if guess > min {
		min = guess
	}

	if verbose {
		fmt.Printf("The target number is between 1 and %d (exclusive).\n", min)
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(min) + 1

		if verbose {
			fmt.Printf("Turn %d: %d\n", turn, n)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
