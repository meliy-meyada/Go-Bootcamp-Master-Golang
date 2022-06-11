package main

import (
	"fmt"
	"os"
)
const (
	usage = "Usage: [username] [password]"
	errUser = "Access denied for %q.\n"
	errPass = "Invalid password for %q.\n"
	okSuccess= "Access granted to %q.\n"
	user,password = "jack", "1888"
)


func main(){

	arg := os.Args
	if len(arg) !=3{
		fmt.Println(usage)
		return
	}


	u, p := arg[1], arg[2]

	if u != user {
		fmt.Printf(errUser,  u)
	} else if p != password {
		fmt.Printf(errPass, u)
	} else {
		fmt.Printf(okSuccess, u)
	}

}