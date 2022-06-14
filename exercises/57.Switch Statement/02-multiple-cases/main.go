package main

import (
	"fmt"
	"os"
)


func main(){

	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")

	case "Tokyo":
		fmt.Println("Japan")
	}

}