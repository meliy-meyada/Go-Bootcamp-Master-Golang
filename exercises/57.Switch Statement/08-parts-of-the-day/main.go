package main

import (
	"fmt"
	"time"
)


func main(){


	switch hour := time.Now().Hour(); {
	case hour >= 18: // 18 to 23
		fmt.Println("good evening")
	case hour >= 12: // 12 to 18
		fmt.Println("good afternoon")
	case hour >= 6: // 6 to 12
		fmt.Println("good morning")
	default: // 0 to 5
		fmt.Println("good night")
	}
}