package main

import "fmt"



func main() {
	var i int

	loop: 
			if i < 3{
				fmt.Println("looping")
				i++
				goto loop
			}
			fmt.Println("Done")	
}