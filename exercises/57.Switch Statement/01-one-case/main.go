package main

import "fmt"

func main(){
	city, f := "Bangkok", "🇹🇭"

	switch city {
	case "Bangkok":
		fmt.Printf("Thailand %s\n", f)
	}
}