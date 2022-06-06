package main

import (
	"fmt"
	"os"
	"strings"
)


func main(){



	pack := os.Args[1]
	fmt.Println(strings.ToLower(pack))

}