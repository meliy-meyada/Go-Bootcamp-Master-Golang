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

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil || guess <= 0 {
		fmt.Println("Please provide a positive integer as the only argument.")
		return
	}
	var maxNumber int
	if guess > 10 {
		maxNumber = guess + guess/2 //set a higher range if guess is greater than 10 
	}else {
		maxNumber = 10 // set a default maximum number if guess is 10 or less
	}

	fmt.Printf("Guess a number between 1 and %d:\n", maxNumber)

	for turn := 1; turn < maxTurns; turn++{
		fmt.Printf("Turn %d. Your guess: ", turn)

		var input int
		fmt.Scanln(&input)

		if input == guess{
			if turn == 1 {
				fmt.Println("First time WINNER!!!")
			}else {
				fmt.Println("You Won!")
			}
			return
		}
		fmt.Println("Try again!")
	}
	fmt.Println("☠️ You Lost... Try again?")
}