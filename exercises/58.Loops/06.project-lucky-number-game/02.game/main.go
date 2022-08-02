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

func main(){
	rand.Seed(time.Now().UnixNano())

	arg := os.Args[1:]
	if len(arg) != 1 {
		// fmt.Println("Pick a number.")
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(arg[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for t := 0; t < maxTurns; t++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("🎊 YOU WIN!")
			return
		}
	}
	fmt.Println("💀 YOU LOST... Try again?")
}