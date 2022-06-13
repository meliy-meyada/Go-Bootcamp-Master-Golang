package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	if year := os.Args; len(year) != 2{
		fmt.Println("Give me a year number")
	} else if year, err := strconv.Atoi(year[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
	} else if year%4==0 && year%100 != 0 || year%400 == 0 {
		fmt.Printf("%v is a leap year.\n", year)
	} else {
		fmt.Printf("%v is not a leap year.\n", year)
	}
}