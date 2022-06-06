package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	msg := os.Args[1]
	s := msg + strings.Repeat("!", utf8.RuneCountInString(msg))


	fmt.Println(s)
}