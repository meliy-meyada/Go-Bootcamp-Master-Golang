package main

import (
	"fmt"
	"os"
)


func main(){
	fmt.Printf("My name is %s and my lastname is %s.\n", os.Args[1], os.Args[2])
}