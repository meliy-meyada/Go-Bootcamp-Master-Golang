package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns    = 5 // less is more difficult
	positiveMsg = "Please pick a positive number."
	notNumber   = "Not a number."
)

const (
	winFirstTime = "🥇 FIRST TIME WINNER!!!"
	winMsg       = "🎉  YOU WON!"
	loseMsg      = "☠️  YOU LOST... Try again?"
)

const (
	welcomeMsg = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(welcomeMsg, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(notNumber)
		return
	}

	if guess <= 0 {
		fmt.Println(positiveMsg)
		return
	}

	if playLuckyNumberGame(guess) {
		fmt.Println(winFirstTime)
	} else {
		fmt.Println(loseMsg)
	}
}

func playLuckyNumberGame(guess int) bool {
	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			if turn == 1 {
				fmt.Println(winFirstTime)
			} else {
				fmt.Println(winMsg)
			}
			return true
		}
	}

	return false
}
