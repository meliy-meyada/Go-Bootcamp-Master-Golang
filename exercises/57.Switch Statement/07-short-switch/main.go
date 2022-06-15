package main

import "fmt"


func main() {
	switch i := 10; /* true */ {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}