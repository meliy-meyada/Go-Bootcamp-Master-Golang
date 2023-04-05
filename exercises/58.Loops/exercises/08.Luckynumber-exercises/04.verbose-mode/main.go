package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
        maxTurns   = 5 // less is more difficult
        usage      = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
        winEmoji   = "🎉"
)

func main () {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	var verbose bool

	if args[0] == "-v" {
		verbose = true
		args = args[:1]
	}

	targetNumber, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}
	if targetNumber <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}
	if verbose {
		fmt.Println(strings.Repeat("=", 20))
	}
	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(targetNumber) + 1
		if verbose {
			fmt.Printf("%d", n)
		}
		if n == targetNumber{
			if turn == 1 {
				fmt.Printf("%s FIRST TIME WINNER!!!\n", winEmoji)
			}else {
				fmt.Printf("%s YOU WIN!\n", winEmoji)
			}
			return
		}
	}
	if verbose {
		fmt.Println("\n" + strings.Repeat("=", 20))
	}
	fmt.Printf("☠️ YOU LOST... Try again?")

}