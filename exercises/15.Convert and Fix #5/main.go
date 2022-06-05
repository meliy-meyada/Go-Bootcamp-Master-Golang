package main

import "fmt"



func main() {
	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))
}