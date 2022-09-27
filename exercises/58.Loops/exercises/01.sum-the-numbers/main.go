package main


import "fmt"

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += 1
	}
	fmt.Println("Sum: ", sum)
}