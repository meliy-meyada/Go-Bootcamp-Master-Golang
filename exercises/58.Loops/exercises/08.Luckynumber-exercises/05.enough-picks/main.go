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
The program will pick a random number between 1 and 10 (inclusive).
Your mission is to guess that number.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("Usage: lucky [-v]")
		return
	}

	fmt.Printf(usage)

	guess, err := strconv.Atoi(getInput("Your guess: "))
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 1 || guess > 10 {
		fmt.Println("Please pick a number between 1 and 10 (inclusive).")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(10) + 1

		if verbose {
			fmt.Printf("Turn %d: picked %d\n", turn, n)
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

func getInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
