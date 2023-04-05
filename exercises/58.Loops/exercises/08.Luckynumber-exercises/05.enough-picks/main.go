package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick a random number between 1 and 10.
Your mission is to guess the number.
You have %d turns.
Wanna play? (y/n)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(usage, maxTurns)

		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		answer = strings.TrimSpace(strings.ToLower(answer))
		if answer != "y" {
			return
		}

		guess := getGuess(reader)
		if guess == 0 {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if playGame(guess) {
			fmt.Println("🎉  YOU WIN!")
		} else {
			fmt.Println("☠️  YOU LOST... Try again?")
		}
	}
}

func getGuess(reader *bufio.Reader) int {
	fmt.Print("Pick a number between 1 and 10: ")

	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return 0
	}

	answer = strings.TrimSpace(answer)

	guess, err := strconv.Atoi(answer)
	if err != nil {
		return 0
	}

	if guess < 1 || guess > 10 {
		return 0
	}

	return guess
}

func playGame(guess int) bool {
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(10) + 1

		if n == guess {
			return true
		}
	}

	return false
}
