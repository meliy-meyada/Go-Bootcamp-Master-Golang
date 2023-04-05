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
The program will pick a random number between 1 and the maximum of your two choices.
Your mission is to guess one of those numbers.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage)
		return
	}

	guess1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("First input is not a number.")
		return
	}

	guess2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Second input is not a number.")
		return
	}

	if guess1 <= 0 || guess2 <= 0 {
		fmt.Println("Please pick positive numbers.")
		return
	}

	max := guess1
	if guess2 > guess1 {
		max = guess2
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(max) + 1

		if n == guess1 || n == guess2 {
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
