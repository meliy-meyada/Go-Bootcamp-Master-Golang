package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = "[command] [string]\n\nAvailable commands: lower, upper and title"


func main(){
	arg := os.Args
	if len(arg) != 3 {
		fmt.Println(usage)
		return
	}

	c, s := arg[1], arg[2]

	switch c {
	case "lower":
		fmt.Println(strings.ToLower(s))
	case "upper":
		fmt.Println(strings.ToUpper(s))
	case "title":
		fmt.Println(strings.Title(s))
	default:
		fmt.Printf("Unknown command: %q\n",c)
	}

}