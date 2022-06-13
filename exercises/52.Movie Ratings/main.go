package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){

	
	if age := os.Args; len(age) != 2{
		println("Requires age")
		return
	}else if age, err := strconv.Atoi(age[1]); err != nil || age < 0 {
		fmt.Printf(`Wrong age: "%v"` + "\n", age)
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}