package main

import (
	"fmt"
	"os"
)


func main(){
	arg := os.Args
	
	if len(arg) !=2 || len(arg[1]) != 1  {
		fmt.Println("Give me a letter")
		return
	}
	
	sca := arg[1]
	if sca == "a" || sca == "e" || sca == "i" || sca == "o" || sca == "u" {
		fmt.Printf("%q is a vowel.\n", sca)
	} else if sca == "w" || sca == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", sca)
	} else {
		fmt.Printf("%q is a consonant.\n", sca)
	}

	

}