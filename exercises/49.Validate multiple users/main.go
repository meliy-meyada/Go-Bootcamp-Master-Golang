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
	user1, password1 = "jack", "1888"
	user2, password2 = "meyada", "1889"

)


func main(){

	arg := os.Args
	if len(arg) != 3 {
		fmt.Println(usage)
		return
	}


	u, p := arg[1], arg[2]

	if u != user1 && u != user2 {
		fmt.Printf(errUser,  u)
	}else if u == user1 && p == password1 {
		fmt.Printf(okSuccess, u)
	} else if u == user2 && p == password2 {
		fmt.Printf(okSuccess, u)
	} else {
		fmt.Printf(errPass, u)
	}

}