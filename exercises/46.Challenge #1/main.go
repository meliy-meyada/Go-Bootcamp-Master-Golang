package main

import (
	"fmt"
	"os"
)


func main(){
	arg := os.Args
	l := len(arg) - 1

	if l == 0 || l == 1 || l > 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}
	u, p := arg[1], arg[2]
	
	if u != "jack" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "1888"{
		fmt.Printf("Invalid password for %q.\n", u)
	}else {
		fmt.Printf("Access granted to %q.\n", u)
	}

}