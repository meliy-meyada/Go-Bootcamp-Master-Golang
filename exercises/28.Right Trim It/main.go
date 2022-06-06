package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main(){


	name := "inanç           "

	names := strings.TrimRight(name, " ")
	fmt.Println(utf8.RuneCountInString(names))
}