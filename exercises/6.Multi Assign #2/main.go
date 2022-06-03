package main

import "fmt"

func main() {

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "\n Air is good on Mars",true, 19.5
	fmt.Println(planet, "\n It's", isTrue,"\n It is", temp, "degrees")

}
