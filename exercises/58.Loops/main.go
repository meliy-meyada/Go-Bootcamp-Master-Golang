package main

import "fmt"

func main(){
	var sum int


	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}